/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package management

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	hypv1alpha1 "github.com/openshift/hypershift/api/v1alpha1"
	awsinfra "github.com/openshift/hypershift/cmd/infra/aws"
	"github.com/openshift/hypershift/cmd/util"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// ProviderPlatformReconciler reconciles a ProviderPlatform object
type ProviderPlatformReconciler struct {
	client.Client
}

const (
	destroyFinalizer    = "openshift.io/destroy-platform"
	oidcStorageProvider = "oidc-storage-provider-s3-config"
	oidcSPNamespace     = "kube-public"
	AutoInfraLabelName  = "hypershift.openshift.io/auto-created-for-infra"
	InfraLabelName      = "hypershift.openshift.io/infra-id"
)

// +kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=ProviderPlatform,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=ProviderPlatform/status,verbs=get;update;patch

func (r *ProviderPlatformReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logr.FromContext(ctx).WithValues("ProviderPlatform", req.NamespacedName)

	// your logic here
	var pp hypv1alpha1.ProviderPlatform

	err := r.Client.Get(ctx, req.NamespacedName, &pp)
	if err != nil {
		log.Info("ProviderPlatform resource has been deleted " + req.NamespacedName.Name)
		return ctrl.Result{}, nil
	}

	condition := metav1.Condition{
		Type:               string(hyperv1.CloudProviderConfigured),
		ObservedGeneration: pp.Generation,
	}
	iamCondition := metav1.Condition{
		Type:               string(hyperv1.CloudProviderIAMConfigured),
		ObservedGeneration: pp.Generation,
	}

	if pp.Spec.InfraID == "" {
		pp.Spec.InfraID = fmt.Sprintf("%s-%s", pp.GetName(), utilrand.String(5))

		controllerutil.AddFinalizer(&pp, destroyFinalizer)

		if err := r.Client.Update(ctx, &pp); err != nil {
			if apierrors.IsConflict(err) {
				return ctrl.Result{Requeue: true}, nil
			}
			return ctrl.Result{}, fmt.Errorf("failed to update status: %w", err)
		}

		//Update the status.conditions
		condition.Status = metav1.ConditionFalse
		iamCondition.Status = metav1.ConditionFalse
		condition.Message = "Configuring platform with infra-id: " + pp.Spec.InfraID
		iamCondition.Message = "Configuring platform IAM with infra-id: " + pp.Spec.InfraID
		condition.Reason = hyperv1.CloudProviderConfiguredAsExpected
		iamCondition.Reason = hyperv1.CloudProviderIAMConfiguredAsExpected
		if pp.Status.Conditions == nil {
			pp.Status.Conditions = []metav1.Condition{}
		}
		meta.SetStatusCondition(&pp.Status.Conditions, condition)
		meta.SetStatusCondition(&pp.Status.Conditions, iamCondition)

		if err := r.Client.Status().Update(ctx, &pp); err != nil {
			if apierrors.IsConflict(err) {
				return ctrl.Result{Requeue: true}, nil
			}
			return ctrl.Result{}, fmt.Errorf("failed to update status: %w", err)
		}
	}

	var providerSecret v1.Secret
	var pullSecret v1.Secret

	err = r.Client.Get(ctx, types.NamespacedName{Namespace: pp.Namespace, Name: pp.Spec.Platform.AWS.ControlPlaneOperatorCreds.Name}, &providerSecret)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.Client.Get(ctx, types.NamespacedName{Namespace: pp.Namespace, Name: pp.Spec.PullSecret.Name}, &pullSecret)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Destroying Platform infrastructure used by the ProviderPlatform scheduled for deletion
	if pp.DeletionTimestamp != nil {
		dOpts := awsinfra.DestroyInfraOptions{
			AWSCredentialsFile: "",
			AWSKey:             string(providerSecret.Data["aws_access_key_id"]),
			AWSSecretKey:       string(providerSecret.Data["aws_secret_access_key"]),
			Region:             pp.Spec.Platform.AWS.Region,
			BaseDomain:         pp.Spec.DNS.BaseDomain,
			InfraID:            pp.Spec.InfraID,
		}

		condition.Status = metav1.ConditionFalse
		condition.Message = "Destroying ProviderPlatform with infra-id: " + pp.Spec.InfraID
		condition.Reason = hyperv1.CloudProviderConfiguredAsExpected
		meta.SetStatusCondition(&pp.Status.Conditions, condition)

		iamCondition.Status = metav1.ConditionFalse
		iamCondition.Message = "Removing ProviderPlatform IAM with infra-id: " + pp.Spec.InfraID
		iamCondition.Reason = hyperv1.CloudProviderIAMConfiguredAsExpected
		meta.SetStatusCondition(&pp.Status.Conditions, iamCondition)

		if err := r.Client.Status().Update(ctx, &pp); err != nil {
			if apierrors.IsConflict(err) {
				return ctrl.Result{Requeue: true}, nil
			}
			return ctrl.Result{}, fmt.Errorf("failed to update status: %w", err)
		}

		if err = dOpts.DestroyInfra(ctx); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to destroy ProviderPlatform: %w", err)
		}

		iamOpt := awsinfra.DestroyIAMOptions{
			Region:       pp.Spec.Platform.AWS.Region,
			AWSKey:       dOpts.AWSKey,
			AWSSecretKey: dOpts.AWSSecretKey,
			InfraID:      dOpts.InfraID,
		}

		if err := iamOpt.DestroyIAM(ctx); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to delete IAM ProviderPlatform: %w", err)
		}

		if err := destroyOIDCSecrets(r, &pp); err != nil {
			log.Error(err, "Encountered an issue while deleting secrets")
		}

		if err = r.Client.Get(ctx, req.NamespacedName, &pp); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to update ProviderPlatform values when removing finalizer: %w", err)
		}

		controllerutil.RemoveFinalizer(&pp, destroyFinalizer)

		if err := r.Client.Update(ctx, &pp); err != nil {
			if apierrors.IsConflict(err) {
				return ctrl.Result{Requeue: true}, nil
			}
			return ctrl.Result{}, fmt.Errorf("failed to remove finalizer, update status: %w", err)
		}
		return ctrl.Result{}, nil
	}

	origPp := pp.DeepCopy()

	// Skip reconcile based on condition
	if !meta.IsStatusConditionTrue(pp.Status.Conditions, string(hyperv1.CloudProviderConfigured)) {
		// Creating Platform infrastructure used by the ProviderPlatform NodePools and ingress
		o := awsinfra.CreateInfraOptions{
			AWSKey:       string(providerSecret.Data["aws_access_key_id"]),
			AWSSecretKey: string(providerSecret.Data["aws_secret_access_key"]),
			Region:       pp.Spec.Platform.AWS.Region,
			InfraID:      pp.Spec.InfraID,
			Name:         pp.GetName(),
			BaseDomain:   pp.Spec.DNS.BaseDomain,
		}

		result, err := o.CreateInfra(ctx)
		if err != nil {
			log.Error(err, "Could not create infrastructure")
			condition.Status = metav1.ConditionFalse
			condition.Message = err.Error()
			condition.Reason = hyperv1.CloudProviderMisConfiguredReason

			updateStatusConditions(r, pp.DeepCopy(), condition)
			return ctrl.Result{RequeueAfter: 1 * time.Minute, Requeue: true}, nil
		}

		pp.Spec.DNS.PrivateZoneID = result.PrivateZoneID
		pp.Spec.DNS.PublicZoneID = result.PublicZoneID
		pp.Spec.Networking.MachineCIDR = result.ComputeCIDR
		if pp.Spec.Platform.AWS.CloudProviderConfig == nil || pp.Spec.Platform.AWS.CloudProviderConfig.Subnet == nil {
			pp.Spec.Platform.AWS.CloudProviderConfig = &hypv1alpha1.AWSCloudProviderConfig{
				Subnet: &hypv1alpha1.AWSResourceReference{},
			}
		}
		pp.Spec.Platform.AWS.CloudProviderConfig.Subnet.ID = &result.PrivateSubnetID

		pp.Spec.Platform.AWS.CloudProviderConfig.Zone = result.Zone
		pp.Spec.Platform.AWS.CloudProviderConfig.VPC = result.VPCID
		pp.Spec.SecurityGroups = []hypv1alpha1.AWSResourceReference{
			hypv1alpha1.AWSResourceReference{ID: &result.SecurityGroupID},
		}
		condition.Status = metav1.ConditionTrue
		condition.Message = ""
		condition.Reason = hyperv1.CloudProviderConfiguredAsExpected
		meta.SetStatusCondition(&pp.Status.Conditions, condition)
	}

	if !meta.IsStatusConditionTrue(pp.Status.Conditions, string(hyperv1.CloudProviderIAMConfigured)) {
		oidcSPName, oidcSPRegion, iamErr := oidcDiscoveryURL(r, pp.Spec.InfraID)
		if iamErr == nil {
			iamOpt := awsinfra.CreateIAMOptions{
				Region:                          pp.Spec.Platform.AWS.Region,
				AWSKey:                          string(providerSecret.Data["aws_access_key_id"]),
				AWSSecretKey:                    string(providerSecret.Data["aws_secret_access_key"]),
				InfraID:                         pp.Spec.InfraID,
				IssuerURL:                       pp.Spec.IssuerURL,
				AdditionalTags:                  []string{},
				OIDCStorageProviderS3BucketName: oidcSPName,
				OIDCStorageProviderS3Region:     oidcSPRegion,
			}

			var iamInfo *awsinfra.CreateIAMOutput
			iamInfo, iamErr = iamOpt.CreateIAM(ctx, r.Client)
			if iamErr == nil {
				pp.Spec.Platform.AWS.Roles = iamInfo.Roles
				pp.Spec.IssuerURL = iamInfo.IssuerURL
				if iamErr = createOIDCSecrets(r, &pp, iamInfo); iamErr == nil {
					iamCondition.Status = metav1.ConditionTrue
					iamCondition.Message = ""
					iamCondition.Reason = hyperv1.CloudProviderIAMConfiguredAsExpected
				}
			}
		}
		if iamErr != nil {
			iamCondition.Status = metav1.ConditionFalse
			iamCondition.Message = iamErr.Error()
			iamCondition.Reason = hyperv1.CloudProviderIAMMisConfiguredReason
			updateStatusConditions(r, pp.DeepCopy(), iamCondition)
		}
		// Include the condition
		meta.SetStatusCondition(&pp.Status.Conditions, iamCondition)
	}

	if !reflect.DeepEqual(origPp.Spec, pp.Spec) {
		tc := pp.Status.Conditions
		metav1.SetMetaDataLabel(&pp.ObjectMeta, InfraLabelName, pp.Spec.InfraID)
		if err := r.Client.Update(ctx, &pp); err != nil {
			if apierrors.IsConflict(err) {
				log.Error(err, "Conflict encountered when updating ProviderPlatform")
				return ctrl.Result{Requeue: true}, nil
			}
			return ctrl.Result{}, fmt.Errorf("failed to update spec: %w", err)
		}
		pp.Status.Conditions = tc
		if err := r.Client.Status().Update(ctx, &pp); err != nil {
			if apierrors.IsConflict(err) {
				log.Error(err, "Conflict encountered when updating ProviderPlatform.Status")
				return ctrl.Result{Requeue: true}, nil
			}
			return ctrl.Result{}, fmt.Errorf("failed to update status: %w", err)
		}
		log.Info("Finished reconciling ProviderPlatform " + req.Name)
	}
	return ctrl.Result{}, err
}

func (r *ProviderPlatformReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hypv1alpha1.ProviderPlatform{}).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: 1,
		}).
		Complete(r)
}

func oidcDiscoveryURL(r *ProviderPlatformReconciler, infraID string) (string, string, error) {

	cm := &v1.ConfigMap{}
	if err := r.Client.Get(context.Background(), types.NamespacedName{Name: oidcStorageProvider, Namespace: oidcSPNamespace}, cm); err != nil {
		return "", "", err
	}
	return cm.Data["name"], cm.Data["region"], nil
}

func createOIDCSecrets(r *ProviderPlatformReconciler, pp *hypv1alpha1.ProviderPlatform, iamInfo *awsinfra.CreateIAMOutput) error {

	ctx := context.Background()

	buildAWSCreds := func(name, arn string) *corev1.Secret {
		return &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: corev1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: pp.Namespace,
				Name:      name,
				Labels: map[string]string{
					AutoInfraLabelName: pp.Spec.InfraID,
				},
			},
			Data: map[string][]byte{
				"credentials": []byte(fmt.Sprintf(`[default]
	role_arn = %s
	web_identity_token_file = /var/run/secrets/openshift/serviceaccount/token
	`, arn)),
			},
		}
	}

	secretResource := buildAWSCreds(pp.Name+"-cloud-ctrl-creds", iamInfo.KubeCloudControllerRoleARN)
	pp.Spec.Platform.AWS.KubeCloudControllerCreds = corev1.LocalObjectReference{Name: secretResource.Name}
	if err := r.Create(ctx, secretResource); apierrors.IsAlreadyExists(err) {
		if err := r.Update(ctx, secretResource); err != nil {
			return err
		}
	}

	secretResource = buildAWSCreds(pp.Name+"-node-mgmt-creds", iamInfo.NodePoolManagementRoleARN)
	pp.Spec.Platform.AWS.NodePoolManagementCreds = corev1.LocalObjectReference{Name: secretResource.Name}
	if err := r.Create(ctx, secretResource); apierrors.IsAlreadyExists(err) {
		if err := r.Update(ctx, secretResource); err != nil {
			return err
		}

	}
	return nil
}

func destroyOIDCSecrets(r *ProviderPlatformReconciler, pp *hypv1alpha1.ProviderPlatform) error {
	//clean up CLI generated secrets
	return r.DeleteAllOf(context.Background(), &v1.Secret{}, client.InNamespace(pp.GetNamespace()), client.MatchingLabels{util.AutoInfraLabelName: pp.Spec.InfraID})

}

func updateStatusConditions(r *ProviderPlatformReconciler, pp *hypv1alpha1.ProviderPlatform, condition metav1.Condition) error {
	cc := meta.FindStatusCondition(pp.Status.Conditions, condition.Type)
	if cc == nil || cc.ObservedGeneration != condition.ObservedGeneration || cc.Status != condition.Status {
		meta.SetStatusCondition(&pp.Status.Conditions, condition)
		return r.Client.Status().Update(context.Background(), pp)
	}
	return nil
}

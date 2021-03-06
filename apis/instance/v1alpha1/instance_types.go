/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecDataDisks struct {
	// +optional
	AutoSnapshotPolicyID *string `json:"autoSnapshotPolicyID,omitempty" tf:"auto_snapshot_policy_id"`
	// +optional
	Category *string `json:"category,omitempty" tf:"category"`
	// +optional
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty" tf:"delete_with_instance"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PerformanceLevel *string `json:"performanceLevel,omitempty" tf:"performance_level"`
	Size             *int64  `json:"size" tf:"size"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	AllocatePublicIP *bool `json:"allocatePublicIP,omitempty" tf:"allocate_public_ip"`
	// +optional
	AutoReleaseTime *string `json:"autoReleaseTime,omitempty" tf:"auto_release_time"`
	// +optional
	AutoRenewPeriod *int64 `json:"autoRenewPeriod,omitempty" tf:"auto_renew_period"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	CreditSpecification *string `json:"creditSpecification,omitempty" tf:"credit_specification"`
	// +optional
	// +kubebuilder:validation:MaxItems=16
	// +kubebuilder:validation:MinItems=1
	DataDisks []InstanceSpecDataDisks `json:"dataDisks,omitempty" tf:"data_disks"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	DeploymentSetGroupNo *string `json:"deploymentSetGroupNo,omitempty" tf:"deployment_set_group_no"`
	// +optional
	DeploymentSetID *string `json:"deploymentSetID,omitempty" tf:"deployment_set_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run"`
	// +optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete"`
	// +optional
	HostName *string `json:"hostName,omitempty" tf:"host_name"`
	// +optional
	HpcClusterID *string `json:"hpcClusterID,omitempty" tf:"hpc_cluster_id"`
	ImageID      *string `json:"imageID" tf:"image_id"`
	// +optional
	IncludeDataDisks *bool `json:"includeDataDisks,omitempty" tf:"include_data_disks"`
	// +optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name"`
	InstanceType *string `json:"instanceType" tf:"instance_type"`
	// +optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type"`
	// +optional
	// Deprecated
	InternetMaxBandwidthIn *int64 `json:"internetMaxBandwidthIn,omitempty" tf:"internet_max_bandwidth_in"`
	// +optional
	InternetMaxBandwidthOut *int64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out"`
	// +optional
	// Deprecated
	IoOptimized *string `json:"ioOptimized,omitempty" tf:"io_optimized"`
	// +optional
	IsOutdated *bool `json:"isOutdated,omitempty" tf:"is_outdated"`
	// +optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`
	// +optional
	KmsEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password"`
	// +optional
	KmsEncryptionContext map[string]string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	PublicIP *string `json:"publicIP,omitempty" tf:"public_ip"`
	// +optional
	RenewalStatus *string `json:"renewalStatus,omitempty" tf:"renewal_status"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name"`
	// +optional
	SecondaryPrivateIPAddressCount *int64 `json:"secondaryPrivateIPAddressCount,omitempty" tf:"secondary_private_ip_address_count"`
	// +optional
	SecondaryPrivateIPS []string `json:"secondaryPrivateIPS,omitempty" tf:"secondary_private_ips"`
	// +optional
	SecurityEnhancementStrategy *string  `json:"securityEnhancementStrategy,omitempty" tf:"security_enhancement_strategy"`
	SecurityGroups              []string `json:"securityGroups" tf:"security_groups"`
	// +optional
	SpotPriceLimit *float64 `json:"spotPriceLimit,omitempty" tf:"spot_price_limit"`
	// +optional
	SpotStrategy *string `json:"spotStrategy,omitempty" tf:"spot_strategy"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	SystemDiskAutoSnapshotPolicyID *string `json:"systemDiskAutoSnapshotPolicyID,omitempty" tf:"system_disk_auto_snapshot_policy_id"`
	// +optional
	SystemDiskCategory *string `json:"systemDiskCategory,omitempty" tf:"system_disk_category"`
	// +optional
	SystemDiskDescription *string `json:"systemDiskDescription,omitempty" tf:"system_disk_description"`
	// +optional
	SystemDiskName *string `json:"systemDiskName,omitempty" tf:"system_disk_name"`
	// +optional
	SystemDiskPerformanceLevel *string `json:"systemDiskPerformanceLevel,omitempty" tf:"system_disk_performance_level"`
	// +optional
	SystemDiskSize *int64 `json:"systemDiskSize,omitempty" tf:"system_disk_size"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	VolumeTags map[string]string `json:"volumeTags,omitempty" tf:"volume_tags"`
	// +optional
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}

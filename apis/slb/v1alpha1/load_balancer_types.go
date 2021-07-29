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

type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerSpec   `json:"spec,omitempty"`
	Status            LoadBalancerStatus `json:"status,omitempty"`
}

type LoadBalancerSpec struct {
	State *LoadBalancerSpecResource `json:"state,omitempty" tf:"-"`

	Resource LoadBalancerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LoadBalancerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// +optional
	AddressIPVersion *string `json:"addressIPVersion,omitempty" tf:"address_ip_version"`
	// +optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type"`
	// +optional
	Bandwidth *int64 `json:"bandwidth,omitempty" tf:"bandwidth"`
	// +optional
	DeleteProtection *string `json:"deleteProtection,omitempty" tf:"delete_protection"`
	// +optional
	// Deprecated
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +optional
	Internet *bool `json:"internet,omitempty" tf:"internet"`
	// +optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type"`
	// +optional
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name"`
	// +optional
	LoadBalancerSpec *string `json:"loadBalancerSpec,omitempty" tf:"load_balancer_spec"`
	// +optional
	MasterZoneID *string `json:"masterZoneID,omitempty" tf:"master_zone_id"`
	// +optional
	ModificationProtectionReason *string `json:"modificationProtectionReason,omitempty" tf:"modification_protection_reason"`
	// +optional
	ModificationProtectionStatus *string `json:"modificationProtectionStatus,omitempty" tf:"modification_protection_status"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	SlaveZoneID *string `json:"slaveZoneID,omitempty" tf:"slave_zone_id"`
	// +optional
	// Deprecated
	Specification *string `json:"specification,omitempty" tf:"specification"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
}

type LoadBalancerStatus struct {
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

// LoadBalancerList is a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancer CRD objects
	Items []LoadBalancer `json:"items,omitempty"`
}

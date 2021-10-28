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

type Gateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec,omitempty"`
	Status            GatewayStatus `json:"status,omitempty"`
}

type GatewaySpecBandwidthPackages struct {
	Bandwidth *int64 `json:"bandwidth" tf:"bandwidth"`
	IpCount   *int64 `json:"ipCount" tf:"ip_count"`
	// +optional
	PublicIPAddresses *string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses"`
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type GatewaySpec struct {
	State *GatewaySpecResource `json:"state,omitempty" tf:"-"`

	Resource GatewaySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GatewaySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BandwidthPackageIDS *string `json:"bandwidthPackageIDS,omitempty" tf:"bandwidth_package_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	BandwidthPackages []GatewaySpecBandwidthPackages `json:"bandwidthPackages,omitempty" tf:"bandwidth_packages"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run"`
	// +optional
	Force *bool `json:"force,omitempty" tf:"force"`
	// +optional
	ForwardTableIDS *string `json:"forwardTableIDS,omitempty" tf:"forward_table_ids"`
	// +optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NatGatewayName *string `json:"natGatewayName,omitempty" tf:"nat_gateway_name"`
	// +optional
	NatType *string `json:"natType,omitempty" tf:"nat_type"`
	// +optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type"`
	// +optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	SnatTableIDS *string `json:"snatTableIDS,omitempty" tf:"snat_table_ids"`
	// +optional
	Spec *string `json:"spec,omitempty" tf:"spec"`
	// +optional
	Specification *string `json:"specification,omitempty" tf:"specification"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty" tf:"tags"`
	VpcID *string           `json:"vpcID" tf:"vpc_id"`
	// +optional
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
}

type GatewayStatus struct {
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

// GatewayList is a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Gateway CRD objects
	Items []Gateway `json:"items,omitempty"`
}

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

type GatewaySpec struct {
	State *GatewaySpecResource `json:"state,omitempty" tf:"-"`

	Resource GatewaySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GatewaySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bandwidth *int64 `json:"bandwidth" tf:"bandwidth"`
	// +optional
	BusinessStatus *string `json:"businessStatus,omitempty" tf:"business_status"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	EnableIpsec *bool `json:"enableIpsec,omitempty" tf:"enable_ipsec"`
	// +optional
	EnableSsl *bool `json:"enableSsl,omitempty" tf:"enable_ssl"`
	// +optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type"`
	// +optional
	InternetIP *string `json:"internetIP,omitempty" tf:"internet_ip"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Period *int64 `json:"period,omitempty" tf:"period"`
	// +optional
	SslConnections *int64 `json:"sslConnections,omitempty" tf:"ssl_connections"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	VpcID  *string `json:"vpcID" tf:"vpc_id"`
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

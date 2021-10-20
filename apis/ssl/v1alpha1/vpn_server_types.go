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

type VpnServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnServerSpec   `json:"spec,omitempty"`
	Status            VpnServerStatus `json:"status,omitempty"`
}

type VpnServerSpec struct {
	State *VpnServerSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpnServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VpnServerSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cipher       *string `json:"cipher,omitempty" tf:"cipher"`
	ClientIPPool *string `json:"clientIPPool" tf:"client_ip_pool"`
	// +optional
	Compress *bool `json:"compress,omitempty" tf:"compress"`
	// +optional
	Connections *int64 `json:"connections,omitempty" tf:"connections"`
	// +optional
	InternetIP  *string `json:"internetIP,omitempty" tf:"internet_ip"`
	LocalSubnet *string `json:"localSubnet" tf:"local_subnet"`
	// +optional
	MaxConnections *int64 `json:"maxConnections,omitempty" tf:"max_connections"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Protocol     *string `json:"protocol,omitempty" tf:"protocol"`
	VpnGatewayID *string `json:"vpnGatewayID" tf:"vpn_gateway_id"`
}

type VpnServerStatus struct {
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

// VpnServerList is a list of VpnServers
type VpnServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnServer CRD objects
	Items []VpnServer `json:"items,omitempty"`
}

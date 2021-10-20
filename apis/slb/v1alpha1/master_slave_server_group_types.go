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

type MasterSlaveServerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MasterSlaveServerGroupSpec   `json:"spec,omitempty"`
	Status            MasterSlaveServerGroupStatus `json:"status,omitempty"`
}

type MasterSlaveServerGroupSpecServers struct {
	// +optional
	IsBackup *int64  `json:"isBackup,omitempty" tf:"is_backup"`
	Port     *int64  `json:"port" tf:"port"`
	ServerID *string `json:"serverID" tf:"server_id"`
	// +optional
	ServerType *string `json:"serverType,omitempty" tf:"server_type"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type MasterSlaveServerGroupSpec struct {
	State *MasterSlaveServerGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource MasterSlaveServerGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MasterSlaveServerGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DeleteProtectionValidation *bool   `json:"deleteProtectionValidation,omitempty" tf:"delete_protection_validation"`
	LoadBalancerID             *string `json:"loadBalancerID" tf:"load_balancer_id"`
	Name                       *string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=2
	Servers []MasterSlaveServerGroupSpecServers `json:"servers,omitempty" tf:"servers"`
}

type MasterSlaveServerGroupStatus struct {
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

// MasterSlaveServerGroupList is a list of MasterSlaveServerGroups
type MasterSlaveServerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MasterSlaveServerGroup CRD objects
	Items []MasterSlaveServerGroup `json:"items,omitempty"`
}

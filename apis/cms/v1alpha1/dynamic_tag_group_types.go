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

type DynamicTagGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamicTagGroupSpec   `json:"spec,omitempty"`
	Status            DynamicTagGroupStatus `json:"status,omitempty"`
}

type DynamicTagGroupSpecMatchExpress struct {
	TagValue              *string `json:"tagValue" tf:"tag_value"`
	TagValueMatchFunction *string `json:"tagValueMatchFunction" tf:"tag_value_match_function"`
}

type DynamicTagGroupSpec struct {
	State *DynamicTagGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource DynamicTagGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DynamicTagGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ContactGroupList []string                          `json:"contactGroupList" tf:"contact_group_list"`
	MatchExpress     []DynamicTagGroupSpecMatchExpress `json:"matchExpress" tf:"match_express"`
	// +optional
	MatchExpressFilterRelation *string `json:"matchExpressFilterRelation,omitempty" tf:"match_express_filter_relation"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	TagKey *string `json:"tagKey" tf:"tag_key"`
	// +optional
	TemplateIDList []string `json:"templateIDList,omitempty" tf:"template_id_list"`
}

type DynamicTagGroupStatus struct {
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

// DynamicTagGroupList is a list of DynamicTagGroups
type DynamicTagGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DynamicTagGroup CRD objects
	Items []DynamicTagGroup `json:"items,omitempty"`
}
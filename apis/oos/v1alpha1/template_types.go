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

type Template struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateSpec   `json:"spec,omitempty"`
	Status            TemplateStatus `json:"status,omitempty"`
}

type TemplateSpec struct {
	State *TemplateSpecResource `json:"state,omitempty" tf:"-"`

	Resource TemplateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TemplateSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoDeleteExecutions *bool   `json:"autoDeleteExecutions,omitempty" tf:"auto_delete_executions"`
	Content              *string `json:"content" tf:"content"`
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// +optional
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	HasTrigger *bool `json:"hasTrigger,omitempty" tf:"has_trigger"`
	// +optional
	ShareType *string `json:"shareType,omitempty" tf:"share_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TemplateFormat *string `json:"templateFormat,omitempty" tf:"template_format"`
	// +optional
	TemplateID   *string `json:"templateID,omitempty" tf:"template_id"`
	TemplateName *string `json:"templateName" tf:"template_name"`
	// +optional
	TemplateType *string `json:"templateType,omitempty" tf:"template_type"`
	// +optional
	TemplateVersion *string `json:"templateVersion,omitempty" tf:"template_version"`
	// +optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by"`
	// +optional
	UpdatedDate *string `json:"updatedDate,omitempty" tf:"updated_date"`
	// +optional
	VersionName *string `json:"versionName,omitempty" tf:"version_name"`
}

type TemplateStatus struct {
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

// TemplateList is a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Template CRD objects
	Items []Template `json:"items,omitempty"`
}

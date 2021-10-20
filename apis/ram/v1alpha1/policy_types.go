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

type Policy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec,omitempty"`
	Status            PolicyStatus `json:"status,omitempty"`
}

type PolicySpecStatement struct {
	Action   []string `json:"action" tf:"action"`
	Effect   *string  `json:"effect" tf:"effect"`
	Resource []string `json:"resource" tf:"resource"`
}

type PolicySpec struct {
	State *PolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource PolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AttachmentCount *int64 `json:"attachmentCount,omitempty" tf:"attachment_count"`
	// +optional
	DefaultVersion *string `json:"defaultVersion,omitempty" tf:"default_version"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	// Deprecated
	Document *string `json:"document,omitempty" tf:"document"`
	// +optional
	Force *bool `json:"force,omitempty" tf:"force"`
	// +optional
	// Deprecated
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document"`
	// +optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name"`
	// +optional
	RotateStrategy *string `json:"rotateStrategy,omitempty" tf:"rotate_strategy"`
	// +optional
	// Deprecated
	Statement []PolicySpecStatement `json:"statement,omitempty" tf:"statement"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	// Deprecated
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VersionID *string `json:"versionID,omitempty" tf:"version_id"`
}

type PolicyStatus struct {
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

// PolicyList is a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Policy CRD objects
	Items []Policy `json:"items,omitempty"`
}

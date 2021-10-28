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

type Domain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

type DomainSpecCertInfos struct {
	// +optional
	CertName *string `json:"certName,omitempty" tf:"cert_name"`
	// +optional
	CertType *string `json:"certType,omitempty" tf:"cert_type"`
	// +optional
	SslPri *string `json:"-" sensitive:"true" tf:"ssl_pri"`
	// +optional
	SslProtocol *string `json:"sslProtocol,omitempty" tf:"ssl_protocol"`
	// +optional
	SslPub *string `json:"sslPub,omitempty" tf:"ssl_pub"`
}

type DomainSpecSources struct {
	Content *string `json:"content" tf:"content"`
	// +optional
	Enabled  *string `json:"enabled,omitempty" tf:"enabled"`
	Port     *int64  `json:"port" tf:"port"`
	Priority *string `json:"priority" tf:"priority"`
	Type     *string `json:"type" tf:"type"`
}

type DomainSpec struct {
	State *DomainSpecResource `json:"state,omitempty" tf:"-"`

	Resource DomainSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DomainSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BizName *string `json:"bizName,omitempty" tf:"biz_name"`
	// +optional
	CertInfos []DomainSpecCertInfos `json:"certInfos,omitempty" tf:"cert_infos"`
	// +optional
	CheckURL   *string `json:"checkURL,omitempty" tf:"check_url"`
	DomainName *string `json:"domainName" tf:"domain_name"`
	// +optional
	ForceSet *string `json:"forceSet,omitempty" tf:"force_set"`
	// +optional
	ResourceGroupID *string             `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	Sources         []DomainSpecSources `json:"sources" tf:"sources"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type DomainStatus struct {
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

// DomainList is a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Domain CRD objects
	Items []Domain `json:"items,omitempty"`
}

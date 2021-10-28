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

type SimpleOfficeSite struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SimpleOfficeSiteSpec   `json:"spec,omitempty"`
	Status            SimpleOfficeSiteStatus `json:"status,omitempty"`
}

type SimpleOfficeSiteSpec struct {
	State *SimpleOfficeSiteSpecResource `json:"state,omitempty" tf:"-"`

	Resource SimpleOfficeSiteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SimpleOfficeSiteSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Bandwidth *int64 `json:"bandwidth,omitempty" tf:"bandwidth"`
	// +optional
	CenID *string `json:"cenID,omitempty" tf:"cen_id"`
	// +optional
	CenOwnerID *string `json:"cenOwnerID,omitempty" tf:"cen_owner_id"`
	CidrBlock  *string `json:"cidrBlock" tf:"cidr_block"`
	// +optional
	DesktopAccessType *string `json:"desktopAccessType,omitempty" tf:"desktop_access_type"`
	// +optional
	EnableAdminAccess *bool `json:"enableAdminAccess,omitempty" tf:"enable_admin_access"`
	// +optional
	EnableCrossDesktopAccess *bool `json:"enableCrossDesktopAccess,omitempty" tf:"enable_cross_desktop_access"`
	// +optional
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access"`
	// +optional
	MfaEnabled *bool `json:"mfaEnabled,omitempty" tf:"mfa_enabled"`
	// +optional
	OfficeSiteName *string `json:"officeSiteName,omitempty" tf:"office_site_name"`
	// +optional
	SsoEnabled *bool `json:"ssoEnabled,omitempty" tf:"sso_enabled"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type SimpleOfficeSiteStatus struct {
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

// SimpleOfficeSiteList is a list of SimpleOfficeSites
type SimpleOfficeSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SimpleOfficeSite CRD objects
	Items []SimpleOfficeSite `json:"items,omitempty"`
}
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

type BandwidthPackage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BandwidthPackageSpec   `json:"spec,omitempty"`
	Status            BandwidthPackageStatus `json:"status,omitempty"`
}

type BandwidthPackageSpec struct {
	State *BandwidthPackageSpecResource `json:"state,omitempty" tf:"-"`

	Resource BandwidthPackageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BandwidthPackageSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoPay *bool `json:"autoPay,omitempty" tf:"auto_pay"`
	// +optional
	AutoUseCoupon *bool  `json:"autoUseCoupon,omitempty" tf:"auto_use_coupon"`
	Bandwidth     *int64 `json:"bandwidth" tf:"bandwidth"`
	// +optional
	BandwidthPackageName *string `json:"bandwidthPackageName,omitempty" tf:"bandwidth_package_name"`
	// +optional
	BandwidthType *string `json:"bandwidthType,omitempty" tf:"bandwidth_type"`
	// +optional
	BillingType *string `json:"billingType,omitempty" tf:"billing_type"`
	// +optional
	CbnGeographicRegionIda *string `json:"cbnGeographicRegionIda,omitempty" tf:"cbn_geographic_region_ida"`
	// +optional
	CbnGeographicRegionIdb *string `json:"cbnGeographicRegionIdb,omitempty" tf:"cbn_geographic_region_idb"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Duration *string `json:"duration,omitempty" tf:"duration"`
	// +optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type"`
	// +optional
	Ratio *int64 `json:"ratio,omitempty" tf:"ratio"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	Type   *string `json:"type" tf:"type"`
}

type BandwidthPackageStatus struct {
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

// BandwidthPackageList is a list of BandwidthPackages
type BandwidthPackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BandwidthPackage CRD objects
	Items []BandwidthPackage `json:"items,omitempty"`
}

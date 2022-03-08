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

type TrafficMirrorFilterIngressRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficMirrorFilterIngressRuleSpec   `json:"spec,omitempty"`
	Status            TrafficMirrorFilterIngressRuleStatus `json:"status,omitempty"`
}

type TrafficMirrorFilterIngressRuleSpec struct {
	State *TrafficMirrorFilterIngressRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource TrafficMirrorFilterIngressRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TrafficMirrorFilterIngressRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DestinationCIDRBlock *string `json:"destinationCIDRBlock" tf:"destination_cidr_block"`
	// +optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range"`
	// +optional
	DryRun          *bool   `json:"dryRun,omitempty" tf:"dry_run"`
	Priority        *int64  `json:"priority" tf:"priority"`
	Protocol        *string `json:"protocol" tf:"protocol"`
	RuleAction      *string `json:"ruleAction" tf:"rule_action"`
	SourceCIDRBlock *string `json:"sourceCIDRBlock" tf:"source_cidr_block"`
	// +optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range"`
	// +optional
	Status                *string `json:"status,omitempty" tf:"status"`
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterID" tf:"traffic_mirror_filter_id"`
	// +optional
	TrafficMirrorFilterIngressRuleID *string `json:"trafficMirrorFilterIngressRuleID,omitempty" tf:"traffic_mirror_filter_ingress_rule_id"`
}

type TrafficMirrorFilterIngressRuleStatus struct {
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

// TrafficMirrorFilterIngressRuleList is a list of TrafficMirrorFilterIngressRules
type TrafficMirrorFilterIngressRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TrafficMirrorFilterIngressRule CRD objects
	Items []TrafficMirrorFilterIngressRule `json:"items,omitempty"`
}

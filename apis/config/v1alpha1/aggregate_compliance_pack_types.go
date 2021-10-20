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

type AggregateCompliancePack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AggregateCompliancePackSpec   `json:"spec,omitempty"`
	Status            AggregateCompliancePackStatus `json:"status,omitempty"`
}

type AggregateCompliancePackSpecConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"parameterName" tf:"parameter_name"`
	ParameterValue *string `json:"parameterValue" tf:"parameter_value"`
}

type AggregateCompliancePackSpecConfigRules struct {
	ConfigRuleParameters  []AggregateCompliancePackSpecConfigRulesConfigRuleParameters `json:"configRuleParameters" tf:"config_rule_parameters"`
	ManagedRuleIdentifier *string                                                      `json:"managedRuleIdentifier" tf:"managed_rule_identifier"`
}

type AggregateCompliancePackSpec struct {
	State *AggregateCompliancePackSpecResource `json:"state,omitempty" tf:"-"`

	Resource AggregateCompliancePackSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AggregateCompliancePackSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AggregateCompliancePackName *string                                  `json:"aggregateCompliancePackName" tf:"aggregate_compliance_pack_name"`
	AggregatorID                *string                                  `json:"aggregatorID" tf:"aggregator_id"`
	CompliancePackTemplateID    *string                                  `json:"compliancePackTemplateID" tf:"compliance_pack_template_id"`
	ConfigRules                 []AggregateCompliancePackSpecConfigRules `json:"configRules" tf:"config_rules"`
	Description                 *string                                  `json:"description" tf:"description"`
	RiskLevel                   *int64                                   `json:"riskLevel" tf:"risk_level"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type AggregateCompliancePackStatus struct {
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

// AggregateCompliancePackList is a list of AggregateCompliancePacks
type AggregateCompliancePackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AggregateCompliancePack CRD objects
	Items []AggregateCompliancePack `json:"items,omitempty"`
}

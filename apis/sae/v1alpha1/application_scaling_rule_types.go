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

type ApplicationScalingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationScalingRuleSpec   `json:"spec,omitempty"`
	Status            ApplicationScalingRuleStatus `json:"status,omitempty"`
}

type ApplicationScalingRuleSpecScalingRuleMetricMetrics struct {
	// +optional
	MetricTargetAverageUtilization *int64 `json:"metricTargetAverageUtilization,omitempty" tf:"metric_target_average_utilization"`
	// +optional
	MetricType *string `json:"metricType,omitempty" tf:"metric_type"`
}

type ApplicationScalingRuleSpecScalingRuleMetricScaleDownRules struct {
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	StabilizationWindowSeconds *int64 `json:"stabilizationWindowSeconds,omitempty" tf:"stabilization_window_seconds"`
	// +optional
	Step *int64 `json:"step,omitempty" tf:"step"`
}

type ApplicationScalingRuleSpecScalingRuleMetricScaleUpRules struct {
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	StabilizationWindowSeconds *int64 `json:"stabilizationWindowSeconds,omitempty" tf:"stabilization_window_seconds"`
	// +optional
	Step *int64 `json:"step,omitempty" tf:"step"`
}

type ApplicationScalingRuleSpecScalingRuleMetric struct {
	// +optional
	MaxReplicas *int64 `json:"maxReplicas,omitempty" tf:"max_replicas"`
	// +optional
	Metrics []ApplicationScalingRuleSpecScalingRuleMetricMetrics `json:"metrics,omitempty" tf:"metrics"`
	// +optional
	MinReplicas *int64 `json:"minReplicas,omitempty" tf:"min_replicas"`
	// +optional
	ScaleDownRules *ApplicationScalingRuleSpecScalingRuleMetricScaleDownRules `json:"scaleDownRules,omitempty" tf:"scale_down_rules"`
	// +optional
	ScaleUpRules *ApplicationScalingRuleSpecScalingRuleMetricScaleUpRules `json:"scaleUpRules,omitempty" tf:"scale_up_rules"`
}

type ApplicationScalingRuleSpecScalingRuleTimerSchedules struct {
	// +optional
	AtTime *string `json:"atTime,omitempty" tf:"at_time"`
	// +optional
	MaxReplicas *int64 `json:"maxReplicas,omitempty" tf:"max_replicas"`
	// +optional
	MinReplicas *int64 `json:"minReplicas,omitempty" tf:"min_replicas"`
	// +optional
	TargetReplicas *int64 `json:"targetReplicas,omitempty" tf:"target_replicas"`
}

type ApplicationScalingRuleSpecScalingRuleTimer struct {
	// +optional
	BeginDate *string `json:"beginDate,omitempty" tf:"begin_date"`
	// +optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date"`
	// +optional
	Period *string `json:"period,omitempty" tf:"period"`
	// +optional
	Schedules []ApplicationScalingRuleSpecScalingRuleTimerSchedules `json:"schedules,omitempty" tf:"schedules"`
}

type ApplicationScalingRuleSpec struct {
	State *ApplicationScalingRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApplicationScalingRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApplicationScalingRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AppID *string `json:"appID" tf:"app_id"`
	// +optional
	MinReadyInstanceRatio *int64 `json:"minReadyInstanceRatio,omitempty" tf:"min_ready_instance_ratio"`
	// +optional
	MinReadyInstances *int64 `json:"minReadyInstances,omitempty" tf:"min_ready_instances"`
	// +optional
	ScalingRuleEnable *bool `json:"scalingRuleEnable,omitempty" tf:"scaling_rule_enable"`
	// +optional
	ScalingRuleMetric []ApplicationScalingRuleSpecScalingRuleMetric `json:"scalingRuleMetric,omitempty" tf:"scaling_rule_metric"`
	ScalingRuleName   *string                                       `json:"scalingRuleName" tf:"scaling_rule_name"`
	// +optional
	ScalingRuleTimer []ApplicationScalingRuleSpecScalingRuleTimer `json:"scalingRuleTimer,omitempty" tf:"scaling_rule_timer"`
	ScalingRuleType  *string                                      `json:"scalingRuleType" tf:"scaling_rule_type"`
}

type ApplicationScalingRuleStatus struct {
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

// ApplicationScalingRuleList is a list of ApplicationScalingRules
type ApplicationScalingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationScalingRule CRD objects
	Items []ApplicationScalingRule `json:"items,omitempty"`
}

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

type EnterpriseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnterpriseInstanceSpec   `json:"spec,omitempty"`
	Status            EnterpriseInstanceStatus `json:"status,omitempty"`
}

type EnterpriseInstanceSpec struct {
	State *EnterpriseInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource EnterpriseInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EnterpriseInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DataLinkName     *string `json:"dataLinkName,omitempty" tf:"data_link_name"`
	DatabasePassword *string `json:"-" sensitive:"true" tf:"database_password"`
	DatabaseUser     *string `json:"databaseUser" tf:"database_user"`
	// +optional
	DbaID *string `json:"dbaID,omitempty" tf:"dba_id"`
	// +optional
	DbaNickName *string `json:"dbaNickName,omitempty" tf:"dba_nick_name"`
	DbaUid      *int64  `json:"dbaUid" tf:"dba_uid"`
	// +optional
	DdlOnline *int64 `json:"ddlOnline,omitempty" tf:"ddl_online"`
	// +optional
	EcsInstanceID *string `json:"ecsInstanceID,omitempty" tf:"ecs_instance_id"`
	// +optional
	EcsRegion     *string `json:"ecsRegion,omitempty" tf:"ecs_region"`
	EnvType       *string `json:"envType" tf:"env_type"`
	ExportTimeout *int64  `json:"exportTimeout" tf:"export_timeout"`
	Host          *string `json:"host" tf:"host"`
	// +optional
	// Deprecated
	InstanceAlias *string `json:"instanceAlias,omitempty" tf:"instance_alias"`
	// +optional
	InstanceID *string `json:"instanceID,omitempty" tf:"instance_id"`
	// +optional
	InstanceName   *string `json:"instanceName,omitempty" tf:"instance_name"`
	InstanceSource *string `json:"instanceSource" tf:"instance_source"`
	InstanceType   *string `json:"instanceType" tf:"instance_type"`
	NetworkType    *string `json:"networkType" tf:"network_type"`
	Port           *int64  `json:"port" tf:"port"`
	QueryTimeout   *int64  `json:"queryTimeout" tf:"query_timeout"`
	SafeRule       *string `json:"safeRule" tf:"safe_rule"`
	// +optional
	SafeRuleID *string `json:"safeRuleID,omitempty" tf:"safe_rule_id"`
	// +optional
	Sid *string `json:"sid,omitempty" tf:"sid"`
	// +optional
	SkipTest *bool `json:"skipTest,omitempty" tf:"skip_test"`
	// +optional
	// Deprecated
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tid *int64 `json:"tid,omitempty" tf:"tid"`
	// +optional
	UseDsql *int64 `json:"useDsql,omitempty" tf:"use_dsql"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type EnterpriseInstanceStatus struct {
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

// EnterpriseInstanceList is a list of EnterpriseInstances
type EnterpriseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EnterpriseInstance CRD objects
	Items []EnterpriseInstance `json:"items,omitempty"`
}

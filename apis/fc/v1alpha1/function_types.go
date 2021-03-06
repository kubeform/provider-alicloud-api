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

type Function struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec,omitempty"`
	Status            FunctionStatus `json:"status,omitempty"`
}

type FunctionSpecCustomContainerConfig struct {
	// +optional
	Args *string `json:"args,omitempty" tf:"args"`
	// +optional
	Command *string `json:"command,omitempty" tf:"command"`
	Image   *string `json:"image" tf:"image"`
}

type FunctionSpec struct {
	State *FunctionSpecResource `json:"state,omitempty" tf:"-"`

	Resource FunctionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FunctionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CaPort *int64 `json:"caPort,omitempty" tf:"ca_port"`
	// +optional
	CodeChecksum *string `json:"codeChecksum,omitempty" tf:"code_checksum"`
	// +optional
	CustomContainerConfig *FunctionSpecCustomContainerConfig `json:"customContainerConfig,omitempty" tf:"custom_container_config"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`
	// +optional
	Filename *string `json:"filename,omitempty" tf:"filename"`
	// +optional
	FunctionID *string `json:"functionID,omitempty" tf:"function_id"`
	Handler    *string `json:"handler" tf:"handler"`
	// +optional
	InitializationTimeout *int64 `json:"initializationTimeout,omitempty" tf:"initialization_timeout"`
	// +optional
	Initializer *string `json:"initializer,omitempty" tf:"initializer"`
	// +optional
	InstanceConcurrency *int64 `json:"instanceConcurrency,omitempty" tf:"instance_concurrency"`
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified"`
	// +optional
	MemorySize *int64 `json:"memorySize,omitempty" tf:"memory_size"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	OssBucket *string `json:"ossBucket,omitempty" tf:"oss_bucket"`
	// +optional
	OssKey  *string `json:"ossKey,omitempty" tf:"oss_key"`
	Runtime *string `json:"runtime" tf:"runtime"`
	Service *string `json:"service" tf:"service"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
}

type FunctionStatus struct {
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

// FunctionList is a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Function CRD objects
	Items []Function `json:"items,omitempty"`
}

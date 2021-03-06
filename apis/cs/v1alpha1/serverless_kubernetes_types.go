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

type ServerlessKubernetes struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerlessKubernetesSpec   `json:"spec,omitempty"`
	Status            ServerlessKubernetesStatus `json:"status,omitempty"`
}

type ServerlessKubernetesSpecAddons struct {
	// +optional
	Config *string `json:"config,omitempty" tf:"config"`
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ServerlessKubernetesSpec struct {
	State *ServerlessKubernetesSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServerlessKubernetesSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServerlessKubernetesSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Addons []ServerlessKubernetesSpecAddons `json:"addons,omitempty" tf:"addons"`
	// +optional
	ClientCert *string `json:"clientCert,omitempty" tf:"client_cert"`
	// +optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key"`
	// +optional
	ClusterCaCert *string `json:"clusterCaCert,omitempty" tf:"cluster_ca_cert"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	EndpointPublicAccessEnabled *bool `json:"endpointPublicAccessEnabled,omitempty" tf:"endpoint_public_access_enabled"`
	// +optional
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update"`
	// +optional
	KubeConfig *string `json:"kubeConfig,omitempty" tf:"kube_config"`
	// +optional
	LoadBalancerSpec *string `json:"loadBalancerSpec,omitempty" tf:"load_balancer_spec"`
	// +optional
	LoggingType *string `json:"loggingType,omitempty" tf:"logging_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	NewNATGateway *bool `json:"newNATGateway,omitempty" tf:"new_nat_gateway"`
	// +optional
	// Deprecated
	PrivateZone *bool `json:"privateZone,omitempty" tf:"private_zone"`
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// +optional
	RetainResources []string `json:"retainResources,omitempty" tf:"retain_resources"`
	// +optional
	SecurityGroupID *string `json:"securityGroupID,omitempty" tf:"security_group_id"`
	// +optional
	ServiceCIDR *string `json:"serviceCIDR,omitempty" tf:"service_cidr"`
	// +optional
	ServiceDiscoveryTypes []string `json:"serviceDiscoveryTypes,omitempty" tf:"service_discovery_types"`
	// +optional
	SlsProjectName *string `json:"slsProjectName,omitempty" tf:"sls_project_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	VpcID   *string `json:"vpcID" tf:"vpc_id"`
	// +optional
	// Deprecated
	VswitchID *string `json:"vswitchID,omitempty" tf:"vswitch_id"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	VswitchIDS []string `json:"vswitchIDS,omitempty" tf:"vswitch_ids"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type ServerlessKubernetesStatus struct {
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

// ServerlessKubernetesList is a list of ServerlessKubernetess
type ServerlessKubernetesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServerlessKubernetes CRD objects
	Items []ServerlessKubernetes `json:"items,omitempty"`
}

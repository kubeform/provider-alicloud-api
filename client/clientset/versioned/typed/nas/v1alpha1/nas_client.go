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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/nas/v1alpha1"
	"kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type NasV1alpha1Interface interface {
	RESTClient() rest.Interface
	AccessGroupsGetter
	AccessRulesGetter
	AutoSnapshotPoliciesGetter
	DataFlowsGetter
	FileSystemsGetter
	FilesetsGetter
	LifecyclePoliciesGetter
	MountTargetsGetter
	RecycleBinsGetter
	SnapshotsGetter
}

// NasV1alpha1Client is used to interact with features provided by the nas.alicloud.kubeform.com group.
type NasV1alpha1Client struct {
	restClient rest.Interface
}

func (c *NasV1alpha1Client) AccessGroups(namespace string) AccessGroupInterface {
	return newAccessGroups(c, namespace)
}

func (c *NasV1alpha1Client) AccessRules(namespace string) AccessRuleInterface {
	return newAccessRules(c, namespace)
}

func (c *NasV1alpha1Client) AutoSnapshotPolicies(namespace string) AutoSnapshotPolicyInterface {
	return newAutoSnapshotPolicies(c, namespace)
}

func (c *NasV1alpha1Client) DataFlows(namespace string) DataFlowInterface {
	return newDataFlows(c, namespace)
}

func (c *NasV1alpha1Client) FileSystems(namespace string) FileSystemInterface {
	return newFileSystems(c, namespace)
}

func (c *NasV1alpha1Client) Filesets(namespace string) FilesetInterface {
	return newFilesets(c, namespace)
}

func (c *NasV1alpha1Client) LifecyclePolicies(namespace string) LifecyclePolicyInterface {
	return newLifecyclePolicies(c, namespace)
}

func (c *NasV1alpha1Client) MountTargets(namespace string) MountTargetInterface {
	return newMountTargets(c, namespace)
}

func (c *NasV1alpha1Client) RecycleBins(namespace string) RecycleBinInterface {
	return newRecycleBins(c, namespace)
}

func (c *NasV1alpha1Client) Snapshots(namespace string) SnapshotInterface {
	return newSnapshots(c, namespace)
}

// NewForConfig creates a new NasV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*NasV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NasV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new NasV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NasV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NasV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *NasV1alpha1Client {
	return &NasV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NasV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

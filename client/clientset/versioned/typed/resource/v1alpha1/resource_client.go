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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"
	"kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type ResourceV1alpha1Interface interface {
	RESTClient() rest.Interface
	ManagerAccountsGetter
	ManagerControlPoliciesGetter
	ManagerControlPolicyAttachmentsGetter
	ManagerFoldersGetter
	ManagerHandshakesGetter
	ManagerPoliciesGetter
	ManagerPolicyAttachmentsGetter
	ManagerPolicyVersionsGetter
	ManagerResourceDirectoriesGetter
	ManagerResourceGroupsGetter
	ManagerResourceSharesGetter
	ManagerRolesGetter
	ManagerSharedResourcesGetter
	ManagerSharedTargetsGetter
}

// ResourceV1alpha1Client is used to interact with features provided by the resource.alicloud.kubeform.com group.
type ResourceV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ResourceV1alpha1Client) ManagerAccounts(namespace string) ManagerAccountInterface {
	return newManagerAccounts(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerControlPolicies(namespace string) ManagerControlPolicyInterface {
	return newManagerControlPolicies(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerControlPolicyAttachments(namespace string) ManagerControlPolicyAttachmentInterface {
	return newManagerControlPolicyAttachments(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerFolders(namespace string) ManagerFolderInterface {
	return newManagerFolders(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerHandshakes(namespace string) ManagerHandshakeInterface {
	return newManagerHandshakes(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerPolicies(namespace string) ManagerPolicyInterface {
	return newManagerPolicies(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerPolicyAttachments(namespace string) ManagerPolicyAttachmentInterface {
	return newManagerPolicyAttachments(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerPolicyVersions(namespace string) ManagerPolicyVersionInterface {
	return newManagerPolicyVersions(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerResourceDirectories(namespace string) ManagerResourceDirectoryInterface {
	return newManagerResourceDirectories(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerResourceGroups(namespace string) ManagerResourceGroupInterface {
	return newManagerResourceGroups(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerResourceShares(namespace string) ManagerResourceShareInterface {
	return newManagerResourceShares(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerRoles(namespace string) ManagerRoleInterface {
	return newManagerRoles(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerSharedResources(namespace string) ManagerSharedResourceInterface {
	return newManagerSharedResources(c, namespace)
}

func (c *ResourceV1alpha1Client) ManagerSharedTargets(namespace string) ManagerSharedTargetInterface {
	return newManagerSharedTargets(c, namespace)
}

// NewForConfig creates a new ResourceV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ResourceV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ResourceV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ResourceV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ResourceV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ResourceV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ResourceV1alpha1Client {
	return &ResourceV1alpha1Client{c}
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
func (c *ResourceV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
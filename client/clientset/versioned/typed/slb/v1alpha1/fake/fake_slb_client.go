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

package fake

import (
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/slb/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSlbV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSlbV1alpha1) Acls(namespace string) v1alpha1.AclInterface {
	return &FakeAcls{c, namespace}
}

func (c *FakeSlbV1alpha1) Attachments(namespace string) v1alpha1.AttachmentInterface {
	return &FakeAttachments{c, namespace}
}

func (c *FakeSlbV1alpha1) BackendServers(namespace string) v1alpha1.BackendServerInterface {
	return &FakeBackendServers{c, namespace}
}

func (c *FakeSlbV1alpha1) CaCertificates(namespace string) v1alpha1.CaCertificateInterface {
	return &FakeCaCertificates{c, namespace}
}

func (c *FakeSlbV1alpha1) DomainExtensions(namespace string) v1alpha1.DomainExtensionInterface {
	return &FakeDomainExtensions{c, namespace}
}

func (c *FakeSlbV1alpha1) Listeners(namespace string) v1alpha1.ListenerInterface {
	return &FakeListeners{c, namespace}
}

func (c *FakeSlbV1alpha1) LoadBalancers(namespace string) v1alpha1.LoadBalancerInterface {
	return &FakeLoadBalancers{c, namespace}
}

func (c *FakeSlbV1alpha1) MasterSlaveServerGroups(namespace string) v1alpha1.MasterSlaveServerGroupInterface {
	return &FakeMasterSlaveServerGroups{c, namespace}
}

func (c *FakeSlbV1alpha1) Rules(namespace string) v1alpha1.RuleInterface {
	return &FakeRules{c, namespace}
}

func (c *FakeSlbV1alpha1) ServerCertificates(namespace string) v1alpha1.ServerCertificateInterface {
	return &FakeServerCertificates{c, namespace}
}

func (c *FakeSlbV1alpha1) ServerGroups(namespace string) v1alpha1.ServerGroupInterface {
	return &FakeServerGroups{c, namespace}
}

func (c *FakeSlbV1alpha1) Slbs(namespace string) v1alpha1.SlbInterface {
	return &FakeSlbs{c, namespace}
}

func (c *FakeSlbV1alpha1) TlsCipherPolicies(namespace string) v1alpha1.TlsCipherPolicyInterface {
	return &FakeTlsCipherPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSlbV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

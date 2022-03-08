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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/pvtz/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePvtzV1alpha1 struct {
	*testing.Fake
}

func (c *FakePvtzV1alpha1) Endpoints(namespace string) v1alpha1.EndpointInterface {
	return &FakeEndpoints{c, namespace}
}

func (c *FakePvtzV1alpha1) Rules(namespace string) v1alpha1.RuleInterface {
	return &FakeRules{c, namespace}
}

func (c *FakePvtzV1alpha1) RuleAttachments(namespace string) v1alpha1.RuleAttachmentInterface {
	return &FakeRuleAttachments{c, namespace}
}

func (c *FakePvtzV1alpha1) UserVpcAuthorizations(namespace string) v1alpha1.UserVpcAuthorizationInterface {
	return &FakeUserVpcAuthorizations{c, namespace}
}

func (c *FakePvtzV1alpha1) Zones(namespace string) v1alpha1.ZoneInterface {
	return &FakeZones{c, namespace}
}

func (c *FakePvtzV1alpha1) ZoneAttachments(namespace string) v1alpha1.ZoneAttachmentInterface {
	return &FakeZoneAttachments{c, namespace}
}

func (c *FakePvtzV1alpha1) ZoneRecords(namespace string) v1alpha1.ZoneRecordInterface {
	return &FakeZoneRecords{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePvtzV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

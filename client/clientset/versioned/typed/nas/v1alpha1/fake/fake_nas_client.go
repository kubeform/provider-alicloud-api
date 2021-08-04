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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/nas/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNasV1alpha1 struct {
	*testing.Fake
}

func (c *FakeNasV1alpha1) AccessGroups(namespace string) v1alpha1.AccessGroupInterface {
	return &FakeAccessGroups{c, namespace}
}

func (c *FakeNasV1alpha1) AccessRules(namespace string) v1alpha1.AccessRuleInterface {
	return &FakeAccessRules{c, namespace}
}

func (c *FakeNasV1alpha1) FileSystems(namespace string) v1alpha1.FileSystemInterface {
	return &FakeFileSystems{c, namespace}
}

func (c *FakeNasV1alpha1) MountTargets(namespace string) v1alpha1.MountTargetInterface {
	return &FakeMountTargets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNasV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
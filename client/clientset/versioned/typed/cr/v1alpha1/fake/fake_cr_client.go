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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/cr/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCrV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCrV1alpha1) EeInstances(namespace string) v1alpha1.EeInstanceInterface {
	return &FakeEeInstances{c, namespace}
}

func (c *FakeCrV1alpha1) EeNamespaces(namespace string) v1alpha1.EeNamespaceInterface {
	return &FakeEeNamespaces{c, namespace}
}

func (c *FakeCrV1alpha1) EeRepos(namespace string) v1alpha1.EeRepoInterface {
	return &FakeEeRepos{c, namespace}
}

func (c *FakeCrV1alpha1) EeSyncRules(namespace string) v1alpha1.EeSyncRuleInterface {
	return &FakeEeSyncRules{c, namespace}
}

func (c *FakeCrV1alpha1) Namespaces(namespace string) v1alpha1.NamespaceInterface {
	return &FakeNamespaces{c, namespace}
}

func (c *FakeCrV1alpha1) Repos(namespace string) v1alpha1.RepoInterface {
	return &FakeRepos{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

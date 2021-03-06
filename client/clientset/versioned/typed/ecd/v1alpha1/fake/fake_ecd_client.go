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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/ecd/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEcdV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEcdV1alpha1) Commands(namespace string) v1alpha1.CommandInterface {
	return &FakeCommands{c, namespace}
}

func (c *FakeEcdV1alpha1) Desktops(namespace string) v1alpha1.DesktopInterface {
	return &FakeDesktops{c, namespace}
}

func (c *FakeEcdV1alpha1) Images(namespace string) v1alpha1.ImageInterface {
	return &FakeImages{c, namespace}
}

func (c *FakeEcdV1alpha1) NasFileSystems(namespace string) v1alpha1.NasFileSystemInterface {
	return &FakeNasFileSystems{c, namespace}
}

func (c *FakeEcdV1alpha1) NetworkPackages(namespace string) v1alpha1.NetworkPackageInterface {
	return &FakeNetworkPackages{c, namespace}
}

func (c *FakeEcdV1alpha1) PolicyGroups(namespace string) v1alpha1.PolicyGroupInterface {
	return &FakePolicyGroups{c, namespace}
}

func (c *FakeEcdV1alpha1) SimpleOfficeSites(namespace string) v1alpha1.SimpleOfficeSiteInterface {
	return &FakeSimpleOfficeSites{c, namespace}
}

func (c *FakeEcdV1alpha1) Users(namespace string) v1alpha1.UserInterface {
	return &FakeUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEcdV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

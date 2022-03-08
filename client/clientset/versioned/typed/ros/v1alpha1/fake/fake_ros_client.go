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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/ros/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRosV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRosV1alpha1) ChangeSets(namespace string) v1alpha1.ChangeSetInterface {
	return &FakeChangeSets{c, namespace}
}

func (c *FakeRosV1alpha1) Stacks(namespace string) v1alpha1.StackInterface {
	return &FakeStacks{c, namespace}
}

func (c *FakeRosV1alpha1) StackGroups(namespace string) v1alpha1.StackGroupInterface {
	return &FakeStackGroups{c, namespace}
}

func (c *FakeRosV1alpha1) StackInstances(namespace string) v1alpha1.StackInstanceInterface {
	return &FakeStackInstances{c, namespace}
}

func (c *FakeRosV1alpha1) Templates(namespace string) v1alpha1.TemplateInterface {
	return &FakeTemplates{c, namespace}
}

func (c *FakeRosV1alpha1) TemplateScratches(namespace string) v1alpha1.TemplateScratchInterface {
	return &FakeTemplateScratches{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRosV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

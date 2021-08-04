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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/cs/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCsV1alpha1) Applications(namespace string) v1alpha1.ApplicationInterface {
	return &FakeApplications{c, namespace}
}

func (c *FakeCsV1alpha1) AutoscalingConfigs(namespace string) v1alpha1.AutoscalingConfigInterface {
	return &FakeAutoscalingConfigs{c, namespace}
}

func (c *FakeCsV1alpha1) EdgeKuberneteses(namespace string) v1alpha1.EdgeKubernetesInterface {
	return &FakeEdgeKuberneteses{c, namespace}
}

func (c *FakeCsV1alpha1) Kuberneteses(namespace string) v1alpha1.KubernetesInterface {
	return &FakeKuberneteses{c, namespace}
}

func (c *FakeCsV1alpha1) KubernetesAutoscalers(namespace string) v1alpha1.KubernetesAutoscalerInterface {
	return &FakeKubernetesAutoscalers{c, namespace}
}

func (c *FakeCsV1alpha1) KubernetesNodePools(namespace string) v1alpha1.KubernetesNodePoolInterface {
	return &FakeKubernetesNodePools{c, namespace}
}

func (c *FakeCsV1alpha1) KubernetesPermissionses(namespace string) v1alpha1.KubernetesPermissionsInterface {
	return &FakeKubernetesPermissionses{c, namespace}
}

func (c *FakeCsV1alpha1) ManagedKuberneteses(namespace string) v1alpha1.ManagedKubernetesInterface {
	return &FakeManagedKuberneteses{c, namespace}
}

func (c *FakeCsV1alpha1) ServerlessKuberneteses(namespace string) v1alpha1.ServerlessKubernetesInterface {
	return &FakeServerlessKuberneteses{c, namespace}
}

func (c *FakeCsV1alpha1) Swarms(namespace string) v1alpha1.SwarmInterface {
	return &FakeSwarms{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
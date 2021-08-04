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
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/clientset/versioned/typed/log/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLogV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLogV1alpha1) Alerts(namespace string) v1alpha1.AlertInterface {
	return &FakeAlerts{c, namespace}
}

func (c *FakeLogV1alpha1) Audits(namespace string) v1alpha1.AuditInterface {
	return &FakeAudits{c, namespace}
}

func (c *FakeLogV1alpha1) Dashboards(namespace string) v1alpha1.DashboardInterface {
	return &FakeDashboards{c, namespace}
}

func (c *FakeLogV1alpha1) Etls(namespace string) v1alpha1.EtlInterface {
	return &FakeEtls{c, namespace}
}

func (c *FakeLogV1alpha1) MachineGroups(namespace string) v1alpha1.MachineGroupInterface {
	return &FakeMachineGroups{c, namespace}
}

func (c *FakeLogV1alpha1) OssShippers(namespace string) v1alpha1.OssShipperInterface {
	return &FakeOssShippers{c, namespace}
}

func (c *FakeLogV1alpha1) Projects(namespace string) v1alpha1.ProjectInterface {
	return &FakeProjects{c, namespace}
}

func (c *FakeLogV1alpha1) Stores(namespace string) v1alpha1.StoreInterface {
	return &FakeStores{c, namespace}
}

func (c *FakeLogV1alpha1) StoreIndexes(namespace string) v1alpha1.StoreIndexInterface {
	return &FakeStoreIndexes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLogV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
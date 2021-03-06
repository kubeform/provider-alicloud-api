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
	"context"

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDedicatedHostClusters implements DedicatedHostClusterInterface
type FakeDedicatedHostClusters struct {
	Fake *FakeEcsV1alpha1
	ns   string
}

var dedicatedhostclustersResource = schema.GroupVersionResource{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "dedicatedhostclusters"}

var dedicatedhostclustersKind = schema.GroupVersionKind{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "DedicatedHostCluster"}

// Get takes name of the dedicatedHostCluster, and returns the corresponding dedicatedHostCluster object, and an error if there is any.
func (c *FakeDedicatedHostClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DedicatedHostCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dedicatedhostclustersResource, c.ns, name), &v1alpha1.DedicatedHostCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostCluster), err
}

// List takes label and field selectors, and returns the list of DedicatedHostClusters that match those selectors.
func (c *FakeDedicatedHostClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DedicatedHostClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dedicatedhostclustersResource, dedicatedhostclustersKind, c.ns, opts), &v1alpha1.DedicatedHostClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DedicatedHostClusterList{ListMeta: obj.(*v1alpha1.DedicatedHostClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DedicatedHostClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dedicatedHostClusters.
func (c *FakeDedicatedHostClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dedicatedhostclustersResource, c.ns, opts))

}

// Create takes the representation of a dedicatedHostCluster and creates it.  Returns the server's representation of the dedicatedHostCluster, and an error, if there is any.
func (c *FakeDedicatedHostClusters) Create(ctx context.Context, dedicatedHostCluster *v1alpha1.DedicatedHostCluster, opts v1.CreateOptions) (result *v1alpha1.DedicatedHostCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dedicatedhostclustersResource, c.ns, dedicatedHostCluster), &v1alpha1.DedicatedHostCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostCluster), err
}

// Update takes the representation of a dedicatedHostCluster and updates it. Returns the server's representation of the dedicatedHostCluster, and an error, if there is any.
func (c *FakeDedicatedHostClusters) Update(ctx context.Context, dedicatedHostCluster *v1alpha1.DedicatedHostCluster, opts v1.UpdateOptions) (result *v1alpha1.DedicatedHostCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dedicatedhostclustersResource, c.ns, dedicatedHostCluster), &v1alpha1.DedicatedHostCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDedicatedHostClusters) UpdateStatus(ctx context.Context, dedicatedHostCluster *v1alpha1.DedicatedHostCluster, opts v1.UpdateOptions) (*v1alpha1.DedicatedHostCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dedicatedhostclustersResource, "status", c.ns, dedicatedHostCluster), &v1alpha1.DedicatedHostCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostCluster), err
}

// Delete takes name of the dedicatedHostCluster and deletes it. Returns an error if one occurs.
func (c *FakeDedicatedHostClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dedicatedhostclustersResource, c.ns, name), &v1alpha1.DedicatedHostCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDedicatedHostClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dedicatedhostclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DedicatedHostClusterList{})
	return err
}

// Patch applies the patch and returns the patched dedicatedHostCluster.
func (c *FakeDedicatedHostClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DedicatedHostCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dedicatedhostclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DedicatedHostCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostCluster), err
}

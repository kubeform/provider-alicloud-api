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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/amqp/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualHosts implements VirtualHostInterface
type FakeVirtualHosts struct {
	Fake *FakeAmqpV1alpha1
	ns   string
}

var virtualhostsResource = schema.GroupVersionResource{Group: "amqp.alicloud.kubeform.com", Version: "v1alpha1", Resource: "virtualhosts"}

var virtualhostsKind = schema.GroupVersionKind{Group: "amqp.alicloud.kubeform.com", Version: "v1alpha1", Kind: "VirtualHost"}

// Get takes name of the virtualHost, and returns the corresponding virtualHost object, and an error if there is any.
func (c *FakeVirtualHosts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualHost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualhostsResource, c.ns, name), &v1alpha1.VirtualHost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualHost), err
}

// List takes label and field selectors, and returns the list of VirtualHosts that match those selectors.
func (c *FakeVirtualHosts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualHostList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualhostsResource, virtualhostsKind, c.ns, opts), &v1alpha1.VirtualHostList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualHostList{ListMeta: obj.(*v1alpha1.VirtualHostList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualHostList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualHosts.
func (c *FakeVirtualHosts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualhostsResource, c.ns, opts))

}

// Create takes the representation of a virtualHost and creates it.  Returns the server's representation of the virtualHost, and an error, if there is any.
func (c *FakeVirtualHosts) Create(ctx context.Context, virtualHost *v1alpha1.VirtualHost, opts v1.CreateOptions) (result *v1alpha1.VirtualHost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualhostsResource, c.ns, virtualHost), &v1alpha1.VirtualHost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualHost), err
}

// Update takes the representation of a virtualHost and updates it. Returns the server's representation of the virtualHost, and an error, if there is any.
func (c *FakeVirtualHosts) Update(ctx context.Context, virtualHost *v1alpha1.VirtualHost, opts v1.UpdateOptions) (result *v1alpha1.VirtualHost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualhostsResource, c.ns, virtualHost), &v1alpha1.VirtualHost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualHost), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualHosts) UpdateStatus(ctx context.Context, virtualHost *v1alpha1.VirtualHost, opts v1.UpdateOptions) (*v1alpha1.VirtualHost, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualhostsResource, "status", c.ns, virtualHost), &v1alpha1.VirtualHost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualHost), err
}

// Delete takes name of the virtualHost and deletes it. Returns an error if one occurs.
func (c *FakeVirtualHosts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualhostsResource, c.ns, name), &v1alpha1.VirtualHost{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualHosts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualhostsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualHostList{})
	return err
}

// Patch applies the patch and returns the patched virtualHost.
func (c *FakeVirtualHosts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualHost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualHost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualHost), err
}

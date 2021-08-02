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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vpn/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConnections implements ConnectionInterface
type FakeConnections struct {
	Fake *FakeVpnV1alpha1
	ns   string
}

var connectionsResource = schema.GroupVersionResource{Group: "vpn.alicloud.kubeform.com", Version: "v1alpha1", Resource: "connections"}

var connectionsKind = schema.GroupVersionKind{Group: "vpn.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Connection"}

// Get takes name of the connection, and returns the corresponding connection object, and an error if there is any.
func (c *FakeConnections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Connection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(connectionsResource, c.ns, name), &v1alpha1.Connection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Connection), err
}

// List takes label and field selectors, and returns the list of Connections that match those selectors.
func (c *FakeConnections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(connectionsResource, connectionsKind, c.ns, opts), &v1alpha1.ConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConnectionList{ListMeta: obj.(*v1alpha1.ConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested connections.
func (c *FakeConnections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(connectionsResource, c.ns, opts))

}

// Create takes the representation of a connection and creates it.  Returns the server's representation of the connection, and an error, if there is any.
func (c *FakeConnections) Create(ctx context.Context, connection *v1alpha1.Connection, opts v1.CreateOptions) (result *v1alpha1.Connection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(connectionsResource, c.ns, connection), &v1alpha1.Connection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Connection), err
}

// Update takes the representation of a connection and updates it. Returns the server's representation of the connection, and an error, if there is any.
func (c *FakeConnections) Update(ctx context.Context, connection *v1alpha1.Connection, opts v1.UpdateOptions) (result *v1alpha1.Connection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(connectionsResource, c.ns, connection), &v1alpha1.Connection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Connection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConnections) UpdateStatus(ctx context.Context, connection *v1alpha1.Connection, opts v1.UpdateOptions) (*v1alpha1.Connection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(connectionsResource, "status", c.ns, connection), &v1alpha1.Connection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Connection), err
}

// Delete takes name of the connection and deletes it. Returns an error if one occurs.
func (c *FakeConnections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(connectionsResource, c.ns, name), &v1alpha1.Connection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConnections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(connectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConnectionList{})
	return err
}

// Patch applies the patch and returns the patched connection.
func (c *FakeConnections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Connection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(connectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Connection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Connection), err
}

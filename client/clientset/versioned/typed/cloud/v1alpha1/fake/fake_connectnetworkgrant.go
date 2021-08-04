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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConnectNetworkGrants implements ConnectNetworkGrantInterface
type FakeConnectNetworkGrants struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var connectnetworkgrantsResource = schema.GroupVersionResource{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Resource: "connectnetworkgrants"}

var connectnetworkgrantsKind = schema.GroupVersionKind{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ConnectNetworkGrant"}

// Get takes name of the connectNetworkGrant, and returns the corresponding connectNetworkGrant object, and an error if there is any.
func (c *FakeConnectNetworkGrants) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConnectNetworkGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(connectnetworkgrantsResource, c.ns, name), &v1alpha1.ConnectNetworkGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectNetworkGrant), err
}

// List takes label and field selectors, and returns the list of ConnectNetworkGrants that match those selectors.
func (c *FakeConnectNetworkGrants) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConnectNetworkGrantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(connectnetworkgrantsResource, connectnetworkgrantsKind, c.ns, opts), &v1alpha1.ConnectNetworkGrantList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConnectNetworkGrantList{ListMeta: obj.(*v1alpha1.ConnectNetworkGrantList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConnectNetworkGrantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested connectNetworkGrants.
func (c *FakeConnectNetworkGrants) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(connectnetworkgrantsResource, c.ns, opts))

}

// Create takes the representation of a connectNetworkGrant and creates it.  Returns the server's representation of the connectNetworkGrant, and an error, if there is any.
func (c *FakeConnectNetworkGrants) Create(ctx context.Context, connectNetworkGrant *v1alpha1.ConnectNetworkGrant, opts v1.CreateOptions) (result *v1alpha1.ConnectNetworkGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(connectnetworkgrantsResource, c.ns, connectNetworkGrant), &v1alpha1.ConnectNetworkGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectNetworkGrant), err
}

// Update takes the representation of a connectNetworkGrant and updates it. Returns the server's representation of the connectNetworkGrant, and an error, if there is any.
func (c *FakeConnectNetworkGrants) Update(ctx context.Context, connectNetworkGrant *v1alpha1.ConnectNetworkGrant, opts v1.UpdateOptions) (result *v1alpha1.ConnectNetworkGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(connectnetworkgrantsResource, c.ns, connectNetworkGrant), &v1alpha1.ConnectNetworkGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectNetworkGrant), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConnectNetworkGrants) UpdateStatus(ctx context.Context, connectNetworkGrant *v1alpha1.ConnectNetworkGrant, opts v1.UpdateOptions) (*v1alpha1.ConnectNetworkGrant, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(connectnetworkgrantsResource, "status", c.ns, connectNetworkGrant), &v1alpha1.ConnectNetworkGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectNetworkGrant), err
}

// Delete takes name of the connectNetworkGrant and deletes it. Returns an error if one occurs.
func (c *FakeConnectNetworkGrants) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(connectnetworkgrantsResource, c.ns, name), &v1alpha1.ConnectNetworkGrant{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConnectNetworkGrants) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(connectnetworkgrantsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConnectNetworkGrantList{})
	return err
}

// Patch applies the patch and returns the patched connectNetworkGrant.
func (c *FakeConnectNetworkGrants) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConnectNetworkGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(connectnetworkgrantsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConnectNetworkGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectNetworkGrant), err
}
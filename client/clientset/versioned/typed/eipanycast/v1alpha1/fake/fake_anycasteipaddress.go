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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/eipanycast/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAnycastEipAddresses implements AnycastEipAddressInterface
type FakeAnycastEipAddresses struct {
	Fake *FakeEipanycastV1alpha1
	ns   string
}

var anycasteipaddressesResource = schema.GroupVersionResource{Group: "eipanycast.alicloud.kubeform.com", Version: "v1alpha1", Resource: "anycasteipaddresses"}

var anycasteipaddressesKind = schema.GroupVersionKind{Group: "eipanycast.alicloud.kubeform.com", Version: "v1alpha1", Kind: "AnycastEipAddress"}

// Get takes name of the anycastEipAddress, and returns the corresponding anycastEipAddress object, and an error if there is any.
func (c *FakeAnycastEipAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnycastEipAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(anycasteipaddressesResource, c.ns, name), &v1alpha1.AnycastEipAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnycastEipAddress), err
}

// List takes label and field selectors, and returns the list of AnycastEipAddresses that match those selectors.
func (c *FakeAnycastEipAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnycastEipAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(anycasteipaddressesResource, anycasteipaddressesKind, c.ns, opts), &v1alpha1.AnycastEipAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AnycastEipAddressList{ListMeta: obj.(*v1alpha1.AnycastEipAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.AnycastEipAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested anycastEipAddresses.
func (c *FakeAnycastEipAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(anycasteipaddressesResource, c.ns, opts))

}

// Create takes the representation of a anycastEipAddress and creates it.  Returns the server's representation of the anycastEipAddress, and an error, if there is any.
func (c *FakeAnycastEipAddresses) Create(ctx context.Context, anycastEipAddress *v1alpha1.AnycastEipAddress, opts v1.CreateOptions) (result *v1alpha1.AnycastEipAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(anycasteipaddressesResource, c.ns, anycastEipAddress), &v1alpha1.AnycastEipAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnycastEipAddress), err
}

// Update takes the representation of a anycastEipAddress and updates it. Returns the server's representation of the anycastEipAddress, and an error, if there is any.
func (c *FakeAnycastEipAddresses) Update(ctx context.Context, anycastEipAddress *v1alpha1.AnycastEipAddress, opts v1.UpdateOptions) (result *v1alpha1.AnycastEipAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(anycasteipaddressesResource, c.ns, anycastEipAddress), &v1alpha1.AnycastEipAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnycastEipAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnycastEipAddresses) UpdateStatus(ctx context.Context, anycastEipAddress *v1alpha1.AnycastEipAddress, opts v1.UpdateOptions) (*v1alpha1.AnycastEipAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(anycasteipaddressesResource, "status", c.ns, anycastEipAddress), &v1alpha1.AnycastEipAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnycastEipAddress), err
}

// Delete takes name of the anycastEipAddress and deletes it. Returns an error if one occurs.
func (c *FakeAnycastEipAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(anycasteipaddressesResource, c.ns, name), &v1alpha1.AnycastEipAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnycastEipAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(anycasteipaddressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AnycastEipAddressList{})
	return err
}

// Patch applies the patch and returns the patched anycastEipAddress.
func (c *FakeAnycastEipAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnycastEipAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(anycasteipaddressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AnycastEipAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnycastEipAddress), err
}

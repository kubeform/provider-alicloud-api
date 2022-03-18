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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/mongodb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShardingNetworkPublicAddresses implements ShardingNetworkPublicAddressInterface
type FakeShardingNetworkPublicAddresses struct {
	Fake *FakeMongodbV1alpha1
	ns   string
}

var shardingnetworkpublicaddressesResource = schema.GroupVersionResource{Group: "mongodb.alicloud.kubeform.com", Version: "v1alpha1", Resource: "shardingnetworkpublicaddresses"}

var shardingnetworkpublicaddressesKind = schema.GroupVersionKind{Group: "mongodb.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ShardingNetworkPublicAddress"}

// Get takes name of the shardingNetworkPublicAddress, and returns the corresponding shardingNetworkPublicAddress object, and an error if there is any.
func (c *FakeShardingNetworkPublicAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShardingNetworkPublicAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shardingnetworkpublicaddressesResource, c.ns, name), &v1alpha1.ShardingNetworkPublicAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShardingNetworkPublicAddress), err
}

// List takes label and field selectors, and returns the list of ShardingNetworkPublicAddresses that match those selectors.
func (c *FakeShardingNetworkPublicAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShardingNetworkPublicAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shardingnetworkpublicaddressesResource, shardingnetworkpublicaddressesKind, c.ns, opts), &v1alpha1.ShardingNetworkPublicAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShardingNetworkPublicAddressList{ListMeta: obj.(*v1alpha1.ShardingNetworkPublicAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShardingNetworkPublicAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shardingNetworkPublicAddresses.
func (c *FakeShardingNetworkPublicAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shardingnetworkpublicaddressesResource, c.ns, opts))

}

// Create takes the representation of a shardingNetworkPublicAddress and creates it.  Returns the server's representation of the shardingNetworkPublicAddress, and an error, if there is any.
func (c *FakeShardingNetworkPublicAddresses) Create(ctx context.Context, shardingNetworkPublicAddress *v1alpha1.ShardingNetworkPublicAddress, opts v1.CreateOptions) (result *v1alpha1.ShardingNetworkPublicAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shardingnetworkpublicaddressesResource, c.ns, shardingNetworkPublicAddress), &v1alpha1.ShardingNetworkPublicAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShardingNetworkPublicAddress), err
}

// Update takes the representation of a shardingNetworkPublicAddress and updates it. Returns the server's representation of the shardingNetworkPublicAddress, and an error, if there is any.
func (c *FakeShardingNetworkPublicAddresses) Update(ctx context.Context, shardingNetworkPublicAddress *v1alpha1.ShardingNetworkPublicAddress, opts v1.UpdateOptions) (result *v1alpha1.ShardingNetworkPublicAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shardingnetworkpublicaddressesResource, c.ns, shardingNetworkPublicAddress), &v1alpha1.ShardingNetworkPublicAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShardingNetworkPublicAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShardingNetworkPublicAddresses) UpdateStatus(ctx context.Context, shardingNetworkPublicAddress *v1alpha1.ShardingNetworkPublicAddress, opts v1.UpdateOptions) (*v1alpha1.ShardingNetworkPublicAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(shardingnetworkpublicaddressesResource, "status", c.ns, shardingNetworkPublicAddress), &v1alpha1.ShardingNetworkPublicAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShardingNetworkPublicAddress), err
}

// Delete takes name of the shardingNetworkPublicAddress and deletes it. Returns an error if one occurs.
func (c *FakeShardingNetworkPublicAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shardingnetworkpublicaddressesResource, c.ns, name), &v1alpha1.ShardingNetworkPublicAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShardingNetworkPublicAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shardingnetworkpublicaddressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShardingNetworkPublicAddressList{})
	return err
}

// Patch applies the patch and returns the patched shardingNetworkPublicAddress.
func (c *FakeShardingNetworkPublicAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShardingNetworkPublicAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shardingnetworkpublicaddressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShardingNetworkPublicAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShardingNetworkPublicAddress), err
}
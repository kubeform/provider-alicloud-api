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

// FakeStorageGatewayGateways implements StorageGatewayGatewayInterface
type FakeStorageGatewayGateways struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var storagegatewaygatewaysResource = schema.GroupVersionResource{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Resource: "storagegatewaygateways"}

var storagegatewaygatewaysKind = schema.GroupVersionKind{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Kind: "StorageGatewayGateway"}

// Get takes name of the storageGatewayGateway, and returns the corresponding storageGatewayGateway object, and an error if there is any.
func (c *FakeStorageGatewayGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagegatewaygatewaysResource, c.ns, name), &v1alpha1.StorageGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageGatewayGateway), err
}

// List takes label and field selectors, and returns the list of StorageGatewayGateways that match those selectors.
func (c *FakeStorageGatewayGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageGatewayGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagegatewaygatewaysResource, storagegatewaygatewaysKind, c.ns, opts), &v1alpha1.StorageGatewayGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageGatewayGatewayList{ListMeta: obj.(*v1alpha1.StorageGatewayGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageGatewayGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageGatewayGateways.
func (c *FakeStorageGatewayGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagegatewaygatewaysResource, c.ns, opts))

}

// Create takes the representation of a storageGatewayGateway and creates it.  Returns the server's representation of the storageGatewayGateway, and an error, if there is any.
func (c *FakeStorageGatewayGateways) Create(ctx context.Context, storageGatewayGateway *v1alpha1.StorageGatewayGateway, opts v1.CreateOptions) (result *v1alpha1.StorageGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagegatewaygatewaysResource, c.ns, storageGatewayGateway), &v1alpha1.StorageGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageGatewayGateway), err
}

// Update takes the representation of a storageGatewayGateway and updates it. Returns the server's representation of the storageGatewayGateway, and an error, if there is any.
func (c *FakeStorageGatewayGateways) Update(ctx context.Context, storageGatewayGateway *v1alpha1.StorageGatewayGateway, opts v1.UpdateOptions) (result *v1alpha1.StorageGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagegatewaygatewaysResource, c.ns, storageGatewayGateway), &v1alpha1.StorageGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageGatewayGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageGatewayGateways) UpdateStatus(ctx context.Context, storageGatewayGateway *v1alpha1.StorageGatewayGateway, opts v1.UpdateOptions) (*v1alpha1.StorageGatewayGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagegatewaygatewaysResource, "status", c.ns, storageGatewayGateway), &v1alpha1.StorageGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageGatewayGateway), err
}

// Delete takes name of the storageGatewayGateway and deletes it. Returns an error if one occurs.
func (c *FakeStorageGatewayGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagegatewaygatewaysResource, c.ns, name), &v1alpha1.StorageGatewayGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageGatewayGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagegatewaygatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageGatewayGatewayList{})
	return err
}

// Patch applies the patch and returns the patched storageGatewayGateway.
func (c *FakeStorageGatewayGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagegatewaygatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageGatewayGateway), err
}

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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/log/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStores implements StoreInterface
type FakeStores struct {
	Fake *FakeLogV1alpha1
	ns   string
}

var storesResource = schema.GroupVersionResource{Group: "log.alicloud.kubeform.com", Version: "v1alpha1", Resource: "stores"}

var storesKind = schema.GroupVersionKind{Group: "log.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Store"}

// Get takes name of the store, and returns the corresponding store object, and an error if there is any.
func (c *FakeStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storesResource, c.ns, name), &v1alpha1.Store{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// List takes label and field selectors, and returns the list of Stores that match those selectors.
func (c *FakeStores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storesResource, storesKind, c.ns, opts), &v1alpha1.StoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoreList{ListMeta: obj.(*v1alpha1.StoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stores.
func (c *FakeStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storesResource, c.ns, opts))

}

// Create takes the representation of a store and creates it.  Returns the server's representation of the store, and an error, if there is any.
func (c *FakeStores) Create(ctx context.Context, store *v1alpha1.Store, opts v1.CreateOptions) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storesResource, c.ns, store), &v1alpha1.Store{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// Update takes the representation of a store and updates it. Returns the server's representation of the store, and an error, if there is any.
func (c *FakeStores) Update(ctx context.Context, store *v1alpha1.Store, opts v1.UpdateOptions) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storesResource, c.ns, store), &v1alpha1.Store{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStores) UpdateStatus(ctx context.Context, store *v1alpha1.Store, opts v1.UpdateOptions) (*v1alpha1.Store, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storesResource, "status", c.ns, store), &v1alpha1.Store{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// Delete takes name of the store and deletes it. Returns an error if one occurs.
func (c *FakeStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storesResource, c.ns, name), &v1alpha1.Store{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoreList{})
	return err
}

// Patch applies the patch and returns the patched store.
func (c *FakeStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Store{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

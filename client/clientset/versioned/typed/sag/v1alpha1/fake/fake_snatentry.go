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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/sag/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSnatEntries implements SnatEntryInterface
type FakeSnatEntries struct {
	Fake *FakeSagV1alpha1
	ns   string
}

var snatentriesResource = schema.GroupVersionResource{Group: "sag.alicloud.kubeform.com", Version: "v1alpha1", Resource: "snatentries"}

var snatentriesKind = schema.GroupVersionKind{Group: "sag.alicloud.kubeform.com", Version: "v1alpha1", Kind: "SnatEntry"}

// Get takes name of the snatEntry, and returns the corresponding snatEntry object, and an error if there is any.
func (c *FakeSnatEntries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SnatEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snatentriesResource, c.ns, name), &v1alpha1.SnatEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnatEntry), err
}

// List takes label and field selectors, and returns the list of SnatEntries that match those selectors.
func (c *FakeSnatEntries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SnatEntryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snatentriesResource, snatentriesKind, c.ns, opts), &v1alpha1.SnatEntryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SnatEntryList{ListMeta: obj.(*v1alpha1.SnatEntryList).ListMeta}
	for _, item := range obj.(*v1alpha1.SnatEntryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snatEntries.
func (c *FakeSnatEntries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snatentriesResource, c.ns, opts))

}

// Create takes the representation of a snatEntry and creates it.  Returns the server's representation of the snatEntry, and an error, if there is any.
func (c *FakeSnatEntries) Create(ctx context.Context, snatEntry *v1alpha1.SnatEntry, opts v1.CreateOptions) (result *v1alpha1.SnatEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snatentriesResource, c.ns, snatEntry), &v1alpha1.SnatEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnatEntry), err
}

// Update takes the representation of a snatEntry and updates it. Returns the server's representation of the snatEntry, and an error, if there is any.
func (c *FakeSnatEntries) Update(ctx context.Context, snatEntry *v1alpha1.SnatEntry, opts v1.UpdateOptions) (result *v1alpha1.SnatEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snatentriesResource, c.ns, snatEntry), &v1alpha1.SnatEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnatEntry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSnatEntries) UpdateStatus(ctx context.Context, snatEntry *v1alpha1.SnatEntry, opts v1.UpdateOptions) (*v1alpha1.SnatEntry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(snatentriesResource, "status", c.ns, snatEntry), &v1alpha1.SnatEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnatEntry), err
}

// Delete takes name of the snatEntry and deletes it. Returns an error if one occurs.
func (c *FakeSnatEntries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snatentriesResource, c.ns, name), &v1alpha1.SnatEntry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnatEntries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snatentriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SnatEntryList{})
	return err
}

// Patch applies the patch and returns the patched snatEntry.
func (c *FakeSnatEntries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SnatEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snatentriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SnatEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnatEntry), err
}

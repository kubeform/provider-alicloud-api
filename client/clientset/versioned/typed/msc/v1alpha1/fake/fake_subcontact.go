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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/msc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubContacts implements SubContactInterface
type FakeSubContacts struct {
	Fake *FakeMscV1alpha1
	ns   string
}

var subcontactsResource = schema.GroupVersionResource{Group: "msc.alicloud.kubeform.com", Version: "v1alpha1", Resource: "subcontacts"}

var subcontactsKind = schema.GroupVersionKind{Group: "msc.alicloud.kubeform.com", Version: "v1alpha1", Kind: "SubContact"}

// Get takes name of the subContact, and returns the corresponding subContact object, and an error if there is any.
func (c *FakeSubContacts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SubContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subcontactsResource, c.ns, name), &v1alpha1.SubContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubContact), err
}

// List takes label and field selectors, and returns the list of SubContacts that match those selectors.
func (c *FakeSubContacts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubContactList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subcontactsResource, subcontactsKind, c.ns, opts), &v1alpha1.SubContactList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubContactList{ListMeta: obj.(*v1alpha1.SubContactList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubContactList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subContacts.
func (c *FakeSubContacts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subcontactsResource, c.ns, opts))

}

// Create takes the representation of a subContact and creates it.  Returns the server's representation of the subContact, and an error, if there is any.
func (c *FakeSubContacts) Create(ctx context.Context, subContact *v1alpha1.SubContact, opts v1.CreateOptions) (result *v1alpha1.SubContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subcontactsResource, c.ns, subContact), &v1alpha1.SubContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubContact), err
}

// Update takes the representation of a subContact and updates it. Returns the server's representation of the subContact, and an error, if there is any.
func (c *FakeSubContacts) Update(ctx context.Context, subContact *v1alpha1.SubContact, opts v1.UpdateOptions) (result *v1alpha1.SubContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subcontactsResource, c.ns, subContact), &v1alpha1.SubContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubContact), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubContacts) UpdateStatus(ctx context.Context, subContact *v1alpha1.SubContact, opts v1.UpdateOptions) (*v1alpha1.SubContact, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subcontactsResource, "status", c.ns, subContact), &v1alpha1.SubContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubContact), err
}

// Delete takes name of the subContact and deletes it. Returns an error if one occurs.
func (c *FakeSubContacts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subcontactsResource, c.ns, name), &v1alpha1.SubContact{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubContacts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subcontactsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubContactList{})
	return err
}

// Patch applies the patch and returns the patched subContact.
func (c *FakeSubContacts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subcontactsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubContact), err
}
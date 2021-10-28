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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/arms/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlertContacts implements AlertContactInterface
type FakeAlertContacts struct {
	Fake *FakeArmsV1alpha1
	ns   string
}

var alertcontactsResource = schema.GroupVersionResource{Group: "arms.alicloud.kubeform.com", Version: "v1alpha1", Resource: "alertcontacts"}

var alertcontactsKind = schema.GroupVersionKind{Group: "arms.alicloud.kubeform.com", Version: "v1alpha1", Kind: "AlertContact"}

// Get takes name of the alertContact, and returns the corresponding alertContact object, and an error if there is any.
func (c *FakeAlertContacts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlertContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertcontactsResource, c.ns, name), &v1alpha1.AlertContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertContact), err
}

// List takes label and field selectors, and returns the list of AlertContacts that match those selectors.
func (c *FakeAlertContacts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlertContactList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertcontactsResource, alertcontactsKind, c.ns, opts), &v1alpha1.AlertContactList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlertContactList{ListMeta: obj.(*v1alpha1.AlertContactList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlertContactList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertContacts.
func (c *FakeAlertContacts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertcontactsResource, c.ns, opts))

}

// Create takes the representation of a alertContact and creates it.  Returns the server's representation of the alertContact, and an error, if there is any.
func (c *FakeAlertContacts) Create(ctx context.Context, alertContact *v1alpha1.AlertContact, opts v1.CreateOptions) (result *v1alpha1.AlertContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertcontactsResource, c.ns, alertContact), &v1alpha1.AlertContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertContact), err
}

// Update takes the representation of a alertContact and updates it. Returns the server's representation of the alertContact, and an error, if there is any.
func (c *FakeAlertContacts) Update(ctx context.Context, alertContact *v1alpha1.AlertContact, opts v1.UpdateOptions) (result *v1alpha1.AlertContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertcontactsResource, c.ns, alertContact), &v1alpha1.AlertContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertContact), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlertContacts) UpdateStatus(ctx context.Context, alertContact *v1alpha1.AlertContact, opts v1.UpdateOptions) (*v1alpha1.AlertContact, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alertcontactsResource, "status", c.ns, alertContact), &v1alpha1.AlertContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertContact), err
}

// Delete takes name of the alertContact and deletes it. Returns an error if one occurs.
func (c *FakeAlertContacts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alertcontactsResource, c.ns, name), &v1alpha1.AlertContact{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlertContacts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertcontactsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlertContactList{})
	return err
}

// Patch applies the patch and returns the patched alertContact.
func (c *FakeAlertContacts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlertContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertcontactsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlertContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertContact), err
}
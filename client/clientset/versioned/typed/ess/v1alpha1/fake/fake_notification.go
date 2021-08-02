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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ess/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNotifications implements NotificationInterface
type FakeNotifications struct {
	Fake *FakeEssV1alpha1
	ns   string
}

var notificationsResource = schema.GroupVersionResource{Group: "ess.alicloud.kubeform.com", Version: "v1alpha1", Resource: "notifications"}

var notificationsKind = schema.GroupVersionKind{Group: "ess.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Notification"}

// Get takes name of the notification, and returns the corresponding notification object, and an error if there is any.
func (c *FakeNotifications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Notification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(notificationsResource, c.ns, name), &v1alpha1.Notification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Notification), err
}

// List takes label and field selectors, and returns the list of Notifications that match those selectors.
func (c *FakeNotifications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(notificationsResource, notificationsKind, c.ns, opts), &v1alpha1.NotificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NotificationList{ListMeta: obj.(*v1alpha1.NotificationList).ListMeta}
	for _, item := range obj.(*v1alpha1.NotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested notifications.
func (c *FakeNotifications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(notificationsResource, c.ns, opts))

}

// Create takes the representation of a notification and creates it.  Returns the server's representation of the notification, and an error, if there is any.
func (c *FakeNotifications) Create(ctx context.Context, notification *v1alpha1.Notification, opts v1.CreateOptions) (result *v1alpha1.Notification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(notificationsResource, c.ns, notification), &v1alpha1.Notification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Notification), err
}

// Update takes the representation of a notification and updates it. Returns the server's representation of the notification, and an error, if there is any.
func (c *FakeNotifications) Update(ctx context.Context, notification *v1alpha1.Notification, opts v1.UpdateOptions) (result *v1alpha1.Notification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(notificationsResource, c.ns, notification), &v1alpha1.Notification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Notification), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNotifications) UpdateStatus(ctx context.Context, notification *v1alpha1.Notification, opts v1.UpdateOptions) (*v1alpha1.Notification, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(notificationsResource, "status", c.ns, notification), &v1alpha1.Notification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Notification), err
}

// Delete takes name of the notification and deletes it. Returns an error if one occurs.
func (c *FakeNotifications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(notificationsResource, c.ns, name), &v1alpha1.Notification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNotifications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(notificationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NotificationList{})
	return err
}

// Patch applies the patch and returns the patched notification.
func (c *FakeNotifications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Notification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(notificationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Notification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Notification), err
}

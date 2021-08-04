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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ram/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoleAttachments implements RoleAttachmentInterface
type FakeRoleAttachments struct {
	Fake *FakeRamV1alpha1
	ns   string
}

var roleattachmentsResource = schema.GroupVersionResource{Group: "ram.alicloud.kubeform.com", Version: "v1alpha1", Resource: "roleattachments"}

var roleattachmentsKind = schema.GroupVersionKind{Group: "ram.alicloud.kubeform.com", Version: "v1alpha1", Kind: "RoleAttachment"}

// Get takes name of the roleAttachment, and returns the corresponding roleAttachment object, and an error if there is any.
func (c *FakeRoleAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RoleAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(roleattachmentsResource, c.ns, name), &v1alpha1.RoleAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAttachment), err
}

// List takes label and field selectors, and returns the list of RoleAttachments that match those selectors.
func (c *FakeRoleAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RoleAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(roleattachmentsResource, roleattachmentsKind, c.ns, opts), &v1alpha1.RoleAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoleAttachmentList{ListMeta: obj.(*v1alpha1.RoleAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoleAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleAttachments.
func (c *FakeRoleAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(roleattachmentsResource, c.ns, opts))

}

// Create takes the representation of a roleAttachment and creates it.  Returns the server's representation of the roleAttachment, and an error, if there is any.
func (c *FakeRoleAttachments) Create(ctx context.Context, roleAttachment *v1alpha1.RoleAttachment, opts v1.CreateOptions) (result *v1alpha1.RoleAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(roleattachmentsResource, c.ns, roleAttachment), &v1alpha1.RoleAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAttachment), err
}

// Update takes the representation of a roleAttachment and updates it. Returns the server's representation of the roleAttachment, and an error, if there is any.
func (c *FakeRoleAttachments) Update(ctx context.Context, roleAttachment *v1alpha1.RoleAttachment, opts v1.UpdateOptions) (result *v1alpha1.RoleAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(roleattachmentsResource, c.ns, roleAttachment), &v1alpha1.RoleAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoleAttachments) UpdateStatus(ctx context.Context, roleAttachment *v1alpha1.RoleAttachment, opts v1.UpdateOptions) (*v1alpha1.RoleAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(roleattachmentsResource, "status", c.ns, roleAttachment), &v1alpha1.RoleAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAttachment), err
}

// Delete takes name of the roleAttachment and deletes it. Returns an error if one occurs.
func (c *FakeRoleAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(roleattachmentsResource, c.ns, name), &v1alpha1.RoleAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(roleattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RoleAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched roleAttachment.
func (c *FakeRoleAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RoleAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(roleattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RoleAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAttachment), err
}
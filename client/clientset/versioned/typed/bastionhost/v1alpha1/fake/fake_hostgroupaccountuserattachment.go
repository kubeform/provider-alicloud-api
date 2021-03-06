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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/bastionhost/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostGroupAccountUserAttachments implements HostGroupAccountUserAttachmentInterface
type FakeHostGroupAccountUserAttachments struct {
	Fake *FakeBastionhostV1alpha1
	ns   string
}

var hostgroupaccountuserattachmentsResource = schema.GroupVersionResource{Group: "bastionhost.alicloud.kubeform.com", Version: "v1alpha1", Resource: "hostgroupaccountuserattachments"}

var hostgroupaccountuserattachmentsKind = schema.GroupVersionKind{Group: "bastionhost.alicloud.kubeform.com", Version: "v1alpha1", Kind: "HostGroupAccountUserAttachment"}

// Get takes name of the hostGroupAccountUserAttachment, and returns the corresponding hostGroupAccountUserAttachment object, and an error if there is any.
func (c *FakeHostGroupAccountUserAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HostGroupAccountUserAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostgroupaccountuserattachmentsResource, c.ns, name), &v1alpha1.HostGroupAccountUserAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostGroupAccountUserAttachment), err
}

// List takes label and field selectors, and returns the list of HostGroupAccountUserAttachments that match those selectors.
func (c *FakeHostGroupAccountUserAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HostGroupAccountUserAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostgroupaccountuserattachmentsResource, hostgroupaccountuserattachmentsKind, c.ns, opts), &v1alpha1.HostGroupAccountUserAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HostGroupAccountUserAttachmentList{ListMeta: obj.(*v1alpha1.HostGroupAccountUserAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.HostGroupAccountUserAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostGroupAccountUserAttachments.
func (c *FakeHostGroupAccountUserAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostgroupaccountuserattachmentsResource, c.ns, opts))

}

// Create takes the representation of a hostGroupAccountUserAttachment and creates it.  Returns the server's representation of the hostGroupAccountUserAttachment, and an error, if there is any.
func (c *FakeHostGroupAccountUserAttachments) Create(ctx context.Context, hostGroupAccountUserAttachment *v1alpha1.HostGroupAccountUserAttachment, opts v1.CreateOptions) (result *v1alpha1.HostGroupAccountUserAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostgroupaccountuserattachmentsResource, c.ns, hostGroupAccountUserAttachment), &v1alpha1.HostGroupAccountUserAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostGroupAccountUserAttachment), err
}

// Update takes the representation of a hostGroupAccountUserAttachment and updates it. Returns the server's representation of the hostGroupAccountUserAttachment, and an error, if there is any.
func (c *FakeHostGroupAccountUserAttachments) Update(ctx context.Context, hostGroupAccountUserAttachment *v1alpha1.HostGroupAccountUserAttachment, opts v1.UpdateOptions) (result *v1alpha1.HostGroupAccountUserAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostgroupaccountuserattachmentsResource, c.ns, hostGroupAccountUserAttachment), &v1alpha1.HostGroupAccountUserAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostGroupAccountUserAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHostGroupAccountUserAttachments) UpdateStatus(ctx context.Context, hostGroupAccountUserAttachment *v1alpha1.HostGroupAccountUserAttachment, opts v1.UpdateOptions) (*v1alpha1.HostGroupAccountUserAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hostgroupaccountuserattachmentsResource, "status", c.ns, hostGroupAccountUserAttachment), &v1alpha1.HostGroupAccountUserAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostGroupAccountUserAttachment), err
}

// Delete takes name of the hostGroupAccountUserAttachment and deletes it. Returns an error if one occurs.
func (c *FakeHostGroupAccountUserAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hostgroupaccountuserattachmentsResource, c.ns, name), &v1alpha1.HostGroupAccountUserAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostGroupAccountUserAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostgroupaccountuserattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HostGroupAccountUserAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched hostGroupAccountUserAttachment.
func (c *FakeHostGroupAccountUserAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HostGroupAccountUserAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostgroupaccountuserattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HostGroupAccountUserAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostGroupAccountUserAttachment), err
}

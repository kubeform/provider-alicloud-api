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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/edas/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstanceClusterAttachments implements InstanceClusterAttachmentInterface
type FakeInstanceClusterAttachments struct {
	Fake *FakeEdasV1alpha1
	ns   string
}

var instanceclusterattachmentsResource = schema.GroupVersionResource{Group: "edas.alicloud.kubeform.com", Version: "v1alpha1", Resource: "instanceclusterattachments"}

var instanceclusterattachmentsKind = schema.GroupVersionKind{Group: "edas.alicloud.kubeform.com", Version: "v1alpha1", Kind: "InstanceClusterAttachment"}

// Get takes name of the instanceClusterAttachment, and returns the corresponding instanceClusterAttachment object, and an error if there is any.
func (c *FakeInstanceClusterAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstanceClusterAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(instanceclusterattachmentsResource, c.ns, name), &v1alpha1.InstanceClusterAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceClusterAttachment), err
}

// List takes label and field selectors, and returns the list of InstanceClusterAttachments that match those selectors.
func (c *FakeInstanceClusterAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstanceClusterAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(instanceclusterattachmentsResource, instanceclusterattachmentsKind, c.ns, opts), &v1alpha1.InstanceClusterAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstanceClusterAttachmentList{ListMeta: obj.(*v1alpha1.InstanceClusterAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstanceClusterAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instanceClusterAttachments.
func (c *FakeInstanceClusterAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(instanceclusterattachmentsResource, c.ns, opts))

}

// Create takes the representation of a instanceClusterAttachment and creates it.  Returns the server's representation of the instanceClusterAttachment, and an error, if there is any.
func (c *FakeInstanceClusterAttachments) Create(ctx context.Context, instanceClusterAttachment *v1alpha1.InstanceClusterAttachment, opts v1.CreateOptions) (result *v1alpha1.InstanceClusterAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(instanceclusterattachmentsResource, c.ns, instanceClusterAttachment), &v1alpha1.InstanceClusterAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceClusterAttachment), err
}

// Update takes the representation of a instanceClusterAttachment and updates it. Returns the server's representation of the instanceClusterAttachment, and an error, if there is any.
func (c *FakeInstanceClusterAttachments) Update(ctx context.Context, instanceClusterAttachment *v1alpha1.InstanceClusterAttachment, opts v1.UpdateOptions) (result *v1alpha1.InstanceClusterAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(instanceclusterattachmentsResource, c.ns, instanceClusterAttachment), &v1alpha1.InstanceClusterAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceClusterAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstanceClusterAttachments) UpdateStatus(ctx context.Context, instanceClusterAttachment *v1alpha1.InstanceClusterAttachment, opts v1.UpdateOptions) (*v1alpha1.InstanceClusterAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(instanceclusterattachmentsResource, "status", c.ns, instanceClusterAttachment), &v1alpha1.InstanceClusterAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceClusterAttachment), err
}

// Delete takes name of the instanceClusterAttachment and deletes it. Returns an error if one occurs.
func (c *FakeInstanceClusterAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(instanceclusterattachmentsResource, c.ns, name), &v1alpha1.InstanceClusterAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstanceClusterAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(instanceclusterattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstanceClusterAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched instanceClusterAttachment.
func (c *FakeInstanceClusterAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstanceClusterAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(instanceclusterattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstanceClusterAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceClusterAttachment), err
}

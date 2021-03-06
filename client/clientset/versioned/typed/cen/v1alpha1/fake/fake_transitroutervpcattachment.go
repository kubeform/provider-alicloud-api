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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cen/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTransitRouterVpcAttachments implements TransitRouterVpcAttachmentInterface
type FakeTransitRouterVpcAttachments struct {
	Fake *FakeCenV1alpha1
	ns   string
}

var transitroutervpcattachmentsResource = schema.GroupVersionResource{Group: "cen.alicloud.kubeform.com", Version: "v1alpha1", Resource: "transitroutervpcattachments"}

var transitroutervpcattachmentsKind = schema.GroupVersionKind{Group: "cen.alicloud.kubeform.com", Version: "v1alpha1", Kind: "TransitRouterVpcAttachment"}

// Get takes name of the transitRouterVpcAttachment, and returns the corresponding transitRouterVpcAttachment object, and an error if there is any.
func (c *FakeTransitRouterVpcAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitRouterVpcAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitroutervpcattachmentsResource, c.ns, name), &v1alpha1.TransitRouterVpcAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterVpcAttachment), err
}

// List takes label and field selectors, and returns the list of TransitRouterVpcAttachments that match those selectors.
func (c *FakeTransitRouterVpcAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitRouterVpcAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitroutervpcattachmentsResource, transitroutervpcattachmentsKind, c.ns, opts), &v1alpha1.TransitRouterVpcAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitRouterVpcAttachmentList{ListMeta: obj.(*v1alpha1.TransitRouterVpcAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitRouterVpcAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitRouterVpcAttachments.
func (c *FakeTransitRouterVpcAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitroutervpcattachmentsResource, c.ns, opts))

}

// Create takes the representation of a transitRouterVpcAttachment and creates it.  Returns the server's representation of the transitRouterVpcAttachment, and an error, if there is any.
func (c *FakeTransitRouterVpcAttachments) Create(ctx context.Context, transitRouterVpcAttachment *v1alpha1.TransitRouterVpcAttachment, opts v1.CreateOptions) (result *v1alpha1.TransitRouterVpcAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitroutervpcattachmentsResource, c.ns, transitRouterVpcAttachment), &v1alpha1.TransitRouterVpcAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterVpcAttachment), err
}

// Update takes the representation of a transitRouterVpcAttachment and updates it. Returns the server's representation of the transitRouterVpcAttachment, and an error, if there is any.
func (c *FakeTransitRouterVpcAttachments) Update(ctx context.Context, transitRouterVpcAttachment *v1alpha1.TransitRouterVpcAttachment, opts v1.UpdateOptions) (result *v1alpha1.TransitRouterVpcAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitroutervpcattachmentsResource, c.ns, transitRouterVpcAttachment), &v1alpha1.TransitRouterVpcAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterVpcAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitRouterVpcAttachments) UpdateStatus(ctx context.Context, transitRouterVpcAttachment *v1alpha1.TransitRouterVpcAttachment, opts v1.UpdateOptions) (*v1alpha1.TransitRouterVpcAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitroutervpcattachmentsResource, "status", c.ns, transitRouterVpcAttachment), &v1alpha1.TransitRouterVpcAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterVpcAttachment), err
}

// Delete takes name of the transitRouterVpcAttachment and deletes it. Returns an error if one occurs.
func (c *FakeTransitRouterVpcAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitroutervpcattachmentsResource, c.ns, name), &v1alpha1.TransitRouterVpcAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitRouterVpcAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitroutervpcattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitRouterVpcAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched transitRouterVpcAttachment.
func (c *FakeTransitRouterVpcAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitRouterVpcAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitroutervpcattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitRouterVpcAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterVpcAttachment), err
}

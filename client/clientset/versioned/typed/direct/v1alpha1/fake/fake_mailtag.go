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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/direct/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMailTags implements MailTagInterface
type FakeMailTags struct {
	Fake *FakeDirectV1alpha1
	ns   string
}

var mailtagsResource = schema.GroupVersionResource{Group: "direct.alicloud.kubeform.com", Version: "v1alpha1", Resource: "mailtags"}

var mailtagsKind = schema.GroupVersionKind{Group: "direct.alicloud.kubeform.com", Version: "v1alpha1", Kind: "MailTag"}

// Get takes name of the mailTag, and returns the corresponding mailTag object, and an error if there is any.
func (c *FakeMailTags) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MailTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mailtagsResource, c.ns, name), &v1alpha1.MailTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MailTag), err
}

// List takes label and field selectors, and returns the list of MailTags that match those selectors.
func (c *FakeMailTags) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MailTagList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mailtagsResource, mailtagsKind, c.ns, opts), &v1alpha1.MailTagList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MailTagList{ListMeta: obj.(*v1alpha1.MailTagList).ListMeta}
	for _, item := range obj.(*v1alpha1.MailTagList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mailTags.
func (c *FakeMailTags) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mailtagsResource, c.ns, opts))

}

// Create takes the representation of a mailTag and creates it.  Returns the server's representation of the mailTag, and an error, if there is any.
func (c *FakeMailTags) Create(ctx context.Context, mailTag *v1alpha1.MailTag, opts v1.CreateOptions) (result *v1alpha1.MailTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mailtagsResource, c.ns, mailTag), &v1alpha1.MailTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MailTag), err
}

// Update takes the representation of a mailTag and updates it. Returns the server's representation of the mailTag, and an error, if there is any.
func (c *FakeMailTags) Update(ctx context.Context, mailTag *v1alpha1.MailTag, opts v1.UpdateOptions) (result *v1alpha1.MailTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mailtagsResource, c.ns, mailTag), &v1alpha1.MailTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MailTag), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMailTags) UpdateStatus(ctx context.Context, mailTag *v1alpha1.MailTag, opts v1.UpdateOptions) (*v1alpha1.MailTag, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mailtagsResource, "status", c.ns, mailTag), &v1alpha1.MailTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MailTag), err
}

// Delete takes name of the mailTag and deletes it. Returns an error if one occurs.
func (c *FakeMailTags) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mailtagsResource, c.ns, name), &v1alpha1.MailTag{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMailTags) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mailtagsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MailTagList{})
	return err
}

// Patch applies the patch and returns the patched mailTag.
func (c *FakeMailTags) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MailTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mailtagsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MailTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MailTag), err
}

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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/image/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCopies implements CopyInterface
type FakeCopies struct {
	Fake *FakeImageV1alpha1
	ns   string
}

var copiesResource = schema.GroupVersionResource{Group: "image.alicloud.kubeform.com", Version: "v1alpha1", Resource: "copies"}

var copiesKind = schema.GroupVersionKind{Group: "image.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Copy"}

// Get takes name of the copy, and returns the corresponding copy object, and an error if there is any.
func (c *FakeCopies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Copy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(copiesResource, c.ns, name), &v1alpha1.Copy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Copy), err
}

// List takes label and field selectors, and returns the list of Copies that match those selectors.
func (c *FakeCopies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CopyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(copiesResource, copiesKind, c.ns, opts), &v1alpha1.CopyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CopyList{ListMeta: obj.(*v1alpha1.CopyList).ListMeta}
	for _, item := range obj.(*v1alpha1.CopyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested copies.
func (c *FakeCopies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(copiesResource, c.ns, opts))

}

// Create takes the representation of a copy and creates it.  Returns the server's representation of the copy, and an error, if there is any.
func (c *FakeCopies) Create(ctx context.Context, copy *v1alpha1.Copy, opts v1.CreateOptions) (result *v1alpha1.Copy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(copiesResource, c.ns, copy), &v1alpha1.Copy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Copy), err
}

// Update takes the representation of a copy and updates it. Returns the server's representation of the copy, and an error, if there is any.
func (c *FakeCopies) Update(ctx context.Context, copy *v1alpha1.Copy, opts v1.UpdateOptions) (result *v1alpha1.Copy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(copiesResource, c.ns, copy), &v1alpha1.Copy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Copy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCopies) UpdateStatus(ctx context.Context, copy *v1alpha1.Copy, opts v1.UpdateOptions) (*v1alpha1.Copy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(copiesResource, "status", c.ns, copy), &v1alpha1.Copy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Copy), err
}

// Delete takes name of the copy and deletes it. Returns an error if one occurs.
func (c *FakeCopies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(copiesResource, c.ns, name), &v1alpha1.Copy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCopies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(copiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CopyList{})
	return err
}

// Patch applies the patch and returns the patched copy.
func (c *FakeCopies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Copy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(copiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Copy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Copy), err
}

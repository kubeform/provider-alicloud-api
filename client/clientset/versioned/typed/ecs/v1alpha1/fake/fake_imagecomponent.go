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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageComponents implements ImageComponentInterface
type FakeImageComponents struct {
	Fake *FakeEcsV1alpha1
	ns   string
}

var imagecomponentsResource = schema.GroupVersionResource{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "imagecomponents"}

var imagecomponentsKind = schema.GroupVersionKind{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ImageComponent"}

// Get takes name of the imageComponent, and returns the corresponding imageComponent object, and an error if there is any.
func (c *FakeImageComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagecomponentsResource, c.ns, name), &v1alpha1.ImageComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageComponent), err
}

// List takes label and field selectors, and returns the list of ImageComponents that match those selectors.
func (c *FakeImageComponents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageComponentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagecomponentsResource, imagecomponentsKind, c.ns, opts), &v1alpha1.ImageComponentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageComponentList{ListMeta: obj.(*v1alpha1.ImageComponentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageComponentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageComponents.
func (c *FakeImageComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imagecomponentsResource, c.ns, opts))

}

// Create takes the representation of a imageComponent and creates it.  Returns the server's representation of the imageComponent, and an error, if there is any.
func (c *FakeImageComponents) Create(ctx context.Context, imageComponent *v1alpha1.ImageComponent, opts v1.CreateOptions) (result *v1alpha1.ImageComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagecomponentsResource, c.ns, imageComponent), &v1alpha1.ImageComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageComponent), err
}

// Update takes the representation of a imageComponent and updates it. Returns the server's representation of the imageComponent, and an error, if there is any.
func (c *FakeImageComponents) Update(ctx context.Context, imageComponent *v1alpha1.ImageComponent, opts v1.UpdateOptions) (result *v1alpha1.ImageComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagecomponentsResource, c.ns, imageComponent), &v1alpha1.ImageComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageComponent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageComponents) UpdateStatus(ctx context.Context, imageComponent *v1alpha1.ImageComponent, opts v1.UpdateOptions) (*v1alpha1.ImageComponent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imagecomponentsResource, "status", c.ns, imageComponent), &v1alpha1.ImageComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageComponent), err
}

// Delete takes name of the imageComponent and deletes it. Returns an error if one occurs.
func (c *FakeImageComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imagecomponentsResource, c.ns, name), &v1alpha1.ImageComponent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imagecomponentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageComponentList{})
	return err
}

// Patch applies the patch and returns the patched imageComponent.
func (c *FakeImageComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagecomponentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImageComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageComponent), err
}

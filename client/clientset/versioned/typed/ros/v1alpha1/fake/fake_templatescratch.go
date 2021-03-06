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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ros/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplateScratches implements TemplateScratchInterface
type FakeTemplateScratches struct {
	Fake *FakeRosV1alpha1
	ns   string
}

var templatescratchesResource = schema.GroupVersionResource{Group: "ros.alicloud.kubeform.com", Version: "v1alpha1", Resource: "templatescratches"}

var templatescratchesKind = schema.GroupVersionKind{Group: "ros.alicloud.kubeform.com", Version: "v1alpha1", Kind: "TemplateScratch"}

// Get takes name of the templateScratch, and returns the corresponding templateScratch object, and an error if there is any.
func (c *FakeTemplateScratches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TemplateScratch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(templatescratchesResource, c.ns, name), &v1alpha1.TemplateScratch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateScratch), err
}

// List takes label and field selectors, and returns the list of TemplateScratches that match those selectors.
func (c *FakeTemplateScratches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TemplateScratchList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(templatescratchesResource, templatescratchesKind, c.ns, opts), &v1alpha1.TemplateScratchList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TemplateScratchList{ListMeta: obj.(*v1alpha1.TemplateScratchList).ListMeta}
	for _, item := range obj.(*v1alpha1.TemplateScratchList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templateScratches.
func (c *FakeTemplateScratches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(templatescratchesResource, c.ns, opts))

}

// Create takes the representation of a templateScratch and creates it.  Returns the server's representation of the templateScratch, and an error, if there is any.
func (c *FakeTemplateScratches) Create(ctx context.Context, templateScratch *v1alpha1.TemplateScratch, opts v1.CreateOptions) (result *v1alpha1.TemplateScratch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(templatescratchesResource, c.ns, templateScratch), &v1alpha1.TemplateScratch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateScratch), err
}

// Update takes the representation of a templateScratch and updates it. Returns the server's representation of the templateScratch, and an error, if there is any.
func (c *FakeTemplateScratches) Update(ctx context.Context, templateScratch *v1alpha1.TemplateScratch, opts v1.UpdateOptions) (result *v1alpha1.TemplateScratch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(templatescratchesResource, c.ns, templateScratch), &v1alpha1.TemplateScratch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateScratch), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTemplateScratches) UpdateStatus(ctx context.Context, templateScratch *v1alpha1.TemplateScratch, opts v1.UpdateOptions) (*v1alpha1.TemplateScratch, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(templatescratchesResource, "status", c.ns, templateScratch), &v1alpha1.TemplateScratch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateScratch), err
}

// Delete takes name of the templateScratch and deletes it. Returns an error if one occurs.
func (c *FakeTemplateScratches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(templatescratchesResource, c.ns, name), &v1alpha1.TemplateScratch{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplateScratches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(templatescratchesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TemplateScratchList{})
	return err
}

// Patch applies the patch and returns the patched templateScratch.
func (c *FakeTemplateScratches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TemplateScratch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(templatescratchesResource, c.ns, name, pt, data, subresources...), &v1alpha1.TemplateScratch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateScratch), err
}

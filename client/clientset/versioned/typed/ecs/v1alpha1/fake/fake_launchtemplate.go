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

// FakeLaunchTemplates implements LaunchTemplateInterface
type FakeLaunchTemplates struct {
	Fake *FakeEcsV1alpha1
	ns   string
}

var launchtemplatesResource = schema.GroupVersionResource{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "launchtemplates"}

var launchtemplatesKind = schema.GroupVersionKind{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "LaunchTemplate"}

// Get takes name of the launchTemplate, and returns the corresponding launchTemplate object, and an error if there is any.
func (c *FakeLaunchTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LaunchTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(launchtemplatesResource, c.ns, name), &v1alpha1.LaunchTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchTemplate), err
}

// List takes label and field selectors, and returns the list of LaunchTemplates that match those selectors.
func (c *FakeLaunchTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LaunchTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(launchtemplatesResource, launchtemplatesKind, c.ns, opts), &v1alpha1.LaunchTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LaunchTemplateList{ListMeta: obj.(*v1alpha1.LaunchTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.LaunchTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested launchTemplates.
func (c *FakeLaunchTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(launchtemplatesResource, c.ns, opts))

}

// Create takes the representation of a launchTemplate and creates it.  Returns the server's representation of the launchTemplate, and an error, if there is any.
func (c *FakeLaunchTemplates) Create(ctx context.Context, launchTemplate *v1alpha1.LaunchTemplate, opts v1.CreateOptions) (result *v1alpha1.LaunchTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(launchtemplatesResource, c.ns, launchTemplate), &v1alpha1.LaunchTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchTemplate), err
}

// Update takes the representation of a launchTemplate and updates it. Returns the server's representation of the launchTemplate, and an error, if there is any.
func (c *FakeLaunchTemplates) Update(ctx context.Context, launchTemplate *v1alpha1.LaunchTemplate, opts v1.UpdateOptions) (result *v1alpha1.LaunchTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(launchtemplatesResource, c.ns, launchTemplate), &v1alpha1.LaunchTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLaunchTemplates) UpdateStatus(ctx context.Context, launchTemplate *v1alpha1.LaunchTemplate, opts v1.UpdateOptions) (*v1alpha1.LaunchTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(launchtemplatesResource, "status", c.ns, launchTemplate), &v1alpha1.LaunchTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchTemplate), err
}

// Delete takes name of the launchTemplate and deletes it. Returns an error if one occurs.
func (c *FakeLaunchTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(launchtemplatesResource, c.ns, name), &v1alpha1.LaunchTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLaunchTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(launchtemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LaunchTemplateList{})
	return err
}

// Patch applies the patch and returns the patched launchTemplate.
func (c *FakeLaunchTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LaunchTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(launchtemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LaunchTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchTemplate), err
}
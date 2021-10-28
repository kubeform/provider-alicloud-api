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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/alb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHealthCheckTemplates implements HealthCheckTemplateInterface
type FakeHealthCheckTemplates struct {
	Fake *FakeAlbV1alpha1
	ns   string
}

var healthchecktemplatesResource = schema.GroupVersionResource{Group: "alb.alicloud.kubeform.com", Version: "v1alpha1", Resource: "healthchecktemplates"}

var healthchecktemplatesKind = schema.GroupVersionKind{Group: "alb.alicloud.kubeform.com", Version: "v1alpha1", Kind: "HealthCheckTemplate"}

// Get takes name of the healthCheckTemplate, and returns the corresponding healthCheckTemplate object, and an error if there is any.
func (c *FakeHealthCheckTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HealthCheckTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(healthchecktemplatesResource, c.ns, name), &v1alpha1.HealthCheckTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthCheckTemplate), err
}

// List takes label and field selectors, and returns the list of HealthCheckTemplates that match those selectors.
func (c *FakeHealthCheckTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HealthCheckTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(healthchecktemplatesResource, healthchecktemplatesKind, c.ns, opts), &v1alpha1.HealthCheckTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HealthCheckTemplateList{ListMeta: obj.(*v1alpha1.HealthCheckTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.HealthCheckTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested healthCheckTemplates.
func (c *FakeHealthCheckTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(healthchecktemplatesResource, c.ns, opts))

}

// Create takes the representation of a healthCheckTemplate and creates it.  Returns the server's representation of the healthCheckTemplate, and an error, if there is any.
func (c *FakeHealthCheckTemplates) Create(ctx context.Context, healthCheckTemplate *v1alpha1.HealthCheckTemplate, opts v1.CreateOptions) (result *v1alpha1.HealthCheckTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(healthchecktemplatesResource, c.ns, healthCheckTemplate), &v1alpha1.HealthCheckTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthCheckTemplate), err
}

// Update takes the representation of a healthCheckTemplate and updates it. Returns the server's representation of the healthCheckTemplate, and an error, if there is any.
func (c *FakeHealthCheckTemplates) Update(ctx context.Context, healthCheckTemplate *v1alpha1.HealthCheckTemplate, opts v1.UpdateOptions) (result *v1alpha1.HealthCheckTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(healthchecktemplatesResource, c.ns, healthCheckTemplate), &v1alpha1.HealthCheckTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthCheckTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHealthCheckTemplates) UpdateStatus(ctx context.Context, healthCheckTemplate *v1alpha1.HealthCheckTemplate, opts v1.UpdateOptions) (*v1alpha1.HealthCheckTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(healthchecktemplatesResource, "status", c.ns, healthCheckTemplate), &v1alpha1.HealthCheckTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthCheckTemplate), err
}

// Delete takes name of the healthCheckTemplate and deletes it. Returns an error if one occurs.
func (c *FakeHealthCheckTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(healthchecktemplatesResource, c.ns, name), &v1alpha1.HealthCheckTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHealthCheckTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(healthchecktemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HealthCheckTemplateList{})
	return err
}

// Patch applies the patch and returns the patched healthCheckTemplate.
func (c *FakeHealthCheckTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HealthCheckTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(healthchecktemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HealthCheckTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthCheckTemplate), err
}

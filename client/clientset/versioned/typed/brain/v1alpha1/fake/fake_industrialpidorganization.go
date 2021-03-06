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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/brain/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIndustrialPidOrganizations implements IndustrialPidOrganizationInterface
type FakeIndustrialPidOrganizations struct {
	Fake *FakeBrainV1alpha1
	ns   string
}

var industrialpidorganizationsResource = schema.GroupVersionResource{Group: "brain.alicloud.kubeform.com", Version: "v1alpha1", Resource: "industrialpidorganizations"}

var industrialpidorganizationsKind = schema.GroupVersionKind{Group: "brain.alicloud.kubeform.com", Version: "v1alpha1", Kind: "IndustrialPidOrganization"}

// Get takes name of the industrialPidOrganization, and returns the corresponding industrialPidOrganization object, and an error if there is any.
func (c *FakeIndustrialPidOrganizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IndustrialPidOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(industrialpidorganizationsResource, c.ns, name), &v1alpha1.IndustrialPidOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IndustrialPidOrganization), err
}

// List takes label and field selectors, and returns the list of IndustrialPidOrganizations that match those selectors.
func (c *FakeIndustrialPidOrganizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IndustrialPidOrganizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(industrialpidorganizationsResource, industrialpidorganizationsKind, c.ns, opts), &v1alpha1.IndustrialPidOrganizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IndustrialPidOrganizationList{ListMeta: obj.(*v1alpha1.IndustrialPidOrganizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.IndustrialPidOrganizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested industrialPidOrganizations.
func (c *FakeIndustrialPidOrganizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(industrialpidorganizationsResource, c.ns, opts))

}

// Create takes the representation of a industrialPidOrganization and creates it.  Returns the server's representation of the industrialPidOrganization, and an error, if there is any.
func (c *FakeIndustrialPidOrganizations) Create(ctx context.Context, industrialPidOrganization *v1alpha1.IndustrialPidOrganization, opts v1.CreateOptions) (result *v1alpha1.IndustrialPidOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(industrialpidorganizationsResource, c.ns, industrialPidOrganization), &v1alpha1.IndustrialPidOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IndustrialPidOrganization), err
}

// Update takes the representation of a industrialPidOrganization and updates it. Returns the server's representation of the industrialPidOrganization, and an error, if there is any.
func (c *FakeIndustrialPidOrganizations) Update(ctx context.Context, industrialPidOrganization *v1alpha1.IndustrialPidOrganization, opts v1.UpdateOptions) (result *v1alpha1.IndustrialPidOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(industrialpidorganizationsResource, c.ns, industrialPidOrganization), &v1alpha1.IndustrialPidOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IndustrialPidOrganization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIndustrialPidOrganizations) UpdateStatus(ctx context.Context, industrialPidOrganization *v1alpha1.IndustrialPidOrganization, opts v1.UpdateOptions) (*v1alpha1.IndustrialPidOrganization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(industrialpidorganizationsResource, "status", c.ns, industrialPidOrganization), &v1alpha1.IndustrialPidOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IndustrialPidOrganization), err
}

// Delete takes name of the industrialPidOrganization and deletes it. Returns an error if one occurs.
func (c *FakeIndustrialPidOrganizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(industrialpidorganizationsResource, c.ns, name), &v1alpha1.IndustrialPidOrganization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIndustrialPidOrganizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(industrialpidorganizationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IndustrialPidOrganizationList{})
	return err
}

// Patch applies the patch and returns the patched industrialPidOrganization.
func (c *FakeIndustrialPidOrganizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IndustrialPidOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(industrialpidorganizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IndustrialPidOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IndustrialPidOrganization), err
}

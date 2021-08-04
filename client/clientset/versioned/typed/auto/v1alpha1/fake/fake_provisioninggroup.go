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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/auto/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProvisioningGroups implements ProvisioningGroupInterface
type FakeProvisioningGroups struct {
	Fake *FakeAutoV1alpha1
	ns   string
}

var provisioninggroupsResource = schema.GroupVersionResource{Group: "auto.alicloud.kubeform.com", Version: "v1alpha1", Resource: "provisioninggroups"}

var provisioninggroupsKind = schema.GroupVersionKind{Group: "auto.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ProvisioningGroup"}

// Get takes name of the provisioningGroup, and returns the corresponding provisioningGroup object, and an error if there is any.
func (c *FakeProvisioningGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProvisioningGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(provisioninggroupsResource, c.ns, name), &v1alpha1.ProvisioningGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisioningGroup), err
}

// List takes label and field selectors, and returns the list of ProvisioningGroups that match those selectors.
func (c *FakeProvisioningGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProvisioningGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(provisioninggroupsResource, provisioninggroupsKind, c.ns, opts), &v1alpha1.ProvisioningGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProvisioningGroupList{ListMeta: obj.(*v1alpha1.ProvisioningGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProvisioningGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested provisioningGroups.
func (c *FakeProvisioningGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(provisioninggroupsResource, c.ns, opts))

}

// Create takes the representation of a provisioningGroup and creates it.  Returns the server's representation of the provisioningGroup, and an error, if there is any.
func (c *FakeProvisioningGroups) Create(ctx context.Context, provisioningGroup *v1alpha1.ProvisioningGroup, opts v1.CreateOptions) (result *v1alpha1.ProvisioningGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(provisioninggroupsResource, c.ns, provisioningGroup), &v1alpha1.ProvisioningGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisioningGroup), err
}

// Update takes the representation of a provisioningGroup and updates it. Returns the server's representation of the provisioningGroup, and an error, if there is any.
func (c *FakeProvisioningGroups) Update(ctx context.Context, provisioningGroup *v1alpha1.ProvisioningGroup, opts v1.UpdateOptions) (result *v1alpha1.ProvisioningGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(provisioninggroupsResource, c.ns, provisioningGroup), &v1alpha1.ProvisioningGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisioningGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProvisioningGroups) UpdateStatus(ctx context.Context, provisioningGroup *v1alpha1.ProvisioningGroup, opts v1.UpdateOptions) (*v1alpha1.ProvisioningGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(provisioninggroupsResource, "status", c.ns, provisioningGroup), &v1alpha1.ProvisioningGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisioningGroup), err
}

// Delete takes name of the provisioningGroup and deletes it. Returns an error if one occurs.
func (c *FakeProvisioningGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(provisioninggroupsResource, c.ns, name), &v1alpha1.ProvisioningGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProvisioningGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(provisioninggroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProvisioningGroupList{})
	return err
}

// Patch applies the patch and returns the patched provisioningGroup.
func (c *FakeProvisioningGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProvisioningGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(provisioninggroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProvisioningGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisioningGroup), err
}
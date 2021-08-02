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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/nas/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessGroups implements AccessGroupInterface
type FakeAccessGroups struct {
	Fake *FakeNasV1alpha1
	ns   string
}

var accessgroupsResource = schema.GroupVersionResource{Group: "nas.alicloud.kubeform.com", Version: "v1alpha1", Resource: "accessgroups"}

var accessgroupsKind = schema.GroupVersionKind{Group: "nas.alicloud.kubeform.com", Version: "v1alpha1", Kind: "AccessGroup"}

// Get takes name of the accessGroup, and returns the corresponding accessGroup object, and an error if there is any.
func (c *FakeAccessGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accessgroupsResource, c.ns, name), &v1alpha1.AccessGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessGroup), err
}

// List takes label and field selectors, and returns the list of AccessGroups that match those selectors.
func (c *FakeAccessGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accessgroupsResource, accessgroupsKind, c.ns, opts), &v1alpha1.AccessGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessGroupList{ListMeta: obj.(*v1alpha1.AccessGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessGroups.
func (c *FakeAccessGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accessgroupsResource, c.ns, opts))

}

// Create takes the representation of a accessGroup and creates it.  Returns the server's representation of the accessGroup, and an error, if there is any.
func (c *FakeAccessGroups) Create(ctx context.Context, accessGroup *v1alpha1.AccessGroup, opts v1.CreateOptions) (result *v1alpha1.AccessGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accessgroupsResource, c.ns, accessGroup), &v1alpha1.AccessGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessGroup), err
}

// Update takes the representation of a accessGroup and updates it. Returns the server's representation of the accessGroup, and an error, if there is any.
func (c *FakeAccessGroups) Update(ctx context.Context, accessGroup *v1alpha1.AccessGroup, opts v1.UpdateOptions) (result *v1alpha1.AccessGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accessgroupsResource, c.ns, accessGroup), &v1alpha1.AccessGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessGroups) UpdateStatus(ctx context.Context, accessGroup *v1alpha1.AccessGroup, opts v1.UpdateOptions) (*v1alpha1.AccessGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accessgroupsResource, "status", c.ns, accessGroup), &v1alpha1.AccessGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessGroup), err
}

// Delete takes name of the accessGroup and deletes it. Returns an error if one occurs.
func (c *FakeAccessGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accessgroupsResource, c.ns, name), &v1alpha1.AccessGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accessgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessGroupList{})
	return err
}

// Patch applies the patch and returns the patched accessGroup.
func (c *FakeAccessGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accessgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessGroup), err
}

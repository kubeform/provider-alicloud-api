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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagerResourceGroups implements ManagerResourceGroupInterface
type FakeManagerResourceGroups struct {
	Fake *FakeResourceV1alpha1
	ns   string
}

var managerresourcegroupsResource = schema.GroupVersionResource{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Resource: "managerresourcegroups"}

var managerresourcegroupsKind = schema.GroupVersionKind{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ManagerResourceGroup"}

// Get takes name of the managerResourceGroup, and returns the corresponding managerResourceGroup object, and an error if there is any.
func (c *FakeManagerResourceGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managerresourcegroupsResource, c.ns, name), &v1alpha1.ManagerResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerResourceGroup), err
}

// List takes label and field selectors, and returns the list of ManagerResourceGroups that match those selectors.
func (c *FakeManagerResourceGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerResourceGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managerresourcegroupsResource, managerresourcegroupsKind, c.ns, opts), &v1alpha1.ManagerResourceGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerResourceGroupList{ListMeta: obj.(*v1alpha1.ManagerResourceGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerResourceGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerResourceGroups.
func (c *FakeManagerResourceGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managerresourcegroupsResource, c.ns, opts))

}

// Create takes the representation of a managerResourceGroup and creates it.  Returns the server's representation of the managerResourceGroup, and an error, if there is any.
func (c *FakeManagerResourceGroups) Create(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.CreateOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managerresourcegroupsResource, c.ns, managerResourceGroup), &v1alpha1.ManagerResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerResourceGroup), err
}

// Update takes the representation of a managerResourceGroup and updates it. Returns the server's representation of the managerResourceGroup, and an error, if there is any.
func (c *FakeManagerResourceGroups) Update(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.UpdateOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managerresourcegroupsResource, c.ns, managerResourceGroup), &v1alpha1.ManagerResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerResourceGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerResourceGroups) UpdateStatus(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.UpdateOptions) (*v1alpha1.ManagerResourceGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managerresourcegroupsResource, "status", c.ns, managerResourceGroup), &v1alpha1.ManagerResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerResourceGroup), err
}

// Delete takes name of the managerResourceGroup and deletes it. Returns an error if one occurs.
func (c *FakeManagerResourceGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managerresourcegroupsResource, c.ns, name), &v1alpha1.ManagerResourceGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerResourceGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managerresourcegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerResourceGroupList{})
	return err
}

// Patch applies the patch and returns the patched managerResourceGroup.
func (c *FakeManagerResourceGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managerresourcegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerResourceGroup), err
}

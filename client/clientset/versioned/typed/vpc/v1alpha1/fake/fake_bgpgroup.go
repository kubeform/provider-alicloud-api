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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vpc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBgpGroups implements BgpGroupInterface
type FakeBgpGroups struct {
	Fake *FakeVpcV1alpha1
	ns   string
}

var bgpgroupsResource = schema.GroupVersionResource{Group: "vpc.alicloud.kubeform.com", Version: "v1alpha1", Resource: "bgpgroups"}

var bgpgroupsKind = schema.GroupVersionKind{Group: "vpc.alicloud.kubeform.com", Version: "v1alpha1", Kind: "BgpGroup"}

// Get takes name of the bgpGroup, and returns the corresponding bgpGroup object, and an error if there is any.
func (c *FakeBgpGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BgpGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bgpgroupsResource, c.ns, name), &v1alpha1.BgpGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BgpGroup), err
}

// List takes label and field selectors, and returns the list of BgpGroups that match those selectors.
func (c *FakeBgpGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BgpGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bgpgroupsResource, bgpgroupsKind, c.ns, opts), &v1alpha1.BgpGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BgpGroupList{ListMeta: obj.(*v1alpha1.BgpGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.BgpGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bgpGroups.
func (c *FakeBgpGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bgpgroupsResource, c.ns, opts))

}

// Create takes the representation of a bgpGroup and creates it.  Returns the server's representation of the bgpGroup, and an error, if there is any.
func (c *FakeBgpGroups) Create(ctx context.Context, bgpGroup *v1alpha1.BgpGroup, opts v1.CreateOptions) (result *v1alpha1.BgpGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bgpgroupsResource, c.ns, bgpGroup), &v1alpha1.BgpGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BgpGroup), err
}

// Update takes the representation of a bgpGroup and updates it. Returns the server's representation of the bgpGroup, and an error, if there is any.
func (c *FakeBgpGroups) Update(ctx context.Context, bgpGroup *v1alpha1.BgpGroup, opts v1.UpdateOptions) (result *v1alpha1.BgpGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bgpgroupsResource, c.ns, bgpGroup), &v1alpha1.BgpGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BgpGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBgpGroups) UpdateStatus(ctx context.Context, bgpGroup *v1alpha1.BgpGroup, opts v1.UpdateOptions) (*v1alpha1.BgpGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bgpgroupsResource, "status", c.ns, bgpGroup), &v1alpha1.BgpGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BgpGroup), err
}

// Delete takes name of the bgpGroup and deletes it. Returns an error if one occurs.
func (c *FakeBgpGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bgpgroupsResource, c.ns, name), &v1alpha1.BgpGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBgpGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bgpgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BgpGroupList{})
	return err
}

// Patch applies the patch and returns the patched bgpGroup.
func (c *FakeBgpGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BgpGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bgpgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BgpGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BgpGroup), err
}

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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecd/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyGroups implements PolicyGroupInterface
type FakePolicyGroups struct {
	Fake *FakeEcdV1alpha1
	ns   string
}

var policygroupsResource = schema.GroupVersionResource{Group: "ecd.alicloud.kubeform.com", Version: "v1alpha1", Resource: "policygroups"}

var policygroupsKind = schema.GroupVersionKind{Group: "ecd.alicloud.kubeform.com", Version: "v1alpha1", Kind: "PolicyGroup"}

// Get takes name of the policyGroup, and returns the corresponding policyGroup object, and an error if there is any.
func (c *FakePolicyGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policygroupsResource, c.ns, name), &v1alpha1.PolicyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyGroup), err
}

// List takes label and field selectors, and returns the list of PolicyGroups that match those selectors.
func (c *FakePolicyGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policygroupsResource, policygroupsKind, c.ns, opts), &v1alpha1.PolicyGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyGroupList{ListMeta: obj.(*v1alpha1.PolicyGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyGroups.
func (c *FakePolicyGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policygroupsResource, c.ns, opts))

}

// Create takes the representation of a policyGroup and creates it.  Returns the server's representation of the policyGroup, and an error, if there is any.
func (c *FakePolicyGroups) Create(ctx context.Context, policyGroup *v1alpha1.PolicyGroup, opts v1.CreateOptions) (result *v1alpha1.PolicyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policygroupsResource, c.ns, policyGroup), &v1alpha1.PolicyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyGroup), err
}

// Update takes the representation of a policyGroup and updates it. Returns the server's representation of the policyGroup, and an error, if there is any.
func (c *FakePolicyGroups) Update(ctx context.Context, policyGroup *v1alpha1.PolicyGroup, opts v1.UpdateOptions) (result *v1alpha1.PolicyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policygroupsResource, c.ns, policyGroup), &v1alpha1.PolicyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyGroups) UpdateStatus(ctx context.Context, policyGroup *v1alpha1.PolicyGroup, opts v1.UpdateOptions) (*v1alpha1.PolicyGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policygroupsResource, "status", c.ns, policyGroup), &v1alpha1.PolicyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyGroup), err
}

// Delete takes name of the policyGroup and deletes it. Returns an error if one occurs.
func (c *FakePolicyGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policygroupsResource, c.ns, name), &v1alpha1.PolicyGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policygroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyGroupList{})
	return err
}

// Patch applies the patch and returns the patched policyGroup.
func (c *FakePolicyGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policygroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyGroup), err
}

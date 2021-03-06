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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cms/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroupMetricRules implements GroupMetricRuleInterface
type FakeGroupMetricRules struct {
	Fake *FakeCmsV1alpha1
	ns   string
}

var groupmetricrulesResource = schema.GroupVersionResource{Group: "cms.alicloud.kubeform.com", Version: "v1alpha1", Resource: "groupmetricrules"}

var groupmetricrulesKind = schema.GroupVersionKind{Group: "cms.alicloud.kubeform.com", Version: "v1alpha1", Kind: "GroupMetricRule"}

// Get takes name of the groupMetricRule, and returns the corresponding groupMetricRule object, and an error if there is any.
func (c *FakeGroupMetricRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupMetricRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(groupmetricrulesResource, c.ns, name), &v1alpha1.GroupMetricRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMetricRule), err
}

// List takes label and field selectors, and returns the list of GroupMetricRules that match those selectors.
func (c *FakeGroupMetricRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupMetricRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(groupmetricrulesResource, groupmetricrulesKind, c.ns, opts), &v1alpha1.GroupMetricRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupMetricRuleList{ListMeta: obj.(*v1alpha1.GroupMetricRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupMetricRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupMetricRules.
func (c *FakeGroupMetricRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(groupmetricrulesResource, c.ns, opts))

}

// Create takes the representation of a groupMetricRule and creates it.  Returns the server's representation of the groupMetricRule, and an error, if there is any.
func (c *FakeGroupMetricRules) Create(ctx context.Context, groupMetricRule *v1alpha1.GroupMetricRule, opts v1.CreateOptions) (result *v1alpha1.GroupMetricRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(groupmetricrulesResource, c.ns, groupMetricRule), &v1alpha1.GroupMetricRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMetricRule), err
}

// Update takes the representation of a groupMetricRule and updates it. Returns the server's representation of the groupMetricRule, and an error, if there is any.
func (c *FakeGroupMetricRules) Update(ctx context.Context, groupMetricRule *v1alpha1.GroupMetricRule, opts v1.UpdateOptions) (result *v1alpha1.GroupMetricRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(groupmetricrulesResource, c.ns, groupMetricRule), &v1alpha1.GroupMetricRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMetricRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupMetricRules) UpdateStatus(ctx context.Context, groupMetricRule *v1alpha1.GroupMetricRule, opts v1.UpdateOptions) (*v1alpha1.GroupMetricRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(groupmetricrulesResource, "status", c.ns, groupMetricRule), &v1alpha1.GroupMetricRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMetricRule), err
}

// Delete takes name of the groupMetricRule and deletes it. Returns an error if one occurs.
func (c *FakeGroupMetricRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(groupmetricrulesResource, c.ns, name), &v1alpha1.GroupMetricRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupMetricRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(groupmetricrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupMetricRuleList{})
	return err
}

// Patch applies the patch and returns the patched groupMetricRule.
func (c *FakeGroupMetricRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupMetricRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(groupmetricrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupMetricRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMetricRule), err
}

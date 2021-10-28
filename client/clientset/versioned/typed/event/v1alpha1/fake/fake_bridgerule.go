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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/event/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBridgeRules implements BridgeRuleInterface
type FakeBridgeRules struct {
	Fake *FakeEventV1alpha1
	ns   string
}

var bridgerulesResource = schema.GroupVersionResource{Group: "event.alicloud.kubeform.com", Version: "v1alpha1", Resource: "bridgerules"}

var bridgerulesKind = schema.GroupVersionKind{Group: "event.alicloud.kubeform.com", Version: "v1alpha1", Kind: "BridgeRule"}

// Get takes name of the bridgeRule, and returns the corresponding bridgeRule object, and an error if there is any.
func (c *FakeBridgeRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BridgeRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bridgerulesResource, c.ns, name), &v1alpha1.BridgeRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BridgeRule), err
}

// List takes label and field selectors, and returns the list of BridgeRules that match those selectors.
func (c *FakeBridgeRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BridgeRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bridgerulesResource, bridgerulesKind, c.ns, opts), &v1alpha1.BridgeRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BridgeRuleList{ListMeta: obj.(*v1alpha1.BridgeRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.BridgeRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bridgeRules.
func (c *FakeBridgeRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bridgerulesResource, c.ns, opts))

}

// Create takes the representation of a bridgeRule and creates it.  Returns the server's representation of the bridgeRule, and an error, if there is any.
func (c *FakeBridgeRules) Create(ctx context.Context, bridgeRule *v1alpha1.BridgeRule, opts v1.CreateOptions) (result *v1alpha1.BridgeRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bridgerulesResource, c.ns, bridgeRule), &v1alpha1.BridgeRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BridgeRule), err
}

// Update takes the representation of a bridgeRule and updates it. Returns the server's representation of the bridgeRule, and an error, if there is any.
func (c *FakeBridgeRules) Update(ctx context.Context, bridgeRule *v1alpha1.BridgeRule, opts v1.UpdateOptions) (result *v1alpha1.BridgeRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bridgerulesResource, c.ns, bridgeRule), &v1alpha1.BridgeRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BridgeRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBridgeRules) UpdateStatus(ctx context.Context, bridgeRule *v1alpha1.BridgeRule, opts v1.UpdateOptions) (*v1alpha1.BridgeRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bridgerulesResource, "status", c.ns, bridgeRule), &v1alpha1.BridgeRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BridgeRule), err
}

// Delete takes name of the bridgeRule and deletes it. Returns an error if one occurs.
func (c *FakeBridgeRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bridgerulesResource, c.ns, name), &v1alpha1.BridgeRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBridgeRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bridgerulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BridgeRuleList{})
	return err
}

// Patch applies the patch and returns the patched bridgeRule.
func (c *FakeBridgeRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BridgeRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bridgerulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BridgeRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BridgeRule), err
}

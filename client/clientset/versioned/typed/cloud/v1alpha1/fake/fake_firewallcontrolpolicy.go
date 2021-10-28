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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFirewallControlPolicies implements FirewallControlPolicyInterface
type FakeFirewallControlPolicies struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var firewallcontrolpoliciesResource = schema.GroupVersionResource{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Resource: "firewallcontrolpolicies"}

var firewallcontrolpoliciesKind = schema.GroupVersionKind{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Kind: "FirewallControlPolicy"}

// Get takes name of the firewallControlPolicy, and returns the corresponding firewallControlPolicy object, and an error if there is any.
func (c *FakeFirewallControlPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirewallControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firewallcontrolpoliciesResource, c.ns, name), &v1alpha1.FirewallControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicy), err
}

// List takes label and field selectors, and returns the list of FirewallControlPolicies that match those selectors.
func (c *FakeFirewallControlPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirewallControlPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firewallcontrolpoliciesResource, firewallcontrolpoliciesKind, c.ns, opts), &v1alpha1.FirewallControlPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirewallControlPolicyList{ListMeta: obj.(*v1alpha1.FirewallControlPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirewallControlPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firewallControlPolicies.
func (c *FakeFirewallControlPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firewallcontrolpoliciesResource, c.ns, opts))

}

// Create takes the representation of a firewallControlPolicy and creates it.  Returns the server's representation of the firewallControlPolicy, and an error, if there is any.
func (c *FakeFirewallControlPolicies) Create(ctx context.Context, firewallControlPolicy *v1alpha1.FirewallControlPolicy, opts v1.CreateOptions) (result *v1alpha1.FirewallControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firewallcontrolpoliciesResource, c.ns, firewallControlPolicy), &v1alpha1.FirewallControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicy), err
}

// Update takes the representation of a firewallControlPolicy and updates it. Returns the server's representation of the firewallControlPolicy, and an error, if there is any.
func (c *FakeFirewallControlPolicies) Update(ctx context.Context, firewallControlPolicy *v1alpha1.FirewallControlPolicy, opts v1.UpdateOptions) (result *v1alpha1.FirewallControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firewallcontrolpoliciesResource, c.ns, firewallControlPolicy), &v1alpha1.FirewallControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirewallControlPolicies) UpdateStatus(ctx context.Context, firewallControlPolicy *v1alpha1.FirewallControlPolicy, opts v1.UpdateOptions) (*v1alpha1.FirewallControlPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firewallcontrolpoliciesResource, "status", c.ns, firewallControlPolicy), &v1alpha1.FirewallControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicy), err
}

// Delete takes name of the firewallControlPolicy and deletes it. Returns an error if one occurs.
func (c *FakeFirewallControlPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(firewallcontrolpoliciesResource, c.ns, name), &v1alpha1.FirewallControlPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirewallControlPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firewallcontrolpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirewallControlPolicyList{})
	return err
}

// Patch applies the patch and returns the patched firewallControlPolicy.
func (c *FakeFirewallControlPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirewallControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firewallcontrolpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirewallControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicy), err
}
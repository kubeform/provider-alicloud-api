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

// FakeFirewallControlPolicyOrders implements FirewallControlPolicyOrderInterface
type FakeFirewallControlPolicyOrders struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var firewallcontrolpolicyordersResource = schema.GroupVersionResource{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Resource: "firewallcontrolpolicyorders"}

var firewallcontrolpolicyordersKind = schema.GroupVersionKind{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Kind: "FirewallControlPolicyOrder"}

// Get takes name of the firewallControlPolicyOrder, and returns the corresponding firewallControlPolicyOrder object, and an error if there is any.
func (c *FakeFirewallControlPolicyOrders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirewallControlPolicyOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firewallcontrolpolicyordersResource, c.ns, name), &v1alpha1.FirewallControlPolicyOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicyOrder), err
}

// List takes label and field selectors, and returns the list of FirewallControlPolicyOrders that match those selectors.
func (c *FakeFirewallControlPolicyOrders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirewallControlPolicyOrderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firewallcontrolpolicyordersResource, firewallcontrolpolicyordersKind, c.ns, opts), &v1alpha1.FirewallControlPolicyOrderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirewallControlPolicyOrderList{ListMeta: obj.(*v1alpha1.FirewallControlPolicyOrderList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirewallControlPolicyOrderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firewallControlPolicyOrders.
func (c *FakeFirewallControlPolicyOrders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firewallcontrolpolicyordersResource, c.ns, opts))

}

// Create takes the representation of a firewallControlPolicyOrder and creates it.  Returns the server's representation of the firewallControlPolicyOrder, and an error, if there is any.
func (c *FakeFirewallControlPolicyOrders) Create(ctx context.Context, firewallControlPolicyOrder *v1alpha1.FirewallControlPolicyOrder, opts v1.CreateOptions) (result *v1alpha1.FirewallControlPolicyOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firewallcontrolpolicyordersResource, c.ns, firewallControlPolicyOrder), &v1alpha1.FirewallControlPolicyOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicyOrder), err
}

// Update takes the representation of a firewallControlPolicyOrder and updates it. Returns the server's representation of the firewallControlPolicyOrder, and an error, if there is any.
func (c *FakeFirewallControlPolicyOrders) Update(ctx context.Context, firewallControlPolicyOrder *v1alpha1.FirewallControlPolicyOrder, opts v1.UpdateOptions) (result *v1alpha1.FirewallControlPolicyOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firewallcontrolpolicyordersResource, c.ns, firewallControlPolicyOrder), &v1alpha1.FirewallControlPolicyOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicyOrder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirewallControlPolicyOrders) UpdateStatus(ctx context.Context, firewallControlPolicyOrder *v1alpha1.FirewallControlPolicyOrder, opts v1.UpdateOptions) (*v1alpha1.FirewallControlPolicyOrder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firewallcontrolpolicyordersResource, "status", c.ns, firewallControlPolicyOrder), &v1alpha1.FirewallControlPolicyOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicyOrder), err
}

// Delete takes name of the firewallControlPolicyOrder and deletes it. Returns an error if one occurs.
func (c *FakeFirewallControlPolicyOrders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(firewallcontrolpolicyordersResource, c.ns, name), &v1alpha1.FirewallControlPolicyOrder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirewallControlPolicyOrders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firewallcontrolpolicyordersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirewallControlPolicyOrderList{})
	return err
}

// Patch applies the patch and returns the patched firewallControlPolicyOrder.
func (c *FakeFirewallControlPolicyOrders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirewallControlPolicyOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firewallcontrolpolicyordersResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirewallControlPolicyOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallControlPolicyOrder), err
}

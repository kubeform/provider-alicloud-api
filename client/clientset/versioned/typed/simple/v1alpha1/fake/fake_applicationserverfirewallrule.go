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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/simple/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationServerFirewallRules implements ApplicationServerFirewallRuleInterface
type FakeApplicationServerFirewallRules struct {
	Fake *FakeSimpleV1alpha1
	ns   string
}

var applicationserverfirewallrulesResource = schema.GroupVersionResource{Group: "simple.alicloud.kubeform.com", Version: "v1alpha1", Resource: "applicationserverfirewallrules"}

var applicationserverfirewallrulesKind = schema.GroupVersionKind{Group: "simple.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ApplicationServerFirewallRule"}

// Get takes name of the applicationServerFirewallRule, and returns the corresponding applicationServerFirewallRule object, and an error if there is any.
func (c *FakeApplicationServerFirewallRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApplicationServerFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationserverfirewallrulesResource, c.ns, name), &v1alpha1.ApplicationServerFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationServerFirewallRule), err
}

// List takes label and field selectors, and returns the list of ApplicationServerFirewallRules that match those selectors.
func (c *FakeApplicationServerFirewallRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationServerFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationserverfirewallrulesResource, applicationserverfirewallrulesKind, c.ns, opts), &v1alpha1.ApplicationServerFirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationServerFirewallRuleList{ListMeta: obj.(*v1alpha1.ApplicationServerFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationServerFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationServerFirewallRules.
func (c *FakeApplicationServerFirewallRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationserverfirewallrulesResource, c.ns, opts))

}

// Create takes the representation of a applicationServerFirewallRule and creates it.  Returns the server's representation of the applicationServerFirewallRule, and an error, if there is any.
func (c *FakeApplicationServerFirewallRules) Create(ctx context.Context, applicationServerFirewallRule *v1alpha1.ApplicationServerFirewallRule, opts v1.CreateOptions) (result *v1alpha1.ApplicationServerFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationserverfirewallrulesResource, c.ns, applicationServerFirewallRule), &v1alpha1.ApplicationServerFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationServerFirewallRule), err
}

// Update takes the representation of a applicationServerFirewallRule and updates it. Returns the server's representation of the applicationServerFirewallRule, and an error, if there is any.
func (c *FakeApplicationServerFirewallRules) Update(ctx context.Context, applicationServerFirewallRule *v1alpha1.ApplicationServerFirewallRule, opts v1.UpdateOptions) (result *v1alpha1.ApplicationServerFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationserverfirewallrulesResource, c.ns, applicationServerFirewallRule), &v1alpha1.ApplicationServerFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationServerFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationServerFirewallRules) UpdateStatus(ctx context.Context, applicationServerFirewallRule *v1alpha1.ApplicationServerFirewallRule, opts v1.UpdateOptions) (*v1alpha1.ApplicationServerFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationserverfirewallrulesResource, "status", c.ns, applicationServerFirewallRule), &v1alpha1.ApplicationServerFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationServerFirewallRule), err
}

// Delete takes name of the applicationServerFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeApplicationServerFirewallRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationserverfirewallrulesResource, c.ns, name), &v1alpha1.ApplicationServerFirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationServerFirewallRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationserverfirewallrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationServerFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched applicationServerFirewallRule.
func (c *FakeApplicationServerFirewallRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationServerFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationserverfirewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationServerFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationServerFirewallRule), err
}

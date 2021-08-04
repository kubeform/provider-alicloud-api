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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ram/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccountPasswordPolicies implements AccountPasswordPolicyInterface
type FakeAccountPasswordPolicies struct {
	Fake *FakeRamV1alpha1
	ns   string
}

var accountpasswordpoliciesResource = schema.GroupVersionResource{Group: "ram.alicloud.kubeform.com", Version: "v1alpha1", Resource: "accountpasswordpolicies"}

var accountpasswordpoliciesKind = schema.GroupVersionKind{Group: "ram.alicloud.kubeform.com", Version: "v1alpha1", Kind: "AccountPasswordPolicy"}

// Get takes name of the accountPasswordPolicy, and returns the corresponding accountPasswordPolicy object, and an error if there is any.
func (c *FakeAccountPasswordPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountpasswordpoliciesResource, c.ns, name), &v1alpha1.AccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPasswordPolicy), err
}

// List takes label and field selectors, and returns the list of AccountPasswordPolicies that match those selectors.
func (c *FakeAccountPasswordPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountPasswordPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountpasswordpoliciesResource, accountpasswordpoliciesKind, c.ns, opts), &v1alpha1.AccountPasswordPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountPasswordPolicyList{ListMeta: obj.(*v1alpha1.AccountPasswordPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountPasswordPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accountPasswordPolicies.
func (c *FakeAccountPasswordPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountpasswordpoliciesResource, c.ns, opts))

}

// Create takes the representation of a accountPasswordPolicy and creates it.  Returns the server's representation of the accountPasswordPolicy, and an error, if there is any.
func (c *FakeAccountPasswordPolicies) Create(ctx context.Context, accountPasswordPolicy *v1alpha1.AccountPasswordPolicy, opts v1.CreateOptions) (result *v1alpha1.AccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountpasswordpoliciesResource, c.ns, accountPasswordPolicy), &v1alpha1.AccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPasswordPolicy), err
}

// Update takes the representation of a accountPasswordPolicy and updates it. Returns the server's representation of the accountPasswordPolicy, and an error, if there is any.
func (c *FakeAccountPasswordPolicies) Update(ctx context.Context, accountPasswordPolicy *v1alpha1.AccountPasswordPolicy, opts v1.UpdateOptions) (result *v1alpha1.AccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountpasswordpoliciesResource, c.ns, accountPasswordPolicy), &v1alpha1.AccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPasswordPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccountPasswordPolicies) UpdateStatus(ctx context.Context, accountPasswordPolicy *v1alpha1.AccountPasswordPolicy, opts v1.UpdateOptions) (*v1alpha1.AccountPasswordPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountpasswordpoliciesResource, "status", c.ns, accountPasswordPolicy), &v1alpha1.AccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPasswordPolicy), err
}

// Delete takes name of the accountPasswordPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAccountPasswordPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accountpasswordpoliciesResource, c.ns, name), &v1alpha1.AccountPasswordPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccountPasswordPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountpasswordpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountPasswordPolicyList{})
	return err
}

// Patch applies the patch and returns the patched accountPasswordPolicy.
func (c *FakeAccountPasswordPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountpasswordpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPasswordPolicy), err
}
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

// FakeManagerControlPolicies implements ManagerControlPolicyInterface
type FakeManagerControlPolicies struct {
	Fake *FakeResourceV1alpha1
	ns   string
}

var managercontrolpoliciesResource = schema.GroupVersionResource{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Resource: "managercontrolpolicies"}

var managercontrolpoliciesKind = schema.GroupVersionKind{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ManagerControlPolicy"}

// Get takes name of the managerControlPolicy, and returns the corresponding managerControlPolicy object, and an error if there is any.
func (c *FakeManagerControlPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managercontrolpoliciesResource, c.ns, name), &v1alpha1.ManagerControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerControlPolicy), err
}

// List takes label and field selectors, and returns the list of ManagerControlPolicies that match those selectors.
func (c *FakeManagerControlPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerControlPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managercontrolpoliciesResource, managercontrolpoliciesKind, c.ns, opts), &v1alpha1.ManagerControlPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerControlPolicyList{ListMeta: obj.(*v1alpha1.ManagerControlPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerControlPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerControlPolicies.
func (c *FakeManagerControlPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managercontrolpoliciesResource, c.ns, opts))

}

// Create takes the representation of a managerControlPolicy and creates it.  Returns the server's representation of the managerControlPolicy, and an error, if there is any.
func (c *FakeManagerControlPolicies) Create(ctx context.Context, managerControlPolicy *v1alpha1.ManagerControlPolicy, opts v1.CreateOptions) (result *v1alpha1.ManagerControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managercontrolpoliciesResource, c.ns, managerControlPolicy), &v1alpha1.ManagerControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerControlPolicy), err
}

// Update takes the representation of a managerControlPolicy and updates it. Returns the server's representation of the managerControlPolicy, and an error, if there is any.
func (c *FakeManagerControlPolicies) Update(ctx context.Context, managerControlPolicy *v1alpha1.ManagerControlPolicy, opts v1.UpdateOptions) (result *v1alpha1.ManagerControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managercontrolpoliciesResource, c.ns, managerControlPolicy), &v1alpha1.ManagerControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerControlPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerControlPolicies) UpdateStatus(ctx context.Context, managerControlPolicy *v1alpha1.ManagerControlPolicy, opts v1.UpdateOptions) (*v1alpha1.ManagerControlPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managercontrolpoliciesResource, "status", c.ns, managerControlPolicy), &v1alpha1.ManagerControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerControlPolicy), err
}

// Delete takes name of the managerControlPolicy and deletes it. Returns an error if one occurs.
func (c *FakeManagerControlPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managercontrolpoliciesResource, c.ns, name), &v1alpha1.ManagerControlPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerControlPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managercontrolpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerControlPolicyList{})
	return err
}

// Patch applies the patch and returns the patched managerControlPolicy.
func (c *FakeManagerControlPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerControlPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managercontrolpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerControlPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerControlPolicy), err
}

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

// FakeManagerAccounts implements ManagerAccountInterface
type FakeManagerAccounts struct {
	Fake *FakeResourceV1alpha1
	ns   string
}

var manageraccountsResource = schema.GroupVersionResource{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Resource: "manageraccounts"}

var manageraccountsKind = schema.GroupVersionKind{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ManagerAccount"}

// Get takes name of the managerAccount, and returns the corresponding managerAccount object, and an error if there is any.
func (c *FakeManagerAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(manageraccountsResource, c.ns, name), &v1alpha1.ManagerAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerAccount), err
}

// List takes label and field selectors, and returns the list of ManagerAccounts that match those selectors.
func (c *FakeManagerAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(manageraccountsResource, manageraccountsKind, c.ns, opts), &v1alpha1.ManagerAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerAccountList{ListMeta: obj.(*v1alpha1.ManagerAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerAccounts.
func (c *FakeManagerAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(manageraccountsResource, c.ns, opts))

}

// Create takes the representation of a managerAccount and creates it.  Returns the server's representation of the managerAccount, and an error, if there is any.
func (c *FakeManagerAccounts) Create(ctx context.Context, managerAccount *v1alpha1.ManagerAccount, opts v1.CreateOptions) (result *v1alpha1.ManagerAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(manageraccountsResource, c.ns, managerAccount), &v1alpha1.ManagerAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerAccount), err
}

// Update takes the representation of a managerAccount and updates it. Returns the server's representation of the managerAccount, and an error, if there is any.
func (c *FakeManagerAccounts) Update(ctx context.Context, managerAccount *v1alpha1.ManagerAccount, opts v1.UpdateOptions) (result *v1alpha1.ManagerAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(manageraccountsResource, c.ns, managerAccount), &v1alpha1.ManagerAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerAccounts) UpdateStatus(ctx context.Context, managerAccount *v1alpha1.ManagerAccount, opts v1.UpdateOptions) (*v1alpha1.ManagerAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(manageraccountsResource, "status", c.ns, managerAccount), &v1alpha1.ManagerAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerAccount), err
}

// Delete takes name of the managerAccount and deletes it. Returns an error if one occurs.
func (c *FakeManagerAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(manageraccountsResource, c.ns, name), &v1alpha1.ManagerAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(manageraccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerAccountList{})
	return err
}

// Patch applies the patch and returns the patched managerAccount.
func (c *FakeManagerAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(manageraccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerAccount), err
}

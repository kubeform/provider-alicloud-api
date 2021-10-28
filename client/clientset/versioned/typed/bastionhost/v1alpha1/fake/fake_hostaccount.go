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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/bastionhost/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostAccounts implements HostAccountInterface
type FakeHostAccounts struct {
	Fake *FakeBastionhostV1alpha1
	ns   string
}

var hostaccountsResource = schema.GroupVersionResource{Group: "bastionhost.alicloud.kubeform.com", Version: "v1alpha1", Resource: "hostaccounts"}

var hostaccountsKind = schema.GroupVersionKind{Group: "bastionhost.alicloud.kubeform.com", Version: "v1alpha1", Kind: "HostAccount"}

// Get takes name of the hostAccount, and returns the corresponding hostAccount object, and an error if there is any.
func (c *FakeHostAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostaccountsResource, c.ns, name), &v1alpha1.HostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostAccount), err
}

// List takes label and field selectors, and returns the list of HostAccounts that match those selectors.
func (c *FakeHostAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HostAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostaccountsResource, hostaccountsKind, c.ns, opts), &v1alpha1.HostAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HostAccountList{ListMeta: obj.(*v1alpha1.HostAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.HostAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostAccounts.
func (c *FakeHostAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostaccountsResource, c.ns, opts))

}

// Create takes the representation of a hostAccount and creates it.  Returns the server's representation of the hostAccount, and an error, if there is any.
func (c *FakeHostAccounts) Create(ctx context.Context, hostAccount *v1alpha1.HostAccount, opts v1.CreateOptions) (result *v1alpha1.HostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostaccountsResource, c.ns, hostAccount), &v1alpha1.HostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostAccount), err
}

// Update takes the representation of a hostAccount and updates it. Returns the server's representation of the hostAccount, and an error, if there is any.
func (c *FakeHostAccounts) Update(ctx context.Context, hostAccount *v1alpha1.HostAccount, opts v1.UpdateOptions) (result *v1alpha1.HostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostaccountsResource, c.ns, hostAccount), &v1alpha1.HostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHostAccounts) UpdateStatus(ctx context.Context, hostAccount *v1alpha1.HostAccount, opts v1.UpdateOptions) (*v1alpha1.HostAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hostaccountsResource, "status", c.ns, hostAccount), &v1alpha1.HostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostAccount), err
}

// Delete takes name of the hostAccount and deletes it. Returns an error if one occurs.
func (c *FakeHostAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hostaccountsResource, c.ns, name), &v1alpha1.HostAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HostAccountList{})
	return err
}

// Patch applies the patch and returns the patched hostAccount.
func (c *FakeHostAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostAccount), err
}
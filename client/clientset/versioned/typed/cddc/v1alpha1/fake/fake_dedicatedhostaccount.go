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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cddc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDedicatedHostAccounts implements DedicatedHostAccountInterface
type FakeDedicatedHostAccounts struct {
	Fake *FakeCddcV1alpha1
	ns   string
}

var dedicatedhostaccountsResource = schema.GroupVersionResource{Group: "cddc.alicloud.kubeform.com", Version: "v1alpha1", Resource: "dedicatedhostaccounts"}

var dedicatedhostaccountsKind = schema.GroupVersionKind{Group: "cddc.alicloud.kubeform.com", Version: "v1alpha1", Kind: "DedicatedHostAccount"}

// Get takes name of the dedicatedHostAccount, and returns the corresponding dedicatedHostAccount object, and an error if there is any.
func (c *FakeDedicatedHostAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DedicatedHostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dedicatedhostaccountsResource, c.ns, name), &v1alpha1.DedicatedHostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostAccount), err
}

// List takes label and field selectors, and returns the list of DedicatedHostAccounts that match those selectors.
func (c *FakeDedicatedHostAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DedicatedHostAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dedicatedhostaccountsResource, dedicatedhostaccountsKind, c.ns, opts), &v1alpha1.DedicatedHostAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DedicatedHostAccountList{ListMeta: obj.(*v1alpha1.DedicatedHostAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.DedicatedHostAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dedicatedHostAccounts.
func (c *FakeDedicatedHostAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dedicatedhostaccountsResource, c.ns, opts))

}

// Create takes the representation of a dedicatedHostAccount and creates it.  Returns the server's representation of the dedicatedHostAccount, and an error, if there is any.
func (c *FakeDedicatedHostAccounts) Create(ctx context.Context, dedicatedHostAccount *v1alpha1.DedicatedHostAccount, opts v1.CreateOptions) (result *v1alpha1.DedicatedHostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dedicatedhostaccountsResource, c.ns, dedicatedHostAccount), &v1alpha1.DedicatedHostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostAccount), err
}

// Update takes the representation of a dedicatedHostAccount and updates it. Returns the server's representation of the dedicatedHostAccount, and an error, if there is any.
func (c *FakeDedicatedHostAccounts) Update(ctx context.Context, dedicatedHostAccount *v1alpha1.DedicatedHostAccount, opts v1.UpdateOptions) (result *v1alpha1.DedicatedHostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dedicatedhostaccountsResource, c.ns, dedicatedHostAccount), &v1alpha1.DedicatedHostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDedicatedHostAccounts) UpdateStatus(ctx context.Context, dedicatedHostAccount *v1alpha1.DedicatedHostAccount, opts v1.UpdateOptions) (*v1alpha1.DedicatedHostAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dedicatedhostaccountsResource, "status", c.ns, dedicatedHostAccount), &v1alpha1.DedicatedHostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostAccount), err
}

// Delete takes name of the dedicatedHostAccount and deletes it. Returns an error if one occurs.
func (c *FakeDedicatedHostAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dedicatedhostaccountsResource, c.ns, name), &v1alpha1.DedicatedHostAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDedicatedHostAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dedicatedhostaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DedicatedHostAccountList{})
	return err
}

// Patch applies the patch and returns the patched dedicatedHostAccount.
func (c *FakeDedicatedHostAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DedicatedHostAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dedicatedhostaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DedicatedHostAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostAccount), err
}

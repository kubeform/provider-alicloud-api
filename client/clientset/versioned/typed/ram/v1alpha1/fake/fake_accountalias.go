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

// FakeAccountAliases implements AccountAliasInterface
type FakeAccountAliases struct {
	Fake *FakeRamV1alpha1
	ns   string
}

var accountaliasesResource = schema.GroupVersionResource{Group: "ram.alicloud.kubeform.com", Version: "v1alpha1", Resource: "accountaliases"}

var accountaliasesKind = schema.GroupVersionKind{Group: "ram.alicloud.kubeform.com", Version: "v1alpha1", Kind: "AccountAlias"}

// Get takes name of the accountAlias, and returns the corresponding accountAlias object, and an error if there is any.
func (c *FakeAccountAliases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccountAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountaliasesResource, c.ns, name), &v1alpha1.AccountAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountAlias), err
}

// List takes label and field selectors, and returns the list of AccountAliases that match those selectors.
func (c *FakeAccountAliases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountaliasesResource, accountaliasesKind, c.ns, opts), &v1alpha1.AccountAliasList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountAliasList{ListMeta: obj.(*v1alpha1.AccountAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accountAliases.
func (c *FakeAccountAliases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountaliasesResource, c.ns, opts))

}

// Create takes the representation of a accountAlias and creates it.  Returns the server's representation of the accountAlias, and an error, if there is any.
func (c *FakeAccountAliases) Create(ctx context.Context, accountAlias *v1alpha1.AccountAlias, opts v1.CreateOptions) (result *v1alpha1.AccountAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountaliasesResource, c.ns, accountAlias), &v1alpha1.AccountAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountAlias), err
}

// Update takes the representation of a accountAlias and updates it. Returns the server's representation of the accountAlias, and an error, if there is any.
func (c *FakeAccountAliases) Update(ctx context.Context, accountAlias *v1alpha1.AccountAlias, opts v1.UpdateOptions) (result *v1alpha1.AccountAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountaliasesResource, c.ns, accountAlias), &v1alpha1.AccountAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccountAliases) UpdateStatus(ctx context.Context, accountAlias *v1alpha1.AccountAlias, opts v1.UpdateOptions) (*v1alpha1.AccountAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountaliasesResource, "status", c.ns, accountAlias), &v1alpha1.AccountAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountAlias), err
}

// Delete takes name of the accountAlias and deletes it. Returns an error if one occurs.
func (c *FakeAccountAliases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accountaliasesResource, c.ns, name), &v1alpha1.AccountAlias{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccountAliases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountaliasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountAliasList{})
	return err
}

// Patch applies the patch and returns the patched accountAlias.
func (c *FakeAccountAliases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccountAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountaliasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccountAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountAlias), err
}
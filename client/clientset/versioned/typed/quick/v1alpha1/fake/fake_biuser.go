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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/quick/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBiUsers implements BiUserInterface
type FakeBiUsers struct {
	Fake *FakeQuickV1alpha1
	ns   string
}

var biusersResource = schema.GroupVersionResource{Group: "quick.alicloud.kubeform.com", Version: "v1alpha1", Resource: "biusers"}

var biusersKind = schema.GroupVersionKind{Group: "quick.alicloud.kubeform.com", Version: "v1alpha1", Kind: "BiUser"}

// Get takes name of the biUser, and returns the corresponding biUser object, and an error if there is any.
func (c *FakeBiUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BiUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(biusersResource, c.ns, name), &v1alpha1.BiUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BiUser), err
}

// List takes label and field selectors, and returns the list of BiUsers that match those selectors.
func (c *FakeBiUsers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BiUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(biusersResource, biusersKind, c.ns, opts), &v1alpha1.BiUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BiUserList{ListMeta: obj.(*v1alpha1.BiUserList).ListMeta}
	for _, item := range obj.(*v1alpha1.BiUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested biUsers.
func (c *FakeBiUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(biusersResource, c.ns, opts))

}

// Create takes the representation of a biUser and creates it.  Returns the server's representation of the biUser, and an error, if there is any.
func (c *FakeBiUsers) Create(ctx context.Context, biUser *v1alpha1.BiUser, opts v1.CreateOptions) (result *v1alpha1.BiUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(biusersResource, c.ns, biUser), &v1alpha1.BiUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BiUser), err
}

// Update takes the representation of a biUser and updates it. Returns the server's representation of the biUser, and an error, if there is any.
func (c *FakeBiUsers) Update(ctx context.Context, biUser *v1alpha1.BiUser, opts v1.UpdateOptions) (result *v1alpha1.BiUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(biusersResource, c.ns, biUser), &v1alpha1.BiUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BiUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBiUsers) UpdateStatus(ctx context.Context, biUser *v1alpha1.BiUser, opts v1.UpdateOptions) (*v1alpha1.BiUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(biusersResource, "status", c.ns, biUser), &v1alpha1.BiUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BiUser), err
}

// Delete takes name of the biUser and deletes it. Returns an error if one occurs.
func (c *FakeBiUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(biusersResource, c.ns, name), &v1alpha1.BiUser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBiUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(biusersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BiUserList{})
	return err
}

// Patch applies the patch and returns the patched biUser.
func (c *FakeBiUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BiUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(biusersResource, c.ns, name, pt, data, subresources...), &v1alpha1.BiUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BiUser), err
}

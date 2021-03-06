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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/kms/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiphertexts implements CiphertextInterface
type FakeCiphertexts struct {
	Fake *FakeKmsV1alpha1
	ns   string
}

var ciphertextsResource = schema.GroupVersionResource{Group: "kms.alicloud.kubeform.com", Version: "v1alpha1", Resource: "ciphertexts"}

var ciphertextsKind = schema.GroupVersionKind{Group: "kms.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Ciphertext"}

// Get takes name of the ciphertext, and returns the corresponding ciphertext object, and an error if there is any.
func (c *FakeCiphertexts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ciphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ciphertextsResource, c.ns, name), &v1alpha1.Ciphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ciphertext), err
}

// List takes label and field selectors, and returns the list of Ciphertexts that match those selectors.
func (c *FakeCiphertexts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CiphertextList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ciphertextsResource, ciphertextsKind, c.ns, opts), &v1alpha1.CiphertextList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CiphertextList{ListMeta: obj.(*v1alpha1.CiphertextList).ListMeta}
	for _, item := range obj.(*v1alpha1.CiphertextList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciphertexts.
func (c *FakeCiphertexts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ciphertextsResource, c.ns, opts))

}

// Create takes the representation of a ciphertext and creates it.  Returns the server's representation of the ciphertext, and an error, if there is any.
func (c *FakeCiphertexts) Create(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.CreateOptions) (result *v1alpha1.Ciphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ciphertextsResource, c.ns, ciphertext), &v1alpha1.Ciphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ciphertext), err
}

// Update takes the representation of a ciphertext and updates it. Returns the server's representation of the ciphertext, and an error, if there is any.
func (c *FakeCiphertexts) Update(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.UpdateOptions) (result *v1alpha1.Ciphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ciphertextsResource, c.ns, ciphertext), &v1alpha1.Ciphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ciphertext), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCiphertexts) UpdateStatus(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.UpdateOptions) (*v1alpha1.Ciphertext, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ciphertextsResource, "status", c.ns, ciphertext), &v1alpha1.Ciphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ciphertext), err
}

// Delete takes name of the ciphertext and deletes it. Returns an error if one occurs.
func (c *FakeCiphertexts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ciphertextsResource, c.ns, name), &v1alpha1.Ciphertext{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiphertexts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ciphertextsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CiphertextList{})
	return err
}

// Patch applies the patch and returns the patched ciphertext.
func (c *FakeCiphertexts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ciphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ciphertextsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ciphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ciphertext), err
}

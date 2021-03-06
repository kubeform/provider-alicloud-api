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

// FakeSsoScimServerCredentials implements SsoScimServerCredentialInterface
type FakeSsoScimServerCredentials struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var ssoscimservercredentialsResource = schema.GroupVersionResource{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Resource: "ssoscimservercredentials"}

var ssoscimservercredentialsKind = schema.GroupVersionKind{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Kind: "SsoScimServerCredential"}

// Get takes name of the ssoScimServerCredential, and returns the corresponding ssoScimServerCredential object, and an error if there is any.
func (c *FakeSsoScimServerCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SsoScimServerCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssoscimservercredentialsResource, c.ns, name), &v1alpha1.SsoScimServerCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoScimServerCredential), err
}

// List takes label and field selectors, and returns the list of SsoScimServerCredentials that match those selectors.
func (c *FakeSsoScimServerCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SsoScimServerCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssoscimservercredentialsResource, ssoscimservercredentialsKind, c.ns, opts), &v1alpha1.SsoScimServerCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsoScimServerCredentialList{ListMeta: obj.(*v1alpha1.SsoScimServerCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsoScimServerCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssoScimServerCredentials.
func (c *FakeSsoScimServerCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssoscimservercredentialsResource, c.ns, opts))

}

// Create takes the representation of a ssoScimServerCredential and creates it.  Returns the server's representation of the ssoScimServerCredential, and an error, if there is any.
func (c *FakeSsoScimServerCredentials) Create(ctx context.Context, ssoScimServerCredential *v1alpha1.SsoScimServerCredential, opts v1.CreateOptions) (result *v1alpha1.SsoScimServerCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssoscimservercredentialsResource, c.ns, ssoScimServerCredential), &v1alpha1.SsoScimServerCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoScimServerCredential), err
}

// Update takes the representation of a ssoScimServerCredential and updates it. Returns the server's representation of the ssoScimServerCredential, and an error, if there is any.
func (c *FakeSsoScimServerCredentials) Update(ctx context.Context, ssoScimServerCredential *v1alpha1.SsoScimServerCredential, opts v1.UpdateOptions) (result *v1alpha1.SsoScimServerCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssoscimservercredentialsResource, c.ns, ssoScimServerCredential), &v1alpha1.SsoScimServerCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoScimServerCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsoScimServerCredentials) UpdateStatus(ctx context.Context, ssoScimServerCredential *v1alpha1.SsoScimServerCredential, opts v1.UpdateOptions) (*v1alpha1.SsoScimServerCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssoscimservercredentialsResource, "status", c.ns, ssoScimServerCredential), &v1alpha1.SsoScimServerCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoScimServerCredential), err
}

// Delete takes name of the ssoScimServerCredential and deletes it. Returns an error if one occurs.
func (c *FakeSsoScimServerCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssoscimservercredentialsResource, c.ns, name), &v1alpha1.SsoScimServerCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsoScimServerCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssoscimservercredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsoScimServerCredentialList{})
	return err
}

// Patch applies the patch and returns the patched ssoScimServerCredential.
func (c *FakeSsoScimServerCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SsoScimServerCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssoscimservercredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsoScimServerCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoScimServerCredential), err
}

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

// FakeSsoDirectories implements SsoDirectoryInterface
type FakeSsoDirectories struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var ssodirectoriesResource = schema.GroupVersionResource{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Resource: "ssodirectories"}

var ssodirectoriesKind = schema.GroupVersionKind{Group: "cloud.alicloud.kubeform.com", Version: "v1alpha1", Kind: "SsoDirectory"}

// Get takes name of the ssoDirectory, and returns the corresponding ssoDirectory object, and an error if there is any.
func (c *FakeSsoDirectories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SsoDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssodirectoriesResource, c.ns, name), &v1alpha1.SsoDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoDirectory), err
}

// List takes label and field selectors, and returns the list of SsoDirectories that match those selectors.
func (c *FakeSsoDirectories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SsoDirectoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssodirectoriesResource, ssodirectoriesKind, c.ns, opts), &v1alpha1.SsoDirectoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsoDirectoryList{ListMeta: obj.(*v1alpha1.SsoDirectoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsoDirectoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssoDirectories.
func (c *FakeSsoDirectories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssodirectoriesResource, c.ns, opts))

}

// Create takes the representation of a ssoDirectory and creates it.  Returns the server's representation of the ssoDirectory, and an error, if there is any.
func (c *FakeSsoDirectories) Create(ctx context.Context, ssoDirectory *v1alpha1.SsoDirectory, opts v1.CreateOptions) (result *v1alpha1.SsoDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssodirectoriesResource, c.ns, ssoDirectory), &v1alpha1.SsoDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoDirectory), err
}

// Update takes the representation of a ssoDirectory and updates it. Returns the server's representation of the ssoDirectory, and an error, if there is any.
func (c *FakeSsoDirectories) Update(ctx context.Context, ssoDirectory *v1alpha1.SsoDirectory, opts v1.UpdateOptions) (result *v1alpha1.SsoDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssodirectoriesResource, c.ns, ssoDirectory), &v1alpha1.SsoDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoDirectory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsoDirectories) UpdateStatus(ctx context.Context, ssoDirectory *v1alpha1.SsoDirectory, opts v1.UpdateOptions) (*v1alpha1.SsoDirectory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssodirectoriesResource, "status", c.ns, ssoDirectory), &v1alpha1.SsoDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoDirectory), err
}

// Delete takes name of the ssoDirectory and deletes it. Returns an error if one occurs.
func (c *FakeSsoDirectories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssodirectoriesResource, c.ns, name), &v1alpha1.SsoDirectory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsoDirectories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssodirectoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsoDirectoryList{})
	return err
}

// Patch applies the patch and returns the patched ssoDirectory.
func (c *FakeSsoDirectories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SsoDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssodirectoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsoDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsoDirectory), err
}
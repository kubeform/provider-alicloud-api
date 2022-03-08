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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecd/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkPackages implements NetworkPackageInterface
type FakeNetworkPackages struct {
	Fake *FakeEcdV1alpha1
	ns   string
}

var networkpackagesResource = schema.GroupVersionResource{Group: "ecd.alicloud.kubeform.com", Version: "v1alpha1", Resource: "networkpackages"}

var networkpackagesKind = schema.GroupVersionKind{Group: "ecd.alicloud.kubeform.com", Version: "v1alpha1", Kind: "NetworkPackage"}

// Get takes name of the networkPackage, and returns the corresponding networkPackage object, and an error if there is any.
func (c *FakeNetworkPackages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkpackagesResource, c.ns, name), &v1alpha1.NetworkPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPackage), err
}

// List takes label and field selectors, and returns the list of NetworkPackages that match those selectors.
func (c *FakeNetworkPackages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkPackageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkpackagesResource, networkpackagesKind, c.ns, opts), &v1alpha1.NetworkPackageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkPackageList{ListMeta: obj.(*v1alpha1.NetworkPackageList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkPackageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPackages.
func (c *FakeNetworkPackages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkpackagesResource, c.ns, opts))

}

// Create takes the representation of a networkPackage and creates it.  Returns the server's representation of the networkPackage, and an error, if there is any.
func (c *FakeNetworkPackages) Create(ctx context.Context, networkPackage *v1alpha1.NetworkPackage, opts v1.CreateOptions) (result *v1alpha1.NetworkPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkpackagesResource, c.ns, networkPackage), &v1alpha1.NetworkPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPackage), err
}

// Update takes the representation of a networkPackage and updates it. Returns the server's representation of the networkPackage, and an error, if there is any.
func (c *FakeNetworkPackages) Update(ctx context.Context, networkPackage *v1alpha1.NetworkPackage, opts v1.UpdateOptions) (result *v1alpha1.NetworkPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkpackagesResource, c.ns, networkPackage), &v1alpha1.NetworkPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPackage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPackages) UpdateStatus(ctx context.Context, networkPackage *v1alpha1.NetworkPackage, opts v1.UpdateOptions) (*v1alpha1.NetworkPackage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkpackagesResource, "status", c.ns, networkPackage), &v1alpha1.NetworkPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPackage), err
}

// Delete takes name of the networkPackage and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPackages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkpackagesResource, c.ns, name), &v1alpha1.NetworkPackage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPackages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkpackagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkPackageList{})
	return err
}

// Patch applies the patch and returns the patched networkPackage.
func (c *FakeNetworkPackages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkPackage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkpackagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkPackage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPackage), err
}

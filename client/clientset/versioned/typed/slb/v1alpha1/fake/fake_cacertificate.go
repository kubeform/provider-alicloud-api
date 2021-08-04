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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/slb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCaCertificates implements CaCertificateInterface
type FakeCaCertificates struct {
	Fake *FakeSlbV1alpha1
	ns   string
}

var cacertificatesResource = schema.GroupVersionResource{Group: "slb.alicloud.kubeform.com", Version: "v1alpha1", Resource: "cacertificates"}

var cacertificatesKind = schema.GroupVersionKind{Group: "slb.alicloud.kubeform.com", Version: "v1alpha1", Kind: "CaCertificate"}

// Get takes name of the caCertificate, and returns the corresponding caCertificate object, and an error if there is any.
func (c *FakeCaCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CaCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cacertificatesResource, c.ns, name), &v1alpha1.CaCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CaCertificate), err
}

// List takes label and field selectors, and returns the list of CaCertificates that match those selectors.
func (c *FakeCaCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CaCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cacertificatesResource, cacertificatesKind, c.ns, opts), &v1alpha1.CaCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CaCertificateList{ListMeta: obj.(*v1alpha1.CaCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.CaCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested caCertificates.
func (c *FakeCaCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cacertificatesResource, c.ns, opts))

}

// Create takes the representation of a caCertificate and creates it.  Returns the server's representation of the caCertificate, and an error, if there is any.
func (c *FakeCaCertificates) Create(ctx context.Context, caCertificate *v1alpha1.CaCertificate, opts v1.CreateOptions) (result *v1alpha1.CaCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cacertificatesResource, c.ns, caCertificate), &v1alpha1.CaCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CaCertificate), err
}

// Update takes the representation of a caCertificate and updates it. Returns the server's representation of the caCertificate, and an error, if there is any.
func (c *FakeCaCertificates) Update(ctx context.Context, caCertificate *v1alpha1.CaCertificate, opts v1.UpdateOptions) (result *v1alpha1.CaCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cacertificatesResource, c.ns, caCertificate), &v1alpha1.CaCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CaCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCaCertificates) UpdateStatus(ctx context.Context, caCertificate *v1alpha1.CaCertificate, opts v1.UpdateOptions) (*v1alpha1.CaCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cacertificatesResource, "status", c.ns, caCertificate), &v1alpha1.CaCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CaCertificate), err
}

// Delete takes name of the caCertificate and deletes it. Returns an error if one occurs.
func (c *FakeCaCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cacertificatesResource, c.ns, name), &v1alpha1.CaCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCaCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cacertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CaCertificateList{})
	return err
}

// Patch applies the patch and returns the patched caCertificate.
func (c *FakeCaCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CaCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cacertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CaCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CaCertificate), err
}
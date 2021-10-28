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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ssl/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificatesServiceCertificates implements CertificatesServiceCertificateInterface
type FakeCertificatesServiceCertificates struct {
	Fake *FakeSslV1alpha1
	ns   string
}

var certificatesservicecertificatesResource = schema.GroupVersionResource{Group: "ssl.alicloud.kubeform.com", Version: "v1alpha1", Resource: "certificatesservicecertificates"}

var certificatesservicecertificatesKind = schema.GroupVersionKind{Group: "ssl.alicloud.kubeform.com", Version: "v1alpha1", Kind: "CertificatesServiceCertificate"}

// Get takes name of the certificatesServiceCertificate, and returns the corresponding certificatesServiceCertificate object, and an error if there is any.
func (c *FakeCertificatesServiceCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificatesServiceCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificatesservicecertificatesResource, c.ns, name), &v1alpha1.CertificatesServiceCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificatesServiceCertificate), err
}

// List takes label and field selectors, and returns the list of CertificatesServiceCertificates that match those selectors.
func (c *FakeCertificatesServiceCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificatesServiceCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificatesservicecertificatesResource, certificatesservicecertificatesKind, c.ns, opts), &v1alpha1.CertificatesServiceCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificatesServiceCertificateList{ListMeta: obj.(*v1alpha1.CertificatesServiceCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificatesServiceCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificatesServiceCertificates.
func (c *FakeCertificatesServiceCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificatesservicecertificatesResource, c.ns, opts))

}

// Create takes the representation of a certificatesServiceCertificate and creates it.  Returns the server's representation of the certificatesServiceCertificate, and an error, if there is any.
func (c *FakeCertificatesServiceCertificates) Create(ctx context.Context, certificatesServiceCertificate *v1alpha1.CertificatesServiceCertificate, opts v1.CreateOptions) (result *v1alpha1.CertificatesServiceCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificatesservicecertificatesResource, c.ns, certificatesServiceCertificate), &v1alpha1.CertificatesServiceCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificatesServiceCertificate), err
}

// Update takes the representation of a certificatesServiceCertificate and updates it. Returns the server's representation of the certificatesServiceCertificate, and an error, if there is any.
func (c *FakeCertificatesServiceCertificates) Update(ctx context.Context, certificatesServiceCertificate *v1alpha1.CertificatesServiceCertificate, opts v1.UpdateOptions) (result *v1alpha1.CertificatesServiceCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificatesservicecertificatesResource, c.ns, certificatesServiceCertificate), &v1alpha1.CertificatesServiceCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificatesServiceCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificatesServiceCertificates) UpdateStatus(ctx context.Context, certificatesServiceCertificate *v1alpha1.CertificatesServiceCertificate, opts v1.UpdateOptions) (*v1alpha1.CertificatesServiceCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificatesservicecertificatesResource, "status", c.ns, certificatesServiceCertificate), &v1alpha1.CertificatesServiceCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificatesServiceCertificate), err
}

// Delete takes name of the certificatesServiceCertificate and deletes it. Returns an error if one occurs.
func (c *FakeCertificatesServiceCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificatesservicecertificatesResource, c.ns, name), &v1alpha1.CertificatesServiceCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificatesServiceCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificatesservicecertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificatesServiceCertificateList{})
	return err
}

// Patch applies the patch and returns the patched certificatesServiceCertificate.
func (c *FakeCertificatesServiceCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificatesServiceCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatesservicecertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CertificatesServiceCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificatesServiceCertificate), err
}
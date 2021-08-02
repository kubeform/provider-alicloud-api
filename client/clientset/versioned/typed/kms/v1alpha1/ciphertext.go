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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/kms/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiphertextsGetter has a method to return a CiphertextInterface.
// A group's client should implement this interface.
type CiphertextsGetter interface {
	Ciphertexts(namespace string) CiphertextInterface
}

// CiphertextInterface has methods to work with Ciphertext resources.
type CiphertextInterface interface {
	Create(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.CreateOptions) (*v1alpha1.Ciphertext, error)
	Update(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.UpdateOptions) (*v1alpha1.Ciphertext, error)
	UpdateStatus(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.UpdateOptions) (*v1alpha1.Ciphertext, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Ciphertext, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CiphertextList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ciphertext, err error)
	CiphertextExpansion
}

// ciphertexts implements CiphertextInterface
type ciphertexts struct {
	client rest.Interface
	ns     string
}

// newCiphertexts returns a Ciphertexts
func newCiphertexts(c *KmsV1alpha1Client, namespace string) *ciphertexts {
	return &ciphertexts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ciphertext, and returns the corresponding ciphertext object, and an error if there is any.
func (c *ciphertexts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ciphertext, err error) {
	result = &v1alpha1.Ciphertext{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ciphertexts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ciphertexts that match those selectors.
func (c *ciphertexts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CiphertextList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CiphertextList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ciphertexts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciphertexts.
func (c *ciphertexts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ciphertexts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciphertext and creates it.  Returns the server's representation of the ciphertext, and an error, if there is any.
func (c *ciphertexts) Create(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.CreateOptions) (result *v1alpha1.Ciphertext, err error) {
	result = &v1alpha1.Ciphertext{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ciphertexts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciphertext).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciphertext and updates it. Returns the server's representation of the ciphertext, and an error, if there is any.
func (c *ciphertexts) Update(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.UpdateOptions) (result *v1alpha1.Ciphertext, err error) {
	result = &v1alpha1.Ciphertext{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ciphertexts").
		Name(ciphertext.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciphertext).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ciphertexts) UpdateStatus(ctx context.Context, ciphertext *v1alpha1.Ciphertext, opts v1.UpdateOptions) (result *v1alpha1.Ciphertext, err error) {
	result = &v1alpha1.Ciphertext{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ciphertexts").
		Name(ciphertext.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciphertext).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciphertext and deletes it. Returns an error if one occurs.
func (c *ciphertexts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ciphertexts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciphertexts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ciphertexts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciphertext.
func (c *ciphertexts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ciphertext, err error) {
	result = &v1alpha1.Ciphertext{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ciphertexts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

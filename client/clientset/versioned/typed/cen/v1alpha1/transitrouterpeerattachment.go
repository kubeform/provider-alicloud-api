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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cen/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TransitRouterPeerAttachmentsGetter has a method to return a TransitRouterPeerAttachmentInterface.
// A group's client should implement this interface.
type TransitRouterPeerAttachmentsGetter interface {
	TransitRouterPeerAttachments(namespace string) TransitRouterPeerAttachmentInterface
}

// TransitRouterPeerAttachmentInterface has methods to work with TransitRouterPeerAttachment resources.
type TransitRouterPeerAttachmentInterface interface {
	Create(ctx context.Context, transitRouterPeerAttachment *v1alpha1.TransitRouterPeerAttachment, opts v1.CreateOptions) (*v1alpha1.TransitRouterPeerAttachment, error)
	Update(ctx context.Context, transitRouterPeerAttachment *v1alpha1.TransitRouterPeerAttachment, opts v1.UpdateOptions) (*v1alpha1.TransitRouterPeerAttachment, error)
	UpdateStatus(ctx context.Context, transitRouterPeerAttachment *v1alpha1.TransitRouterPeerAttachment, opts v1.UpdateOptions) (*v1alpha1.TransitRouterPeerAttachment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TransitRouterPeerAttachment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TransitRouterPeerAttachmentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitRouterPeerAttachment, err error)
	TransitRouterPeerAttachmentExpansion
}

// transitRouterPeerAttachments implements TransitRouterPeerAttachmentInterface
type transitRouterPeerAttachments struct {
	client rest.Interface
	ns     string
}

// newTransitRouterPeerAttachments returns a TransitRouterPeerAttachments
func newTransitRouterPeerAttachments(c *CenV1alpha1Client, namespace string) *transitRouterPeerAttachments {
	return &transitRouterPeerAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the transitRouterPeerAttachment, and returns the corresponding transitRouterPeerAttachment object, and an error if there is any.
func (c *transitRouterPeerAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitRouterPeerAttachment, err error) {
	result = &v1alpha1.TransitRouterPeerAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TransitRouterPeerAttachments that match those selectors.
func (c *transitRouterPeerAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitRouterPeerAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TransitRouterPeerAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested transitRouterPeerAttachments.
func (c *transitRouterPeerAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a transitRouterPeerAttachment and creates it.  Returns the server's representation of the transitRouterPeerAttachment, and an error, if there is any.
func (c *transitRouterPeerAttachments) Create(ctx context.Context, transitRouterPeerAttachment *v1alpha1.TransitRouterPeerAttachment, opts v1.CreateOptions) (result *v1alpha1.TransitRouterPeerAttachment, err error) {
	result = &v1alpha1.TransitRouterPeerAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transitRouterPeerAttachment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a transitRouterPeerAttachment and updates it. Returns the server's representation of the transitRouterPeerAttachment, and an error, if there is any.
func (c *transitRouterPeerAttachments) Update(ctx context.Context, transitRouterPeerAttachment *v1alpha1.TransitRouterPeerAttachment, opts v1.UpdateOptions) (result *v1alpha1.TransitRouterPeerAttachment, err error) {
	result = &v1alpha1.TransitRouterPeerAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		Name(transitRouterPeerAttachment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transitRouterPeerAttachment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *transitRouterPeerAttachments) UpdateStatus(ctx context.Context, transitRouterPeerAttachment *v1alpha1.TransitRouterPeerAttachment, opts v1.UpdateOptions) (result *v1alpha1.TransitRouterPeerAttachment, err error) {
	result = &v1alpha1.TransitRouterPeerAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		Name(transitRouterPeerAttachment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transitRouterPeerAttachment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the transitRouterPeerAttachment and deletes it. Returns an error if one occurs.
func (c *transitRouterPeerAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *transitRouterPeerAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched transitRouterPeerAttachment.
func (c *transitRouterPeerAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitRouterPeerAttachment, err error) {
	result = &v1alpha1.TransitRouterPeerAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("transitrouterpeerattachments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

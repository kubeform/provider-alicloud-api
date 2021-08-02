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

// BandwidthLimitsGetter has a method to return a BandwidthLimitInterface.
// A group's client should implement this interface.
type BandwidthLimitsGetter interface {
	BandwidthLimits(namespace string) BandwidthLimitInterface
}

// BandwidthLimitInterface has methods to work with BandwidthLimit resources.
type BandwidthLimitInterface interface {
	Create(ctx context.Context, bandwidthLimit *v1alpha1.BandwidthLimit, opts v1.CreateOptions) (*v1alpha1.BandwidthLimit, error)
	Update(ctx context.Context, bandwidthLimit *v1alpha1.BandwidthLimit, opts v1.UpdateOptions) (*v1alpha1.BandwidthLimit, error)
	UpdateStatus(ctx context.Context, bandwidthLimit *v1alpha1.BandwidthLimit, opts v1.UpdateOptions) (*v1alpha1.BandwidthLimit, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BandwidthLimit, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BandwidthLimitList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BandwidthLimit, err error)
	BandwidthLimitExpansion
}

// bandwidthLimits implements BandwidthLimitInterface
type bandwidthLimits struct {
	client rest.Interface
	ns     string
}

// newBandwidthLimits returns a BandwidthLimits
func newBandwidthLimits(c *CenV1alpha1Client, namespace string) *bandwidthLimits {
	return &bandwidthLimits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bandwidthLimit, and returns the corresponding bandwidthLimit object, and an error if there is any.
func (c *bandwidthLimits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BandwidthLimit, err error) {
	result = &v1alpha1.BandwidthLimit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BandwidthLimits that match those selectors.
func (c *bandwidthLimits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BandwidthLimitList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BandwidthLimitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bandwidthLimits.
func (c *bandwidthLimits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bandwidthLimit and creates it.  Returns the server's representation of the bandwidthLimit, and an error, if there is any.
func (c *bandwidthLimits) Create(ctx context.Context, bandwidthLimit *v1alpha1.BandwidthLimit, opts v1.CreateOptions) (result *v1alpha1.BandwidthLimit, err error) {
	result = &v1alpha1.BandwidthLimit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bandwidthLimit).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bandwidthLimit and updates it. Returns the server's representation of the bandwidthLimit, and an error, if there is any.
func (c *bandwidthLimits) Update(ctx context.Context, bandwidthLimit *v1alpha1.BandwidthLimit, opts v1.UpdateOptions) (result *v1alpha1.BandwidthLimit, err error) {
	result = &v1alpha1.BandwidthLimit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		Name(bandwidthLimit.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bandwidthLimit).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bandwidthLimits) UpdateStatus(ctx context.Context, bandwidthLimit *v1alpha1.BandwidthLimit, opts v1.UpdateOptions) (result *v1alpha1.BandwidthLimit, err error) {
	result = &v1alpha1.BandwidthLimit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		Name(bandwidthLimit.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bandwidthLimit).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bandwidthLimit and deletes it. Returns an error if one occurs.
func (c *bandwidthLimits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bandwidthLimits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bandwidthlimits").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bandwidthLimit.
func (c *bandwidthLimits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BandwidthLimit, err error) {
	result = &v1alpha1.BandwidthLimit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bandwidthlimits").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

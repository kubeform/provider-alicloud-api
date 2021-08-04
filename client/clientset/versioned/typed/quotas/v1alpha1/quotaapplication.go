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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/quotas/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QuotaApplicationsGetter has a method to return a QuotaApplicationInterface.
// A group's client should implement this interface.
type QuotaApplicationsGetter interface {
	QuotaApplications(namespace string) QuotaApplicationInterface
}

// QuotaApplicationInterface has methods to work with QuotaApplication resources.
type QuotaApplicationInterface interface {
	Create(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.CreateOptions) (*v1alpha1.QuotaApplication, error)
	Update(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.UpdateOptions) (*v1alpha1.QuotaApplication, error)
	UpdateStatus(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.UpdateOptions) (*v1alpha1.QuotaApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.QuotaApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.QuotaApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.QuotaApplication, err error)
	QuotaApplicationExpansion
}

// quotaApplications implements QuotaApplicationInterface
type quotaApplications struct {
	client rest.Interface
	ns     string
}

// newQuotaApplications returns a QuotaApplications
func newQuotaApplications(c *QuotasV1alpha1Client, namespace string) *quotaApplications {
	return &quotaApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the quotaApplication, and returns the corresponding quotaApplication object, and an error if there is any.
func (c *quotaApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.QuotaApplication, err error) {
	result = &v1alpha1.QuotaApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QuotaApplications that match those selectors.
func (c *quotaApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.QuotaApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.QuotaApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested quotaApplications.
func (c *quotaApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("quotaapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a quotaApplication and creates it.  Returns the server's representation of the quotaApplication, and an error, if there is any.
func (c *quotaApplications) Create(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.CreateOptions) (result *v1alpha1.QuotaApplication, err error) {
	result = &v1alpha1.QuotaApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("quotaapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(quotaApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a quotaApplication and updates it. Returns the server's representation of the quotaApplication, and an error, if there is any.
func (c *quotaApplications) Update(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.UpdateOptions) (result *v1alpha1.QuotaApplication, err error) {
	result = &v1alpha1.QuotaApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quotaapplications").
		Name(quotaApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(quotaApplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *quotaApplications) UpdateStatus(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.UpdateOptions) (result *v1alpha1.QuotaApplication, err error) {
	result = &v1alpha1.QuotaApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quotaapplications").
		Name(quotaApplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(quotaApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the quotaApplication and deletes it. Returns an error if one occurs.
func (c *quotaApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *quotaApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched quotaApplication.
func (c *quotaApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.QuotaApplication, err error) {
	result = &v1alpha1.QuotaApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("quotaapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
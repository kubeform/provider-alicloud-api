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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/log/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuditsGetter has a method to return a AuditInterface.
// A group's client should implement this interface.
type AuditsGetter interface {
	Audits(namespace string) AuditInterface
}

// AuditInterface has methods to work with Audit resources.
type AuditInterface interface {
	Create(ctx context.Context, audit *v1alpha1.Audit, opts v1.CreateOptions) (*v1alpha1.Audit, error)
	Update(ctx context.Context, audit *v1alpha1.Audit, opts v1.UpdateOptions) (*v1alpha1.Audit, error)
	UpdateStatus(ctx context.Context, audit *v1alpha1.Audit, opts v1.UpdateOptions) (*v1alpha1.Audit, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Audit, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AuditList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Audit, err error)
	AuditExpansion
}

// audits implements AuditInterface
type audits struct {
	client rest.Interface
	ns     string
}

// newAudits returns a Audits
func newAudits(c *LogV1alpha1Client, namespace string) *audits {
	return &audits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the audit, and returns the corresponding audit object, and an error if there is any.
func (c *audits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Audit, err error) {
	result = &v1alpha1.Audit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("audits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Audits that match those selectors.
func (c *audits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AuditList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AuditList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("audits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested audits.
func (c *audits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("audits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a audit and creates it.  Returns the server's representation of the audit, and an error, if there is any.
func (c *audits) Create(ctx context.Context, audit *v1alpha1.Audit, opts v1.CreateOptions) (result *v1alpha1.Audit, err error) {
	result = &v1alpha1.Audit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("audits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(audit).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a audit and updates it. Returns the server's representation of the audit, and an error, if there is any.
func (c *audits) Update(ctx context.Context, audit *v1alpha1.Audit, opts v1.UpdateOptions) (result *v1alpha1.Audit, err error) {
	result = &v1alpha1.Audit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("audits").
		Name(audit.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(audit).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *audits) UpdateStatus(ctx context.Context, audit *v1alpha1.Audit, opts v1.UpdateOptions) (result *v1alpha1.Audit, err error) {
	result = &v1alpha1.Audit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("audits").
		Name(audit.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(audit).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the audit and deletes it. Returns an error if one occurs.
func (c *audits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("audits").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *audits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("audits").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched audit.
func (c *audits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Audit, err error) {
	result = &v1alpha1.Audit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("audits").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
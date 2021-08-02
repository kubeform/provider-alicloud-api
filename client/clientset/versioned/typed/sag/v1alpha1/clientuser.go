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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/sag/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClientUsersGetter has a method to return a ClientUserInterface.
// A group's client should implement this interface.
type ClientUsersGetter interface {
	ClientUsers(namespace string) ClientUserInterface
}

// ClientUserInterface has methods to work with ClientUser resources.
type ClientUserInterface interface {
	Create(ctx context.Context, clientUser *v1alpha1.ClientUser, opts v1.CreateOptions) (*v1alpha1.ClientUser, error)
	Update(ctx context.Context, clientUser *v1alpha1.ClientUser, opts v1.UpdateOptions) (*v1alpha1.ClientUser, error)
	UpdateStatus(ctx context.Context, clientUser *v1alpha1.ClientUser, opts v1.UpdateOptions) (*v1alpha1.ClientUser, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClientUser, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClientUserList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClientUser, err error)
	ClientUserExpansion
}

// clientUsers implements ClientUserInterface
type clientUsers struct {
	client rest.Interface
	ns     string
}

// newClientUsers returns a ClientUsers
func newClientUsers(c *SagV1alpha1Client, namespace string) *clientUsers {
	return &clientUsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clientUser, and returns the corresponding clientUser object, and an error if there is any.
func (c *clientUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClientUser, err error) {
	result = &v1alpha1.ClientUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clientusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClientUsers that match those selectors.
func (c *clientUsers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClientUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClientUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clientusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clientUsers.
func (c *clientUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clientusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clientUser and creates it.  Returns the server's representation of the clientUser, and an error, if there is any.
func (c *clientUsers) Create(ctx context.Context, clientUser *v1alpha1.ClientUser, opts v1.CreateOptions) (result *v1alpha1.ClientUser, err error) {
	result = &v1alpha1.ClientUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clientusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clientUser).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clientUser and updates it. Returns the server's representation of the clientUser, and an error, if there is any.
func (c *clientUsers) Update(ctx context.Context, clientUser *v1alpha1.ClientUser, opts v1.UpdateOptions) (result *v1alpha1.ClientUser, err error) {
	result = &v1alpha1.ClientUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clientusers").
		Name(clientUser.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clientUser).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clientUsers) UpdateStatus(ctx context.Context, clientUser *v1alpha1.ClientUser, opts v1.UpdateOptions) (result *v1alpha1.ClientUser, err error) {
	result = &v1alpha1.ClientUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clientusers").
		Name(clientUser.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clientUser).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clientUser and deletes it. Returns an error if one occurs.
func (c *clientUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clientusers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clientUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clientusers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clientUser.
func (c *clientUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClientUser, err error) {
	result = &v1alpha1.ClientUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clientusers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

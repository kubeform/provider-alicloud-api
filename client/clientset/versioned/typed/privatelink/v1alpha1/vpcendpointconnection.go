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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/privatelink/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VpcEndpointConnectionsGetter has a method to return a VpcEndpointConnectionInterface.
// A group's client should implement this interface.
type VpcEndpointConnectionsGetter interface {
	VpcEndpointConnections(namespace string) VpcEndpointConnectionInterface
}

// VpcEndpointConnectionInterface has methods to work with VpcEndpointConnection resources.
type VpcEndpointConnectionInterface interface {
	Create(ctx context.Context, vpcEndpointConnection *v1alpha1.VpcEndpointConnection, opts v1.CreateOptions) (*v1alpha1.VpcEndpointConnection, error)
	Update(ctx context.Context, vpcEndpointConnection *v1alpha1.VpcEndpointConnection, opts v1.UpdateOptions) (*v1alpha1.VpcEndpointConnection, error)
	UpdateStatus(ctx context.Context, vpcEndpointConnection *v1alpha1.VpcEndpointConnection, opts v1.UpdateOptions) (*v1alpha1.VpcEndpointConnection, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VpcEndpointConnection, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VpcEndpointConnectionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcEndpointConnection, err error)
	VpcEndpointConnectionExpansion
}

// vpcEndpointConnections implements VpcEndpointConnectionInterface
type vpcEndpointConnections struct {
	client rest.Interface
	ns     string
}

// newVpcEndpointConnections returns a VpcEndpointConnections
func newVpcEndpointConnections(c *PrivatelinkV1alpha1Client, namespace string) *vpcEndpointConnections {
	return &vpcEndpointConnections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpcEndpointConnection, and returns the corresponding vpcEndpointConnection object, and an error if there is any.
func (c *vpcEndpointConnections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpcEndpointConnection, err error) {
	result = &v1alpha1.VpcEndpointConnection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpcEndpointConnections that match those selectors.
func (c *vpcEndpointConnections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcEndpointConnectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpcEndpointConnectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpcEndpointConnections.
func (c *vpcEndpointConnections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vpcEndpointConnection and creates it.  Returns the server's representation of the vpcEndpointConnection, and an error, if there is any.
func (c *vpcEndpointConnections) Create(ctx context.Context, vpcEndpointConnection *v1alpha1.VpcEndpointConnection, opts v1.CreateOptions) (result *v1alpha1.VpcEndpointConnection, err error) {
	result = &v1alpha1.VpcEndpointConnection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpcEndpointConnection).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vpcEndpointConnection and updates it. Returns the server's representation of the vpcEndpointConnection, and an error, if there is any.
func (c *vpcEndpointConnections) Update(ctx context.Context, vpcEndpointConnection *v1alpha1.VpcEndpointConnection, opts v1.UpdateOptions) (result *v1alpha1.VpcEndpointConnection, err error) {
	result = &v1alpha1.VpcEndpointConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		Name(vpcEndpointConnection.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpcEndpointConnection).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vpcEndpointConnections) UpdateStatus(ctx context.Context, vpcEndpointConnection *v1alpha1.VpcEndpointConnection, opts v1.UpdateOptions) (result *v1alpha1.VpcEndpointConnection, err error) {
	result = &v1alpha1.VpcEndpointConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		Name(vpcEndpointConnection.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpcEndpointConnection).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vpcEndpointConnection and deletes it. Returns an error if one occurs.
func (c *vpcEndpointConnections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpcEndpointConnections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vpcEndpointConnection.
func (c *vpcEndpointConnections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcEndpointConnection, err error) {
	result = &v1alpha1.VpcEndpointConnection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpcendpointconnections").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
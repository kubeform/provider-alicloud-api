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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagerResourceGroupsGetter has a method to return a ManagerResourceGroupInterface.
// A group's client should implement this interface.
type ManagerResourceGroupsGetter interface {
	ManagerResourceGroups(namespace string) ManagerResourceGroupInterface
}

// ManagerResourceGroupInterface has methods to work with ManagerResourceGroup resources.
type ManagerResourceGroupInterface interface {
	Create(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.CreateOptions) (*v1alpha1.ManagerResourceGroup, error)
	Update(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.UpdateOptions) (*v1alpha1.ManagerResourceGroup, error)
	UpdateStatus(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.UpdateOptions) (*v1alpha1.ManagerResourceGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ManagerResourceGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ManagerResourceGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerResourceGroup, err error)
	ManagerResourceGroupExpansion
}

// managerResourceGroups implements ManagerResourceGroupInterface
type managerResourceGroups struct {
	client rest.Interface
	ns     string
}

// newManagerResourceGroups returns a ManagerResourceGroups
func newManagerResourceGroups(c *ResourceV1alpha1Client, namespace string) *managerResourceGroups {
	return &managerResourceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managerResourceGroup, and returns the corresponding managerResourceGroup object, and an error if there is any.
func (c *managerResourceGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	result = &v1alpha1.ManagerResourceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagerResourceGroups that match those selectors.
func (c *managerResourceGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerResourceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ManagerResourceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managerResourceGroups.
func (c *managerResourceGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managerResourceGroup and creates it.  Returns the server's representation of the managerResourceGroup, and an error, if there is any.
func (c *managerResourceGroups) Create(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.CreateOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	result = &v1alpha1.ManagerResourceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerResourceGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managerResourceGroup and updates it. Returns the server's representation of the managerResourceGroup, and an error, if there is any.
func (c *managerResourceGroups) Update(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.UpdateOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	result = &v1alpha1.ManagerResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		Name(managerResourceGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerResourceGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managerResourceGroups) UpdateStatus(ctx context.Context, managerResourceGroup *v1alpha1.ManagerResourceGroup, opts v1.UpdateOptions) (result *v1alpha1.ManagerResourceGroup, err error) {
	result = &v1alpha1.ManagerResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		Name(managerResourceGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerResourceGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managerResourceGroup and deletes it. Returns an error if one occurs.
func (c *managerResourceGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managerResourceGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managerresourcegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managerResourceGroup.
func (c *managerResourceGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerResourceGroup, err error) {
	result = &v1alpha1.ManagerResourceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managerresourcegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

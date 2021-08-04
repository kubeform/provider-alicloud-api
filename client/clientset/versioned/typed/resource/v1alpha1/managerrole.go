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

// ManagerRolesGetter has a method to return a ManagerRoleInterface.
// A group's client should implement this interface.
type ManagerRolesGetter interface {
	ManagerRoles(namespace string) ManagerRoleInterface
}

// ManagerRoleInterface has methods to work with ManagerRole resources.
type ManagerRoleInterface interface {
	Create(ctx context.Context, managerRole *v1alpha1.ManagerRole, opts v1.CreateOptions) (*v1alpha1.ManagerRole, error)
	Update(ctx context.Context, managerRole *v1alpha1.ManagerRole, opts v1.UpdateOptions) (*v1alpha1.ManagerRole, error)
	UpdateStatus(ctx context.Context, managerRole *v1alpha1.ManagerRole, opts v1.UpdateOptions) (*v1alpha1.ManagerRole, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ManagerRole, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ManagerRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerRole, err error)
	ManagerRoleExpansion
}

// managerRoles implements ManagerRoleInterface
type managerRoles struct {
	client rest.Interface
	ns     string
}

// newManagerRoles returns a ManagerRoles
func newManagerRoles(c *ResourceV1alpha1Client, namespace string) *managerRoles {
	return &managerRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managerRole, and returns the corresponding managerRole object, and an error if there is any.
func (c *managerRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerRole, err error) {
	result = &v1alpha1.ManagerRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managerroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagerRoles that match those selectors.
func (c *managerRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ManagerRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managerroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managerRoles.
func (c *managerRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managerroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managerRole and creates it.  Returns the server's representation of the managerRole, and an error, if there is any.
func (c *managerRoles) Create(ctx context.Context, managerRole *v1alpha1.ManagerRole, opts v1.CreateOptions) (result *v1alpha1.ManagerRole, err error) {
	result = &v1alpha1.ManagerRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managerroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerRole).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managerRole and updates it. Returns the server's representation of the managerRole, and an error, if there is any.
func (c *managerRoles) Update(ctx context.Context, managerRole *v1alpha1.ManagerRole, opts v1.UpdateOptions) (result *v1alpha1.ManagerRole, err error) {
	result = &v1alpha1.ManagerRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managerroles").
		Name(managerRole.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerRole).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managerRoles) UpdateStatus(ctx context.Context, managerRole *v1alpha1.ManagerRole, opts v1.UpdateOptions) (result *v1alpha1.ManagerRole, err error) {
	result = &v1alpha1.ManagerRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managerroles").
		Name(managerRole.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerRole).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managerRole and deletes it. Returns an error if one occurs.
func (c *managerRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managerroles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managerRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managerroles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managerRole.
func (c *managerRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerRole, err error) {
	result = &v1alpha1.ManagerRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managerroles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
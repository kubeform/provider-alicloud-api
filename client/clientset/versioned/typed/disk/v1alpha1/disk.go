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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/disk/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DisksGetter has a method to return a DiskInterface.
// A group's client should implement this interface.
type DisksGetter interface {
	Disks(namespace string) DiskInterface
}

// DiskInterface has methods to work with Disk resources.
type DiskInterface interface {
	Create(ctx context.Context, disk *v1alpha1.Disk, opts v1.CreateOptions) (*v1alpha1.Disk, error)
	Update(ctx context.Context, disk *v1alpha1.Disk, opts v1.UpdateOptions) (*v1alpha1.Disk, error)
	UpdateStatus(ctx context.Context, disk *v1alpha1.Disk, opts v1.UpdateOptions) (*v1alpha1.Disk, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Disk, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DiskList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Disk, err error)
	DiskExpansion
}

// disks implements DiskInterface
type disks struct {
	client rest.Interface
	ns     string
}

// newDisks returns a Disks
func newDisks(c *DiskV1alpha1Client, namespace string) *disks {
	return &disks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the disk, and returns the corresponding disk object, and an error if there is any.
func (c *disks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Disks that match those selectors.
func (c *disks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DiskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DiskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested disks.
func (c *disks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("disks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a disk and creates it.  Returns the server's representation of the disk, and an error, if there is any.
func (c *disks) Create(ctx context.Context, disk *v1alpha1.Disk, opts v1.CreateOptions) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("disks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disk).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a disk and updates it. Returns the server's representation of the disk, and an error, if there is any.
func (c *disks) Update(ctx context.Context, disk *v1alpha1.Disk, opts v1.UpdateOptions) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("disks").
		Name(disk.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disk).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *disks) UpdateStatus(ctx context.Context, disk *v1alpha1.Disk, opts v1.UpdateOptions) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("disks").
		Name(disk.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disk).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the disk and deletes it. Returns an error if one occurs.
func (c *disks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("disks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *disks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("disks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched disk.
func (c *disks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("disks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
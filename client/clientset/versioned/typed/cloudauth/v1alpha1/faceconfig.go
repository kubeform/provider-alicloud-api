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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloudauth/v1alpha1"
	scheme "kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FaceConfigsGetter has a method to return a FaceConfigInterface.
// A group's client should implement this interface.
type FaceConfigsGetter interface {
	FaceConfigs(namespace string) FaceConfigInterface
}

// FaceConfigInterface has methods to work with FaceConfig resources.
type FaceConfigInterface interface {
	Create(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.CreateOptions) (*v1alpha1.FaceConfig, error)
	Update(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.UpdateOptions) (*v1alpha1.FaceConfig, error)
	UpdateStatus(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.UpdateOptions) (*v1alpha1.FaceConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FaceConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FaceConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FaceConfig, err error)
	FaceConfigExpansion
}

// faceConfigs implements FaceConfigInterface
type faceConfigs struct {
	client rest.Interface
	ns     string
}

// newFaceConfigs returns a FaceConfigs
func newFaceConfigs(c *CloudauthV1alpha1Client, namespace string) *faceConfigs {
	return &faceConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the faceConfig, and returns the corresponding faceConfig object, and an error if there is any.
func (c *faceConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FaceConfig, err error) {
	result = &v1alpha1.FaceConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("faceconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FaceConfigs that match those selectors.
func (c *faceConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FaceConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FaceConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("faceconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested faceConfigs.
func (c *faceConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("faceconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a faceConfig and creates it.  Returns the server's representation of the faceConfig, and an error, if there is any.
func (c *faceConfigs) Create(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.CreateOptions) (result *v1alpha1.FaceConfig, err error) {
	result = &v1alpha1.FaceConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("faceconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(faceConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a faceConfig and updates it. Returns the server's representation of the faceConfig, and an error, if there is any.
func (c *faceConfigs) Update(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.UpdateOptions) (result *v1alpha1.FaceConfig, err error) {
	result = &v1alpha1.FaceConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("faceconfigs").
		Name(faceConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(faceConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *faceConfigs) UpdateStatus(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.UpdateOptions) (result *v1alpha1.FaceConfig, err error) {
	result = &v1alpha1.FaceConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("faceconfigs").
		Name(faceConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(faceConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the faceConfig and deletes it. Returns an error if one occurs.
func (c *faceConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("faceconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *faceConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("faceconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched faceConfig.
func (c *faceConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FaceConfig, err error) {
	result = &v1alpha1.FaceConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("faceconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

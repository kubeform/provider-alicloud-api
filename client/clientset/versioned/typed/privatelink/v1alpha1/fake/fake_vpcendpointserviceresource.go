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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/privatelink/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpcEndpointServiceResources implements VpcEndpointServiceResourceInterface
type FakeVpcEndpointServiceResources struct {
	Fake *FakePrivatelinkV1alpha1
	ns   string
}

var vpcendpointserviceresourcesResource = schema.GroupVersionResource{Group: "privatelink.alicloud.kubeform.com", Version: "v1alpha1", Resource: "vpcendpointserviceresources"}

var vpcendpointserviceresourcesKind = schema.GroupVersionKind{Group: "privatelink.alicloud.kubeform.com", Version: "v1alpha1", Kind: "VpcEndpointServiceResource"}

// Get takes name of the vpcEndpointServiceResource, and returns the corresponding vpcEndpointServiceResource object, and an error if there is any.
func (c *FakeVpcEndpointServiceResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpcEndpointServiceResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcendpointserviceresourcesResource, c.ns, name), &v1alpha1.VpcEndpointServiceResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointServiceResource), err
}

// List takes label and field selectors, and returns the list of VpcEndpointServiceResources that match those selectors.
func (c *FakeVpcEndpointServiceResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcEndpointServiceResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcendpointserviceresourcesResource, vpcendpointserviceresourcesKind, c.ns, opts), &v1alpha1.VpcEndpointServiceResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcEndpointServiceResourceList{ListMeta: obj.(*v1alpha1.VpcEndpointServiceResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcEndpointServiceResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcEndpointServiceResources.
func (c *FakeVpcEndpointServiceResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcendpointserviceresourcesResource, c.ns, opts))

}

// Create takes the representation of a vpcEndpointServiceResource and creates it.  Returns the server's representation of the vpcEndpointServiceResource, and an error, if there is any.
func (c *FakeVpcEndpointServiceResources) Create(ctx context.Context, vpcEndpointServiceResource *v1alpha1.VpcEndpointServiceResource, opts v1.CreateOptions) (result *v1alpha1.VpcEndpointServiceResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcendpointserviceresourcesResource, c.ns, vpcEndpointServiceResource), &v1alpha1.VpcEndpointServiceResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointServiceResource), err
}

// Update takes the representation of a vpcEndpointServiceResource and updates it. Returns the server's representation of the vpcEndpointServiceResource, and an error, if there is any.
func (c *FakeVpcEndpointServiceResources) Update(ctx context.Context, vpcEndpointServiceResource *v1alpha1.VpcEndpointServiceResource, opts v1.UpdateOptions) (result *v1alpha1.VpcEndpointServiceResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcendpointserviceresourcesResource, c.ns, vpcEndpointServiceResource), &v1alpha1.VpcEndpointServiceResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointServiceResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcEndpointServiceResources) UpdateStatus(ctx context.Context, vpcEndpointServiceResource *v1alpha1.VpcEndpointServiceResource, opts v1.UpdateOptions) (*v1alpha1.VpcEndpointServiceResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcendpointserviceresourcesResource, "status", c.ns, vpcEndpointServiceResource), &v1alpha1.VpcEndpointServiceResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointServiceResource), err
}

// Delete takes name of the vpcEndpointServiceResource and deletes it. Returns an error if one occurs.
func (c *FakeVpcEndpointServiceResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcendpointserviceresourcesResource, c.ns, name), &v1alpha1.VpcEndpointServiceResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcEndpointServiceResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcendpointserviceresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcEndpointServiceResourceList{})
	return err
}

// Patch applies the patch and returns the patched vpcEndpointServiceResource.
func (c *FakeVpcEndpointServiceResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcEndpointServiceResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcendpointserviceresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcEndpointServiceResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointServiceResource), err
}

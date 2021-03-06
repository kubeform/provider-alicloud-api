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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cen/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTransitRouterRouteTablePropagations implements TransitRouterRouteTablePropagationInterface
type FakeTransitRouterRouteTablePropagations struct {
	Fake *FakeCenV1alpha1
	ns   string
}

var transitrouterroutetablepropagationsResource = schema.GroupVersionResource{Group: "cen.alicloud.kubeform.com", Version: "v1alpha1", Resource: "transitrouterroutetablepropagations"}

var transitrouterroutetablepropagationsKind = schema.GroupVersionKind{Group: "cen.alicloud.kubeform.com", Version: "v1alpha1", Kind: "TransitRouterRouteTablePropagation"}

// Get takes name of the transitRouterRouteTablePropagation, and returns the corresponding transitRouterRouteTablePropagation object, and an error if there is any.
func (c *FakeTransitRouterRouteTablePropagations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitRouterRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitrouterroutetablepropagationsResource, c.ns, name), &v1alpha1.TransitRouterRouteTablePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterRouteTablePropagation), err
}

// List takes label and field selectors, and returns the list of TransitRouterRouteTablePropagations that match those selectors.
func (c *FakeTransitRouterRouteTablePropagations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitRouterRouteTablePropagationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitrouterroutetablepropagationsResource, transitrouterroutetablepropagationsKind, c.ns, opts), &v1alpha1.TransitRouterRouteTablePropagationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitRouterRouteTablePropagationList{ListMeta: obj.(*v1alpha1.TransitRouterRouteTablePropagationList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitRouterRouteTablePropagationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitRouterRouteTablePropagations.
func (c *FakeTransitRouterRouteTablePropagations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitrouterroutetablepropagationsResource, c.ns, opts))

}

// Create takes the representation of a transitRouterRouteTablePropagation and creates it.  Returns the server's representation of the transitRouterRouteTablePropagation, and an error, if there is any.
func (c *FakeTransitRouterRouteTablePropagations) Create(ctx context.Context, transitRouterRouteTablePropagation *v1alpha1.TransitRouterRouteTablePropagation, opts v1.CreateOptions) (result *v1alpha1.TransitRouterRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitrouterroutetablepropagationsResource, c.ns, transitRouterRouteTablePropagation), &v1alpha1.TransitRouterRouteTablePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterRouteTablePropagation), err
}

// Update takes the representation of a transitRouterRouteTablePropagation and updates it. Returns the server's representation of the transitRouterRouteTablePropagation, and an error, if there is any.
func (c *FakeTransitRouterRouteTablePropagations) Update(ctx context.Context, transitRouterRouteTablePropagation *v1alpha1.TransitRouterRouteTablePropagation, opts v1.UpdateOptions) (result *v1alpha1.TransitRouterRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitrouterroutetablepropagationsResource, c.ns, transitRouterRouteTablePropagation), &v1alpha1.TransitRouterRouteTablePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterRouteTablePropagation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitRouterRouteTablePropagations) UpdateStatus(ctx context.Context, transitRouterRouteTablePropagation *v1alpha1.TransitRouterRouteTablePropagation, opts v1.UpdateOptions) (*v1alpha1.TransitRouterRouteTablePropagation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitrouterroutetablepropagationsResource, "status", c.ns, transitRouterRouteTablePropagation), &v1alpha1.TransitRouterRouteTablePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterRouteTablePropagation), err
}

// Delete takes name of the transitRouterRouteTablePropagation and deletes it. Returns an error if one occurs.
func (c *FakeTransitRouterRouteTablePropagations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitrouterroutetablepropagationsResource, c.ns, name), &v1alpha1.TransitRouterRouteTablePropagation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitRouterRouteTablePropagations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitrouterroutetablepropagationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitRouterRouteTablePropagationList{})
	return err
}

// Patch applies the patch and returns the patched transitRouterRouteTablePropagation.
func (c *FakeTransitRouterRouteTablePropagations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitRouterRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitrouterroutetablepropagationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitRouterRouteTablePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitRouterRouteTablePropagation), err
}

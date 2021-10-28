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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cdn/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRealTimeLogDeliveries implements RealTimeLogDeliveryInterface
type FakeRealTimeLogDeliveries struct {
	Fake *FakeCdnV1alpha1
	ns   string
}

var realtimelogdeliveriesResource = schema.GroupVersionResource{Group: "cdn.alicloud.kubeform.com", Version: "v1alpha1", Resource: "realtimelogdeliveries"}

var realtimelogdeliveriesKind = schema.GroupVersionKind{Group: "cdn.alicloud.kubeform.com", Version: "v1alpha1", Kind: "RealTimeLogDelivery"}

// Get takes name of the realTimeLogDelivery, and returns the corresponding realTimeLogDelivery object, and an error if there is any.
func (c *FakeRealTimeLogDeliveries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RealTimeLogDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(realtimelogdeliveriesResource, c.ns, name), &v1alpha1.RealTimeLogDelivery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealTimeLogDelivery), err
}

// List takes label and field selectors, and returns the list of RealTimeLogDeliveries that match those selectors.
func (c *FakeRealTimeLogDeliveries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RealTimeLogDeliveryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(realtimelogdeliveriesResource, realtimelogdeliveriesKind, c.ns, opts), &v1alpha1.RealTimeLogDeliveryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RealTimeLogDeliveryList{ListMeta: obj.(*v1alpha1.RealTimeLogDeliveryList).ListMeta}
	for _, item := range obj.(*v1alpha1.RealTimeLogDeliveryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested realTimeLogDeliveries.
func (c *FakeRealTimeLogDeliveries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(realtimelogdeliveriesResource, c.ns, opts))

}

// Create takes the representation of a realTimeLogDelivery and creates it.  Returns the server's representation of the realTimeLogDelivery, and an error, if there is any.
func (c *FakeRealTimeLogDeliveries) Create(ctx context.Context, realTimeLogDelivery *v1alpha1.RealTimeLogDelivery, opts v1.CreateOptions) (result *v1alpha1.RealTimeLogDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(realtimelogdeliveriesResource, c.ns, realTimeLogDelivery), &v1alpha1.RealTimeLogDelivery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealTimeLogDelivery), err
}

// Update takes the representation of a realTimeLogDelivery and updates it. Returns the server's representation of the realTimeLogDelivery, and an error, if there is any.
func (c *FakeRealTimeLogDeliveries) Update(ctx context.Context, realTimeLogDelivery *v1alpha1.RealTimeLogDelivery, opts v1.UpdateOptions) (result *v1alpha1.RealTimeLogDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(realtimelogdeliveriesResource, c.ns, realTimeLogDelivery), &v1alpha1.RealTimeLogDelivery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealTimeLogDelivery), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRealTimeLogDeliveries) UpdateStatus(ctx context.Context, realTimeLogDelivery *v1alpha1.RealTimeLogDelivery, opts v1.UpdateOptions) (*v1alpha1.RealTimeLogDelivery, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(realtimelogdeliveriesResource, "status", c.ns, realTimeLogDelivery), &v1alpha1.RealTimeLogDelivery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealTimeLogDelivery), err
}

// Delete takes name of the realTimeLogDelivery and deletes it. Returns an error if one occurs.
func (c *FakeRealTimeLogDeliveries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(realtimelogdeliveriesResource, c.ns, name), &v1alpha1.RealTimeLogDelivery{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRealTimeLogDeliveries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(realtimelogdeliveriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RealTimeLogDeliveryList{})
	return err
}

// Patch applies the patch and returns the patched realTimeLogDelivery.
func (c *FakeRealTimeLogDeliveries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RealTimeLogDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(realtimelogdeliveriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RealTimeLogDelivery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealTimeLogDelivery), err
}

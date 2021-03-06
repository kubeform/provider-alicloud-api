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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/config/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAggregators implements AggregatorInterface
type FakeAggregators struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var aggregatorsResource = schema.GroupVersionResource{Group: "config.alicloud.kubeform.com", Version: "v1alpha1", Resource: "aggregators"}

var aggregatorsKind = schema.GroupVersionKind{Group: "config.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Aggregator"}

// Get takes name of the aggregator, and returns the corresponding aggregator object, and an error if there is any.
func (c *FakeAggregators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Aggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(aggregatorsResource, c.ns, name), &v1alpha1.Aggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Aggregator), err
}

// List takes label and field selectors, and returns the list of Aggregators that match those selectors.
func (c *FakeAggregators) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AggregatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(aggregatorsResource, aggregatorsKind, c.ns, opts), &v1alpha1.AggregatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AggregatorList{ListMeta: obj.(*v1alpha1.AggregatorList).ListMeta}
	for _, item := range obj.(*v1alpha1.AggregatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aggregators.
func (c *FakeAggregators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(aggregatorsResource, c.ns, opts))

}

// Create takes the representation of a aggregator and creates it.  Returns the server's representation of the aggregator, and an error, if there is any.
func (c *FakeAggregators) Create(ctx context.Context, aggregator *v1alpha1.Aggregator, opts v1.CreateOptions) (result *v1alpha1.Aggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(aggregatorsResource, c.ns, aggregator), &v1alpha1.Aggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Aggregator), err
}

// Update takes the representation of a aggregator and updates it. Returns the server's representation of the aggregator, and an error, if there is any.
func (c *FakeAggregators) Update(ctx context.Context, aggregator *v1alpha1.Aggregator, opts v1.UpdateOptions) (result *v1alpha1.Aggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(aggregatorsResource, c.ns, aggregator), &v1alpha1.Aggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Aggregator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAggregators) UpdateStatus(ctx context.Context, aggregator *v1alpha1.Aggregator, opts v1.UpdateOptions) (*v1alpha1.Aggregator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(aggregatorsResource, "status", c.ns, aggregator), &v1alpha1.Aggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Aggregator), err
}

// Delete takes name of the aggregator and deletes it. Returns an error if one occurs.
func (c *FakeAggregators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(aggregatorsResource, c.ns, name), &v1alpha1.Aggregator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAggregators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(aggregatorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AggregatorList{})
	return err
}

// Patch applies the patch and returns the patched aggregator.
func (c *FakeAggregators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Aggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(aggregatorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Aggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Aggregator), err
}

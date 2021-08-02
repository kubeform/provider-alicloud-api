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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/amqp/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExchanges implements ExchangeInterface
type FakeExchanges struct {
	Fake *FakeAmqpV1alpha1
	ns   string
}

var exchangesResource = schema.GroupVersionResource{Group: "amqp.alicloud.kubeform.com", Version: "v1alpha1", Resource: "exchanges"}

var exchangesKind = schema.GroupVersionKind{Group: "amqp.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Exchange"}

// Get takes name of the exchange, and returns the corresponding exchange object, and an error if there is any.
func (c *FakeExchanges) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Exchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(exchangesResource, c.ns, name), &v1alpha1.Exchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Exchange), err
}

// List takes label and field selectors, and returns the list of Exchanges that match those selectors.
func (c *FakeExchanges) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ExchangeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(exchangesResource, exchangesKind, c.ns, opts), &v1alpha1.ExchangeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ExchangeList{ListMeta: obj.(*v1alpha1.ExchangeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ExchangeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested exchanges.
func (c *FakeExchanges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(exchangesResource, c.ns, opts))

}

// Create takes the representation of a exchange and creates it.  Returns the server's representation of the exchange, and an error, if there is any.
func (c *FakeExchanges) Create(ctx context.Context, exchange *v1alpha1.Exchange, opts v1.CreateOptions) (result *v1alpha1.Exchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(exchangesResource, c.ns, exchange), &v1alpha1.Exchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Exchange), err
}

// Update takes the representation of a exchange and updates it. Returns the server's representation of the exchange, and an error, if there is any.
func (c *FakeExchanges) Update(ctx context.Context, exchange *v1alpha1.Exchange, opts v1.UpdateOptions) (result *v1alpha1.Exchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(exchangesResource, c.ns, exchange), &v1alpha1.Exchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Exchange), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExchanges) UpdateStatus(ctx context.Context, exchange *v1alpha1.Exchange, opts v1.UpdateOptions) (*v1alpha1.Exchange, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(exchangesResource, "status", c.ns, exchange), &v1alpha1.Exchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Exchange), err
}

// Delete takes name of the exchange and deletes it. Returns an error if one occurs.
func (c *FakeExchanges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(exchangesResource, c.ns, name), &v1alpha1.Exchange{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExchanges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(exchangesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ExchangeList{})
	return err
}

// Patch applies the patch and returns the patched exchange.
func (c *FakeExchanges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Exchange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(exchangesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Exchange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Exchange), err
}

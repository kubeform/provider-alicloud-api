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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vpc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIpv6InternetBandwidths implements Ipv6InternetBandwidthInterface
type FakeIpv6InternetBandwidths struct {
	Fake *FakeVpcV1alpha1
	ns   string
}

var ipv6internetbandwidthsResource = schema.GroupVersionResource{Group: "vpc.alicloud.kubeform.com", Version: "v1alpha1", Resource: "ipv6internetbandwidths"}

var ipv6internetbandwidthsKind = schema.GroupVersionKind{Group: "vpc.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Ipv6InternetBandwidth"}

// Get takes name of the ipv6InternetBandwidth, and returns the corresponding ipv6InternetBandwidth object, and an error if there is any.
func (c *FakeIpv6InternetBandwidths) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ipv6InternetBandwidth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipv6internetbandwidthsResource, c.ns, name), &v1alpha1.Ipv6InternetBandwidth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6InternetBandwidth), err
}

// List takes label and field selectors, and returns the list of Ipv6InternetBandwidths that match those selectors.
func (c *FakeIpv6InternetBandwidths) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ipv6InternetBandwidthList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipv6internetbandwidthsResource, ipv6internetbandwidthsKind, c.ns, opts), &v1alpha1.Ipv6InternetBandwidthList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ipv6InternetBandwidthList{ListMeta: obj.(*v1alpha1.Ipv6InternetBandwidthList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ipv6InternetBandwidthList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipv6InternetBandwidths.
func (c *FakeIpv6InternetBandwidths) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipv6internetbandwidthsResource, c.ns, opts))

}

// Create takes the representation of a ipv6InternetBandwidth and creates it.  Returns the server's representation of the ipv6InternetBandwidth, and an error, if there is any.
func (c *FakeIpv6InternetBandwidths) Create(ctx context.Context, ipv6InternetBandwidth *v1alpha1.Ipv6InternetBandwidth, opts v1.CreateOptions) (result *v1alpha1.Ipv6InternetBandwidth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipv6internetbandwidthsResource, c.ns, ipv6InternetBandwidth), &v1alpha1.Ipv6InternetBandwidth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6InternetBandwidth), err
}

// Update takes the representation of a ipv6InternetBandwidth and updates it. Returns the server's representation of the ipv6InternetBandwidth, and an error, if there is any.
func (c *FakeIpv6InternetBandwidths) Update(ctx context.Context, ipv6InternetBandwidth *v1alpha1.Ipv6InternetBandwidth, opts v1.UpdateOptions) (result *v1alpha1.Ipv6InternetBandwidth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipv6internetbandwidthsResource, c.ns, ipv6InternetBandwidth), &v1alpha1.Ipv6InternetBandwidth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6InternetBandwidth), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpv6InternetBandwidths) UpdateStatus(ctx context.Context, ipv6InternetBandwidth *v1alpha1.Ipv6InternetBandwidth, opts v1.UpdateOptions) (*v1alpha1.Ipv6InternetBandwidth, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipv6internetbandwidthsResource, "status", c.ns, ipv6InternetBandwidth), &v1alpha1.Ipv6InternetBandwidth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6InternetBandwidth), err
}

// Delete takes name of the ipv6InternetBandwidth and deletes it. Returns an error if one occurs.
func (c *FakeIpv6InternetBandwidths) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipv6internetbandwidthsResource, c.ns, name), &v1alpha1.Ipv6InternetBandwidth{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpv6InternetBandwidths) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipv6internetbandwidthsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ipv6InternetBandwidthList{})
	return err
}

// Patch applies the patch and returns the patched ipv6InternetBandwidth.
func (c *FakeIpv6InternetBandwidths) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipv6InternetBandwidth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipv6internetbandwidthsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ipv6InternetBandwidth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6InternetBandwidth), err
}

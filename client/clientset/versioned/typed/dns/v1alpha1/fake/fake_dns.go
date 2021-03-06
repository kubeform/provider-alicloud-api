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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/dns/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDnses implements DnsInterface
type FakeDnses struct {
	Fake *FakeDnsV1alpha1
	ns   string
}

var dnsesResource = schema.GroupVersionResource{Group: "dns.alicloud.kubeform.com", Version: "v1alpha1", Resource: "dnses"}

var dnsesKind = schema.GroupVersionKind{Group: "dns.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Dns"}

// Get takes name of the dns, and returns the corresponding dns object, and an error if there is any.
func (c *FakeDnses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Dns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsesResource, c.ns, name), &v1alpha1.Dns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dns), err
}

// List takes label and field selectors, and returns the list of Dnses that match those selectors.
func (c *FakeDnses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DnsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsesResource, dnsesKind, c.ns, opts), &v1alpha1.DnsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsList{ListMeta: obj.(*v1alpha1.DnsList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnses.
func (c *FakeDnses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsesResource, c.ns, opts))

}

// Create takes the representation of a dns and creates it.  Returns the server's representation of the dns, and an error, if there is any.
func (c *FakeDnses) Create(ctx context.Context, dns *v1alpha1.Dns, opts v1.CreateOptions) (result *v1alpha1.Dns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsesResource, c.ns, dns), &v1alpha1.Dns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dns), err
}

// Update takes the representation of a dns and updates it. Returns the server's representation of the dns, and an error, if there is any.
func (c *FakeDnses) Update(ctx context.Context, dns *v1alpha1.Dns, opts v1.UpdateOptions) (result *v1alpha1.Dns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsesResource, c.ns, dns), &v1alpha1.Dns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dns), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnses) UpdateStatus(ctx context.Context, dns *v1alpha1.Dns, opts v1.UpdateOptions) (*v1alpha1.Dns, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsesResource, "status", c.ns, dns), &v1alpha1.Dns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dns), err
}

// Delete takes name of the dns and deletes it. Returns an error if one occurs.
func (c *FakeDnses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsesResource, c.ns, name), &v1alpha1.Dns{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsList{})
	return err
}

// Patch applies the patch and returns the patched dns.
func (c *FakeDnses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Dns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Dns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dns), err
}

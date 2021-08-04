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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ddoscoo/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDomainResources implements DomainResourceInterface
type FakeDomainResources struct {
	Fake *FakeDdoscooV1alpha1
	ns   string
}

var domainresourcesResource = schema.GroupVersionResource{Group: "ddoscoo.alicloud.kubeform.com", Version: "v1alpha1", Resource: "domainresources"}

var domainresourcesKind = schema.GroupVersionKind{Group: "ddoscoo.alicloud.kubeform.com", Version: "v1alpha1", Kind: "DomainResource"}

// Get takes name of the domainResource, and returns the corresponding domainResource object, and an error if there is any.
func (c *FakeDomainResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DomainResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainresourcesResource, c.ns, name), &v1alpha1.DomainResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainResource), err
}

// List takes label and field selectors, and returns the list of DomainResources that match those selectors.
func (c *FakeDomainResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DomainResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainresourcesResource, domainresourcesKind, c.ns, opts), &v1alpha1.DomainResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainResourceList{ListMeta: obj.(*v1alpha1.DomainResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domainResources.
func (c *FakeDomainResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainresourcesResource, c.ns, opts))

}

// Create takes the representation of a domainResource and creates it.  Returns the server's representation of the domainResource, and an error, if there is any.
func (c *FakeDomainResources) Create(ctx context.Context, domainResource *v1alpha1.DomainResource, opts v1.CreateOptions) (result *v1alpha1.DomainResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainresourcesResource, c.ns, domainResource), &v1alpha1.DomainResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainResource), err
}

// Update takes the representation of a domainResource and updates it. Returns the server's representation of the domainResource, and an error, if there is any.
func (c *FakeDomainResources) Update(ctx context.Context, domainResource *v1alpha1.DomainResource, opts v1.UpdateOptions) (result *v1alpha1.DomainResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainresourcesResource, c.ns, domainResource), &v1alpha1.DomainResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomainResources) UpdateStatus(ctx context.Context, domainResource *v1alpha1.DomainResource, opts v1.UpdateOptions) (*v1alpha1.DomainResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainresourcesResource, "status", c.ns, domainResource), &v1alpha1.DomainResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainResource), err
}

// Delete takes name of the domainResource and deletes it. Returns an error if one occurs.
func (c *FakeDomainResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainresourcesResource, c.ns, name), &v1alpha1.DomainResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomainResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainResourceList{})
	return err
}

// Patch applies the patch and returns the patched domainResource.
func (c *FakeDomainResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DomainResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DomainResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainResource), err
}
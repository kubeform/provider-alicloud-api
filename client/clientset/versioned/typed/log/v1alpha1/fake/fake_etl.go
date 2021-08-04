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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/log/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEtls implements EtlInterface
type FakeEtls struct {
	Fake *FakeLogV1alpha1
	ns   string
}

var etlsResource = schema.GroupVersionResource{Group: "log.alicloud.kubeform.com", Version: "v1alpha1", Resource: "etls"}

var etlsKind = schema.GroupVersionKind{Group: "log.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Etl"}

// Get takes name of the etl, and returns the corresponding etl object, and an error if there is any.
func (c *FakeEtls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Etl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etlsResource, c.ns, name), &v1alpha1.Etl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Etl), err
}

// List takes label and field selectors, and returns the list of Etls that match those selectors.
func (c *FakeEtls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EtlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etlsResource, etlsKind, c.ns, opts), &v1alpha1.EtlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EtlList{ListMeta: obj.(*v1alpha1.EtlList).ListMeta}
	for _, item := range obj.(*v1alpha1.EtlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etls.
func (c *FakeEtls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etlsResource, c.ns, opts))

}

// Create takes the representation of a etl and creates it.  Returns the server's representation of the etl, and an error, if there is any.
func (c *FakeEtls) Create(ctx context.Context, etl *v1alpha1.Etl, opts v1.CreateOptions) (result *v1alpha1.Etl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etlsResource, c.ns, etl), &v1alpha1.Etl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Etl), err
}

// Update takes the representation of a etl and updates it. Returns the server's representation of the etl, and an error, if there is any.
func (c *FakeEtls) Update(ctx context.Context, etl *v1alpha1.Etl, opts v1.UpdateOptions) (result *v1alpha1.Etl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etlsResource, c.ns, etl), &v1alpha1.Etl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Etl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEtls) UpdateStatus(ctx context.Context, etl *v1alpha1.Etl, opts v1.UpdateOptions) (*v1alpha1.Etl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etlsResource, "status", c.ns, etl), &v1alpha1.Etl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Etl), err
}

// Delete takes name of the etl and deletes it. Returns an error if one occurs.
func (c *FakeEtls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(etlsResource, c.ns, name), &v1alpha1.Etl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etlsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EtlList{})
	return err
}

// Patch applies the patch and returns the patched etl.
func (c *FakeEtls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Etl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etlsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Etl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Etl), err
}
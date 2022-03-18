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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/sddp/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataLimits implements DataLimitInterface
type FakeDataLimits struct {
	Fake *FakeSddpV1alpha1
	ns   string
}

var datalimitsResource = schema.GroupVersionResource{Group: "sddp.alicloud.kubeform.com", Version: "v1alpha1", Resource: "datalimits"}

var datalimitsKind = schema.GroupVersionKind{Group: "sddp.alicloud.kubeform.com", Version: "v1alpha1", Kind: "DataLimit"}

// Get takes name of the dataLimit, and returns the corresponding dataLimit object, and an error if there is any.
func (c *FakeDataLimits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataLimit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datalimitsResource, c.ns, name), &v1alpha1.DataLimit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLimit), err
}

// List takes label and field selectors, and returns the list of DataLimits that match those selectors.
func (c *FakeDataLimits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataLimitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datalimitsResource, datalimitsKind, c.ns, opts), &v1alpha1.DataLimitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataLimitList{ListMeta: obj.(*v1alpha1.DataLimitList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataLimitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataLimits.
func (c *FakeDataLimits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datalimitsResource, c.ns, opts))

}

// Create takes the representation of a dataLimit and creates it.  Returns the server's representation of the dataLimit, and an error, if there is any.
func (c *FakeDataLimits) Create(ctx context.Context, dataLimit *v1alpha1.DataLimit, opts v1.CreateOptions) (result *v1alpha1.DataLimit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datalimitsResource, c.ns, dataLimit), &v1alpha1.DataLimit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLimit), err
}

// Update takes the representation of a dataLimit and updates it. Returns the server's representation of the dataLimit, and an error, if there is any.
func (c *FakeDataLimits) Update(ctx context.Context, dataLimit *v1alpha1.DataLimit, opts v1.UpdateOptions) (result *v1alpha1.DataLimit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datalimitsResource, c.ns, dataLimit), &v1alpha1.DataLimit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLimit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataLimits) UpdateStatus(ctx context.Context, dataLimit *v1alpha1.DataLimit, opts v1.UpdateOptions) (*v1alpha1.DataLimit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datalimitsResource, "status", c.ns, dataLimit), &v1alpha1.DataLimit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLimit), err
}

// Delete takes name of the dataLimit and deletes it. Returns an error if one occurs.
func (c *FakeDataLimits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datalimitsResource, c.ns, name), &v1alpha1.DataLimit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataLimits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datalimitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataLimitList{})
	return err
}

// Patch applies the patch and returns the patched dataLimit.
func (c *FakeDataLimits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataLimit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datalimitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataLimit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLimit), err
}
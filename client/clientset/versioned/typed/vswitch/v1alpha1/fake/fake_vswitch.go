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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vswitch/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVswitches implements VswitchInterface
type FakeVswitches struct {
	Fake *FakeVswitchV1alpha1
	ns   string
}

var vswitchesResource = schema.GroupVersionResource{Group: "vswitch.alicloud.kubeform.com", Version: "v1alpha1", Resource: "vswitches"}

var vswitchesKind = schema.GroupVersionKind{Group: "vswitch.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Vswitch"}

// Get takes name of the vswitch, and returns the corresponding vswitch object, and an error if there is any.
func (c *FakeVswitches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vswitch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vswitchesResource, c.ns, name), &v1alpha1.Vswitch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vswitch), err
}

// List takes label and field selectors, and returns the list of Vswitches that match those selectors.
func (c *FakeVswitches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VswitchList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vswitchesResource, vswitchesKind, c.ns, opts), &v1alpha1.VswitchList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VswitchList{ListMeta: obj.(*v1alpha1.VswitchList).ListMeta}
	for _, item := range obj.(*v1alpha1.VswitchList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vswitches.
func (c *FakeVswitches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vswitchesResource, c.ns, opts))

}

// Create takes the representation of a vswitch and creates it.  Returns the server's representation of the vswitch, and an error, if there is any.
func (c *FakeVswitches) Create(ctx context.Context, vswitch *v1alpha1.Vswitch, opts v1.CreateOptions) (result *v1alpha1.Vswitch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vswitchesResource, c.ns, vswitch), &v1alpha1.Vswitch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vswitch), err
}

// Update takes the representation of a vswitch and updates it. Returns the server's representation of the vswitch, and an error, if there is any.
func (c *FakeVswitches) Update(ctx context.Context, vswitch *v1alpha1.Vswitch, opts v1.UpdateOptions) (result *v1alpha1.Vswitch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vswitchesResource, c.ns, vswitch), &v1alpha1.Vswitch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vswitch), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVswitches) UpdateStatus(ctx context.Context, vswitch *v1alpha1.Vswitch, opts v1.UpdateOptions) (*v1alpha1.Vswitch, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vswitchesResource, "status", c.ns, vswitch), &v1alpha1.Vswitch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vswitch), err
}

// Delete takes name of the vswitch and deletes it. Returns an error if one occurs.
func (c *FakeVswitches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vswitchesResource, c.ns, name), &v1alpha1.Vswitch{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVswitches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vswitchesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VswitchList{})
	return err
}

// Patch applies the patch and returns the patched vswitch.
func (c *FakeVswitches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vswitch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vswitchesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Vswitch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vswitch), err
}

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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/dts/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSynchronizationInstances implements SynchronizationInstanceInterface
type FakeSynchronizationInstances struct {
	Fake *FakeDtsV1alpha1
	ns   string
}

var synchronizationinstancesResource = schema.GroupVersionResource{Group: "dts.alicloud.kubeform.com", Version: "v1alpha1", Resource: "synchronizationinstances"}

var synchronizationinstancesKind = schema.GroupVersionKind{Group: "dts.alicloud.kubeform.com", Version: "v1alpha1", Kind: "SynchronizationInstance"}

// Get takes name of the synchronizationInstance, and returns the corresponding synchronizationInstance object, and an error if there is any.
func (c *FakeSynchronizationInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SynchronizationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(synchronizationinstancesResource, c.ns, name), &v1alpha1.SynchronizationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SynchronizationInstance), err
}

// List takes label and field selectors, and returns the list of SynchronizationInstances that match those selectors.
func (c *FakeSynchronizationInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SynchronizationInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(synchronizationinstancesResource, synchronizationinstancesKind, c.ns, opts), &v1alpha1.SynchronizationInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SynchronizationInstanceList{ListMeta: obj.(*v1alpha1.SynchronizationInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SynchronizationInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested synchronizationInstances.
func (c *FakeSynchronizationInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(synchronizationinstancesResource, c.ns, opts))

}

// Create takes the representation of a synchronizationInstance and creates it.  Returns the server's representation of the synchronizationInstance, and an error, if there is any.
func (c *FakeSynchronizationInstances) Create(ctx context.Context, synchronizationInstance *v1alpha1.SynchronizationInstance, opts v1.CreateOptions) (result *v1alpha1.SynchronizationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(synchronizationinstancesResource, c.ns, synchronizationInstance), &v1alpha1.SynchronizationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SynchronizationInstance), err
}

// Update takes the representation of a synchronizationInstance and updates it. Returns the server's representation of the synchronizationInstance, and an error, if there is any.
func (c *FakeSynchronizationInstances) Update(ctx context.Context, synchronizationInstance *v1alpha1.SynchronizationInstance, opts v1.UpdateOptions) (result *v1alpha1.SynchronizationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(synchronizationinstancesResource, c.ns, synchronizationInstance), &v1alpha1.SynchronizationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SynchronizationInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSynchronizationInstances) UpdateStatus(ctx context.Context, synchronizationInstance *v1alpha1.SynchronizationInstance, opts v1.UpdateOptions) (*v1alpha1.SynchronizationInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(synchronizationinstancesResource, "status", c.ns, synchronizationInstance), &v1alpha1.SynchronizationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SynchronizationInstance), err
}

// Delete takes name of the synchronizationInstance and deletes it. Returns an error if one occurs.
func (c *FakeSynchronizationInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(synchronizationinstancesResource, c.ns, name), &v1alpha1.SynchronizationInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSynchronizationInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(synchronizationinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SynchronizationInstanceList{})
	return err
}

// Patch applies the patch and returns the patched synchronizationInstance.
func (c *FakeSynchronizationInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SynchronizationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(synchronizationinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SynchronizationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SynchronizationInstance), err
}

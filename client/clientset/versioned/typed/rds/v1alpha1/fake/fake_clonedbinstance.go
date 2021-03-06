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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/rds/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloneDbInstances implements CloneDbInstanceInterface
type FakeCloneDbInstances struct {
	Fake *FakeRdsV1alpha1
	ns   string
}

var clonedbinstancesResource = schema.GroupVersionResource{Group: "rds.alicloud.kubeform.com", Version: "v1alpha1", Resource: "clonedbinstances"}

var clonedbinstancesKind = schema.GroupVersionKind{Group: "rds.alicloud.kubeform.com", Version: "v1alpha1", Kind: "CloneDbInstance"}

// Get takes name of the cloneDbInstance, and returns the corresponding cloneDbInstance object, and an error if there is any.
func (c *FakeCloneDbInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloneDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clonedbinstancesResource, c.ns, name), &v1alpha1.CloneDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneDbInstance), err
}

// List takes label and field selectors, and returns the list of CloneDbInstances that match those selectors.
func (c *FakeCloneDbInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloneDbInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clonedbinstancesResource, clonedbinstancesKind, c.ns, opts), &v1alpha1.CloneDbInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloneDbInstanceList{ListMeta: obj.(*v1alpha1.CloneDbInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloneDbInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloneDbInstances.
func (c *FakeCloneDbInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clonedbinstancesResource, c.ns, opts))

}

// Create takes the representation of a cloneDbInstance and creates it.  Returns the server's representation of the cloneDbInstance, and an error, if there is any.
func (c *FakeCloneDbInstances) Create(ctx context.Context, cloneDbInstance *v1alpha1.CloneDbInstance, opts v1.CreateOptions) (result *v1alpha1.CloneDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clonedbinstancesResource, c.ns, cloneDbInstance), &v1alpha1.CloneDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneDbInstance), err
}

// Update takes the representation of a cloneDbInstance and updates it. Returns the server's representation of the cloneDbInstance, and an error, if there is any.
func (c *FakeCloneDbInstances) Update(ctx context.Context, cloneDbInstance *v1alpha1.CloneDbInstance, opts v1.UpdateOptions) (result *v1alpha1.CloneDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clonedbinstancesResource, c.ns, cloneDbInstance), &v1alpha1.CloneDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneDbInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloneDbInstances) UpdateStatus(ctx context.Context, cloneDbInstance *v1alpha1.CloneDbInstance, opts v1.UpdateOptions) (*v1alpha1.CloneDbInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clonedbinstancesResource, "status", c.ns, cloneDbInstance), &v1alpha1.CloneDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneDbInstance), err
}

// Delete takes name of the cloneDbInstance and deletes it. Returns an error if one occurs.
func (c *FakeCloneDbInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clonedbinstancesResource, c.ns, name), &v1alpha1.CloneDbInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloneDbInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clonedbinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloneDbInstanceList{})
	return err
}

// Patch applies the patch and returns the patched cloneDbInstance.
func (c *FakeCloneDbInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloneDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clonedbinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloneDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloneDbInstance), err
}

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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/db/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReadonlyInstances implements ReadonlyInstanceInterface
type FakeReadonlyInstances struct {
	Fake *FakeDbV1alpha1
	ns   string
}

var readonlyinstancesResource = schema.GroupVersionResource{Group: "db.alicloud.kubeform.com", Version: "v1alpha1", Resource: "readonlyinstances"}

var readonlyinstancesKind = schema.GroupVersionKind{Group: "db.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ReadonlyInstance"}

// Get takes name of the readonlyInstance, and returns the corresponding readonlyInstance object, and an error if there is any.
func (c *FakeReadonlyInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReadonlyInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(readonlyinstancesResource, c.ns, name), &v1alpha1.ReadonlyInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReadonlyInstance), err
}

// List takes label and field selectors, and returns the list of ReadonlyInstances that match those selectors.
func (c *FakeReadonlyInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReadonlyInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(readonlyinstancesResource, readonlyinstancesKind, c.ns, opts), &v1alpha1.ReadonlyInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReadonlyInstanceList{ListMeta: obj.(*v1alpha1.ReadonlyInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReadonlyInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested readonlyInstances.
func (c *FakeReadonlyInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(readonlyinstancesResource, c.ns, opts))

}

// Create takes the representation of a readonlyInstance and creates it.  Returns the server's representation of the readonlyInstance, and an error, if there is any.
func (c *FakeReadonlyInstances) Create(ctx context.Context, readonlyInstance *v1alpha1.ReadonlyInstance, opts v1.CreateOptions) (result *v1alpha1.ReadonlyInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(readonlyinstancesResource, c.ns, readonlyInstance), &v1alpha1.ReadonlyInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReadonlyInstance), err
}

// Update takes the representation of a readonlyInstance and updates it. Returns the server's representation of the readonlyInstance, and an error, if there is any.
func (c *FakeReadonlyInstances) Update(ctx context.Context, readonlyInstance *v1alpha1.ReadonlyInstance, opts v1.UpdateOptions) (result *v1alpha1.ReadonlyInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(readonlyinstancesResource, c.ns, readonlyInstance), &v1alpha1.ReadonlyInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReadonlyInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReadonlyInstances) UpdateStatus(ctx context.Context, readonlyInstance *v1alpha1.ReadonlyInstance, opts v1.UpdateOptions) (*v1alpha1.ReadonlyInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(readonlyinstancesResource, "status", c.ns, readonlyInstance), &v1alpha1.ReadonlyInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReadonlyInstance), err
}

// Delete takes name of the readonlyInstance and deletes it. Returns an error if one occurs.
func (c *FakeReadonlyInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(readonlyinstancesResource, c.ns, name), &v1alpha1.ReadonlyInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReadonlyInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(readonlyinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReadonlyInstanceList{})
	return err
}

// Patch applies the patch and returns the patched readonlyInstance.
func (c *FakeReadonlyInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReadonlyInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(readonlyinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReadonlyInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReadonlyInstance), err
}

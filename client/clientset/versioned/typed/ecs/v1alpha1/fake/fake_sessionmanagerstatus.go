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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSessionManagerStatuses implements SessionManagerStatusInterface
type FakeSessionManagerStatuses struct {
	Fake *FakeEcsV1alpha1
	ns   string
}

var sessionmanagerstatusesResource = schema.GroupVersionResource{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "sessionmanagerstatuses"}

var sessionmanagerstatusesKind = schema.GroupVersionKind{Group: "ecs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "SessionManagerStatus"}

// Get takes name of the sessionManagerStatus, and returns the corresponding sessionManagerStatus object, and an error if there is any.
func (c *FakeSessionManagerStatuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SessionManagerStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sessionmanagerstatusesResource, c.ns, name), &v1alpha1.SessionManagerStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionManagerStatus), err
}

// List takes label and field selectors, and returns the list of SessionManagerStatuses that match those selectors.
func (c *FakeSessionManagerStatuses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SessionManagerStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sessionmanagerstatusesResource, sessionmanagerstatusesKind, c.ns, opts), &v1alpha1.SessionManagerStatusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SessionManagerStatusList{ListMeta: obj.(*v1alpha1.SessionManagerStatusList).ListMeta}
	for _, item := range obj.(*v1alpha1.SessionManagerStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sessionManagerStatuses.
func (c *FakeSessionManagerStatuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sessionmanagerstatusesResource, c.ns, opts))

}

// Create takes the representation of a sessionManagerStatus and creates it.  Returns the server's representation of the sessionManagerStatus, and an error, if there is any.
func (c *FakeSessionManagerStatuses) Create(ctx context.Context, sessionManagerStatus *v1alpha1.SessionManagerStatus, opts v1.CreateOptions) (result *v1alpha1.SessionManagerStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sessionmanagerstatusesResource, c.ns, sessionManagerStatus), &v1alpha1.SessionManagerStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionManagerStatus), err
}

// Update takes the representation of a sessionManagerStatus and updates it. Returns the server's representation of the sessionManagerStatus, and an error, if there is any.
func (c *FakeSessionManagerStatuses) Update(ctx context.Context, sessionManagerStatus *v1alpha1.SessionManagerStatus, opts v1.UpdateOptions) (result *v1alpha1.SessionManagerStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sessionmanagerstatusesResource, c.ns, sessionManagerStatus), &v1alpha1.SessionManagerStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionManagerStatus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSessionManagerStatuses) UpdateStatus(ctx context.Context, sessionManagerStatus *v1alpha1.SessionManagerStatus, opts v1.UpdateOptions) (*v1alpha1.SessionManagerStatus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sessionmanagerstatusesResource, "status", c.ns, sessionManagerStatus), &v1alpha1.SessionManagerStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionManagerStatus), err
}

// Delete takes name of the sessionManagerStatus and deletes it. Returns an error if one occurs.
func (c *FakeSessionManagerStatuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sessionmanagerstatusesResource, c.ns, name), &v1alpha1.SessionManagerStatus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSessionManagerStatuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sessionmanagerstatusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SessionManagerStatusList{})
	return err
}

// Patch applies the patch and returns the patched sessionManagerStatus.
func (c *FakeSessionManagerStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SessionManagerStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sessionmanagerstatusesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SessionManagerStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionManagerStatus), err
}

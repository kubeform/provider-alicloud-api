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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ga/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeListeners implements ListenerInterface
type FakeListeners struct {
	Fake *FakeGaV1alpha1
	ns   string
}

var listenersResource = schema.GroupVersionResource{Group: "ga.alicloud.kubeform.com", Version: "v1alpha1", Resource: "listeners"}

var listenersKind = schema.GroupVersionKind{Group: "ga.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Listener"}

// Get takes name of the listener, and returns the corresponding listener object, and an error if there is any.
func (c *FakeListeners) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Listener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(listenersResource, c.ns, name), &v1alpha1.Listener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Listener), err
}

// List takes label and field selectors, and returns the list of Listeners that match those selectors.
func (c *FakeListeners) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(listenersResource, listenersKind, c.ns, opts), &v1alpha1.ListenerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ListenerList{ListMeta: obj.(*v1alpha1.ListenerList).ListMeta}
	for _, item := range obj.(*v1alpha1.ListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested listeners.
func (c *FakeListeners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(listenersResource, c.ns, opts))

}

// Create takes the representation of a listener and creates it.  Returns the server's representation of the listener, and an error, if there is any.
func (c *FakeListeners) Create(ctx context.Context, listener *v1alpha1.Listener, opts v1.CreateOptions) (result *v1alpha1.Listener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(listenersResource, c.ns, listener), &v1alpha1.Listener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Listener), err
}

// Update takes the representation of a listener and updates it. Returns the server's representation of the listener, and an error, if there is any.
func (c *FakeListeners) Update(ctx context.Context, listener *v1alpha1.Listener, opts v1.UpdateOptions) (result *v1alpha1.Listener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(listenersResource, c.ns, listener), &v1alpha1.Listener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Listener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeListeners) UpdateStatus(ctx context.Context, listener *v1alpha1.Listener, opts v1.UpdateOptions) (*v1alpha1.Listener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(listenersResource, "status", c.ns, listener), &v1alpha1.Listener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Listener), err
}

// Delete takes name of the listener and deletes it. Returns an error if one occurs.
func (c *FakeListeners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(listenersResource, c.ns, name), &v1alpha1.Listener{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeListeners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(listenersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ListenerList{})
	return err
}

// Patch applies the patch and returns the patched listener.
func (c *FakeListeners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Listener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(listenersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Listener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Listener), err
}

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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/waf/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProtectionModules implements ProtectionModuleInterface
type FakeProtectionModules struct {
	Fake *FakeWafV1alpha1
	ns   string
}

var protectionmodulesResource = schema.GroupVersionResource{Group: "waf.alicloud.kubeform.com", Version: "v1alpha1", Resource: "protectionmodules"}

var protectionmodulesKind = schema.GroupVersionKind{Group: "waf.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ProtectionModule"}

// Get takes name of the protectionModule, and returns the corresponding protectionModule object, and an error if there is any.
func (c *FakeProtectionModules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProtectionModule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(protectionmodulesResource, c.ns, name), &v1alpha1.ProtectionModule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionModule), err
}

// List takes label and field selectors, and returns the list of ProtectionModules that match those selectors.
func (c *FakeProtectionModules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProtectionModuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(protectionmodulesResource, protectionmodulesKind, c.ns, opts), &v1alpha1.ProtectionModuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProtectionModuleList{ListMeta: obj.(*v1alpha1.ProtectionModuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProtectionModuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested protectionModules.
func (c *FakeProtectionModules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(protectionmodulesResource, c.ns, opts))

}

// Create takes the representation of a protectionModule and creates it.  Returns the server's representation of the protectionModule, and an error, if there is any.
func (c *FakeProtectionModules) Create(ctx context.Context, protectionModule *v1alpha1.ProtectionModule, opts v1.CreateOptions) (result *v1alpha1.ProtectionModule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(protectionmodulesResource, c.ns, protectionModule), &v1alpha1.ProtectionModule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionModule), err
}

// Update takes the representation of a protectionModule and updates it. Returns the server's representation of the protectionModule, and an error, if there is any.
func (c *FakeProtectionModules) Update(ctx context.Context, protectionModule *v1alpha1.ProtectionModule, opts v1.UpdateOptions) (result *v1alpha1.ProtectionModule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(protectionmodulesResource, c.ns, protectionModule), &v1alpha1.ProtectionModule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionModule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProtectionModules) UpdateStatus(ctx context.Context, protectionModule *v1alpha1.ProtectionModule, opts v1.UpdateOptions) (*v1alpha1.ProtectionModule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(protectionmodulesResource, "status", c.ns, protectionModule), &v1alpha1.ProtectionModule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionModule), err
}

// Delete takes name of the protectionModule and deletes it. Returns an error if one occurs.
func (c *FakeProtectionModules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(protectionmodulesResource, c.ns, name), &v1alpha1.ProtectionModule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProtectionModules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(protectionmodulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProtectionModuleList{})
	return err
}

// Patch applies the patch and returns the patched protectionModule.
func (c *FakeProtectionModules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProtectionModule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(protectionmodulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProtectionModule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionModule), err
}
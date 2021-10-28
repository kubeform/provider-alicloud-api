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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorksFolders implements WorksFolderInterface
type FakeWorksFolders struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var worksfoldersResource = schema.GroupVersionResource{Group: "data.alicloud.kubeform.com", Version: "v1alpha1", Resource: "worksfolders"}

var worksfoldersKind = schema.GroupVersionKind{Group: "data.alicloud.kubeform.com", Version: "v1alpha1", Kind: "WorksFolder"}

// Get takes name of the worksFolder, and returns the corresponding worksFolder object, and an error if there is any.
func (c *FakeWorksFolders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorksFolder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(worksfoldersResource, c.ns, name), &v1alpha1.WorksFolder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorksFolder), err
}

// List takes label and field selectors, and returns the list of WorksFolders that match those selectors.
func (c *FakeWorksFolders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorksFolderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(worksfoldersResource, worksfoldersKind, c.ns, opts), &v1alpha1.WorksFolderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorksFolderList{ListMeta: obj.(*v1alpha1.WorksFolderList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorksFolderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested worksFolders.
func (c *FakeWorksFolders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(worksfoldersResource, c.ns, opts))

}

// Create takes the representation of a worksFolder and creates it.  Returns the server's representation of the worksFolder, and an error, if there is any.
func (c *FakeWorksFolders) Create(ctx context.Context, worksFolder *v1alpha1.WorksFolder, opts v1.CreateOptions) (result *v1alpha1.WorksFolder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(worksfoldersResource, c.ns, worksFolder), &v1alpha1.WorksFolder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorksFolder), err
}

// Update takes the representation of a worksFolder and updates it. Returns the server's representation of the worksFolder, and an error, if there is any.
func (c *FakeWorksFolders) Update(ctx context.Context, worksFolder *v1alpha1.WorksFolder, opts v1.UpdateOptions) (result *v1alpha1.WorksFolder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(worksfoldersResource, c.ns, worksFolder), &v1alpha1.WorksFolder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorksFolder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorksFolders) UpdateStatus(ctx context.Context, worksFolder *v1alpha1.WorksFolder, opts v1.UpdateOptions) (*v1alpha1.WorksFolder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(worksfoldersResource, "status", c.ns, worksFolder), &v1alpha1.WorksFolder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorksFolder), err
}

// Delete takes name of the worksFolder and deletes it. Returns an error if one occurs.
func (c *FakeWorksFolders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(worksfoldersResource, c.ns, name), &v1alpha1.WorksFolder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorksFolders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(worksfoldersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorksFolderList{})
	return err
}

// Patch applies the patch and returns the patched worksFolder.
func (c *FakeWorksFolders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorksFolder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(worksfoldersResource, c.ns, name, pt, data, subresources...), &v1alpha1.WorksFolder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorksFolder), err
}
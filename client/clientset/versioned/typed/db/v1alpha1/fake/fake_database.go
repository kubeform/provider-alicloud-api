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

// FakeDatabases implements DatabaseInterface
type FakeDatabases struct {
	Fake *FakeDbV1alpha1
	ns   string
}

var databasesResource = schema.GroupVersionResource{Group: "db.alicloud.kubeform.com", Version: "v1alpha1", Resource: "databases"}

var databasesKind = schema.GroupVersionKind{Group: "db.alicloud.kubeform.com", Version: "v1alpha1", Kind: "Database"}

// Get takes name of the database, and returns the corresponding database object, and an error if there is any.
func (c *FakeDatabases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Database, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(databasesResource, c.ns, name), &v1alpha1.Database{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Database), err
}

// List takes label and field selectors, and returns the list of Databases that match those selectors.
func (c *FakeDatabases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(databasesResource, databasesKind, c.ns, opts), &v1alpha1.DatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatabaseList{ListMeta: obj.(*v1alpha1.DatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested databases.
func (c *FakeDatabases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(databasesResource, c.ns, opts))

}

// Create takes the representation of a database and creates it.  Returns the server's representation of the database, and an error, if there is any.
func (c *FakeDatabases) Create(ctx context.Context, database *v1alpha1.Database, opts v1.CreateOptions) (result *v1alpha1.Database, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(databasesResource, c.ns, database), &v1alpha1.Database{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Database), err
}

// Update takes the representation of a database and updates it. Returns the server's representation of the database, and an error, if there is any.
func (c *FakeDatabases) Update(ctx context.Context, database *v1alpha1.Database, opts v1.UpdateOptions) (result *v1alpha1.Database, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(databasesResource, c.ns, database), &v1alpha1.Database{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Database), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatabases) UpdateStatus(ctx context.Context, database *v1alpha1.Database, opts v1.UpdateOptions) (*v1alpha1.Database, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(databasesResource, "status", c.ns, database), &v1alpha1.Database{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Database), err
}

// Delete takes name of the database and deletes it. Returns an error if one occurs.
func (c *FakeDatabases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(databasesResource, c.ns, name), &v1alpha1.Database{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatabases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(databasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatabaseList{})
	return err
}

// Patch applies the patch and returns the patched database.
func (c *FakeDatabases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Database, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(databasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Database{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Database), err
}

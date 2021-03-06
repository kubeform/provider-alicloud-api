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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/quotas/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQuotaApplications implements QuotaApplicationInterface
type FakeQuotaApplications struct {
	Fake *FakeQuotasV1alpha1
	ns   string
}

var quotaapplicationsResource = schema.GroupVersionResource{Group: "quotas.alicloud.kubeform.com", Version: "v1alpha1", Resource: "quotaapplications"}

var quotaapplicationsKind = schema.GroupVersionKind{Group: "quotas.alicloud.kubeform.com", Version: "v1alpha1", Kind: "QuotaApplication"}

// Get takes name of the quotaApplication, and returns the corresponding quotaApplication object, and an error if there is any.
func (c *FakeQuotaApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.QuotaApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(quotaapplicationsResource, c.ns, name), &v1alpha1.QuotaApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuotaApplication), err
}

// List takes label and field selectors, and returns the list of QuotaApplications that match those selectors.
func (c *FakeQuotaApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.QuotaApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(quotaapplicationsResource, quotaapplicationsKind, c.ns, opts), &v1alpha1.QuotaApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.QuotaApplicationList{ListMeta: obj.(*v1alpha1.QuotaApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.QuotaApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested quotaApplications.
func (c *FakeQuotaApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(quotaapplicationsResource, c.ns, opts))

}

// Create takes the representation of a quotaApplication and creates it.  Returns the server's representation of the quotaApplication, and an error, if there is any.
func (c *FakeQuotaApplications) Create(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.CreateOptions) (result *v1alpha1.QuotaApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(quotaapplicationsResource, c.ns, quotaApplication), &v1alpha1.QuotaApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuotaApplication), err
}

// Update takes the representation of a quotaApplication and updates it. Returns the server's representation of the quotaApplication, and an error, if there is any.
func (c *FakeQuotaApplications) Update(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.UpdateOptions) (result *v1alpha1.QuotaApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(quotaapplicationsResource, c.ns, quotaApplication), &v1alpha1.QuotaApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuotaApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeQuotaApplications) UpdateStatus(ctx context.Context, quotaApplication *v1alpha1.QuotaApplication, opts v1.UpdateOptions) (*v1alpha1.QuotaApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(quotaapplicationsResource, "status", c.ns, quotaApplication), &v1alpha1.QuotaApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuotaApplication), err
}

// Delete takes name of the quotaApplication and deletes it. Returns an error if one occurs.
func (c *FakeQuotaApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(quotaapplicationsResource, c.ns, name), &v1alpha1.QuotaApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQuotaApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(quotaapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.QuotaApplicationList{})
	return err
}

// Patch applies the patch and returns the patched quotaApplication.
func (c *FakeQuotaApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.QuotaApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(quotaapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.QuotaApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuotaApplication), err
}

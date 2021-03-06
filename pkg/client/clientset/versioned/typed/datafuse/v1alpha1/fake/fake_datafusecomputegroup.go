/*
Copyright 2021.

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

	v1alpha1 "datafuselabs.io/datafuse-operator/pkg/apis/datafuse/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatafuseComputeGroups implements DatafuseComputeGroupInterface
type FakeDatafuseComputeGroups struct {
	Fake *FakeDatafuseV1alpha1
	ns   string
}

var datafusecomputegroupsResource = schema.GroupVersionResource{Group: "datafuse", Version: "v1alpha1", Resource: "datafusecomputegroups"}

var datafusecomputegroupsKind = schema.GroupVersionKind{Group: "datafuse", Version: "v1alpha1", Kind: "DatafuseComputeGroup"}

// Get takes name of the datafuseComputeGroup, and returns the corresponding datafuseComputeGroup object, and an error if there is any.
func (c *FakeDatafuseComputeGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatafuseComputeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datafusecomputegroupsResource, c.ns, name), &v1alpha1.DatafuseComputeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatafuseComputeGroup), err
}

// List takes label and field selectors, and returns the list of DatafuseComputeGroups that match those selectors.
func (c *FakeDatafuseComputeGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatafuseComputeGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datafusecomputegroupsResource, datafusecomputegroupsKind, c.ns, opts), &v1alpha1.DatafuseComputeGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatafuseComputeGroupList{ListMeta: obj.(*v1alpha1.DatafuseComputeGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatafuseComputeGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datafuseComputeGroups.
func (c *FakeDatafuseComputeGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datafusecomputegroupsResource, c.ns, opts))

}

// Create takes the representation of a datafuseComputeGroup and creates it.  Returns the server's representation of the datafuseComputeGroup, and an error, if there is any.
func (c *FakeDatafuseComputeGroups) Create(ctx context.Context, datafuseComputeGroup *v1alpha1.DatafuseComputeGroup, opts v1.CreateOptions) (result *v1alpha1.DatafuseComputeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datafusecomputegroupsResource, c.ns, datafuseComputeGroup), &v1alpha1.DatafuseComputeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatafuseComputeGroup), err
}

// Update takes the representation of a datafuseComputeGroup and updates it. Returns the server's representation of the datafuseComputeGroup, and an error, if there is any.
func (c *FakeDatafuseComputeGroups) Update(ctx context.Context, datafuseComputeGroup *v1alpha1.DatafuseComputeGroup, opts v1.UpdateOptions) (result *v1alpha1.DatafuseComputeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datafusecomputegroupsResource, c.ns, datafuseComputeGroup), &v1alpha1.DatafuseComputeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatafuseComputeGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatafuseComputeGroups) UpdateStatus(ctx context.Context, datafuseComputeGroup *v1alpha1.DatafuseComputeGroup, opts v1.UpdateOptions) (*v1alpha1.DatafuseComputeGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datafusecomputegroupsResource, "status", c.ns, datafuseComputeGroup), &v1alpha1.DatafuseComputeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatafuseComputeGroup), err
}

// Delete takes name of the datafuseComputeGroup and deletes it. Returns an error if one occurs.
func (c *FakeDatafuseComputeGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datafusecomputegroupsResource, c.ns, name), &v1alpha1.DatafuseComputeGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatafuseComputeGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datafusecomputegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatafuseComputeGroupList{})
	return err
}

// Patch applies the patch and returns the patched datafuseComputeGroup.
func (c *FakeDatafuseComputeGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatafuseComputeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datafusecomputegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatafuseComputeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatafuseComputeGroup), err
}

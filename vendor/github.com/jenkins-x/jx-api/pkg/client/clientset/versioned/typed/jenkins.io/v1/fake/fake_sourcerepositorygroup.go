// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	jenkinsiov1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSourceRepositoryGroups implements SourceRepositoryGroupInterface
type FakeSourceRepositoryGroups struct {
	Fake *FakeJenkinsV1
	ns   string
}

var sourcerepositorygroupsResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "sourcerepositorygroups"}

var sourcerepositorygroupsKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "SourceRepositoryGroup"}

// Get takes name of the sourceRepositoryGroup, and returns the corresponding sourceRepositoryGroup object, and an error if there is any.
func (c *FakeSourceRepositoryGroups) Get(name string, options v1.GetOptions) (result *jenkinsiov1.SourceRepositoryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sourcerepositorygroupsResource, c.ns, name), &jenkinsiov1.SourceRepositoryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.SourceRepositoryGroup), err
}

// List takes label and field selectors, and returns the list of SourceRepositoryGroups that match those selectors.
func (c *FakeSourceRepositoryGroups) List(opts v1.ListOptions) (result *jenkinsiov1.SourceRepositoryGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sourcerepositorygroupsResource, sourcerepositorygroupsKind, c.ns, opts), &jenkinsiov1.SourceRepositoryGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.SourceRepositoryGroupList{ListMeta: obj.(*jenkinsiov1.SourceRepositoryGroupList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.SourceRepositoryGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sourceRepositoryGroups.
func (c *FakeSourceRepositoryGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sourcerepositorygroupsResource, c.ns, opts))

}

// Create takes the representation of a sourceRepositoryGroup and creates it.  Returns the server's representation of the sourceRepositoryGroup, and an error, if there is any.
func (c *FakeSourceRepositoryGroups) Create(sourceRepositoryGroup *jenkinsiov1.SourceRepositoryGroup) (result *jenkinsiov1.SourceRepositoryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sourcerepositorygroupsResource, c.ns, sourceRepositoryGroup), &jenkinsiov1.SourceRepositoryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.SourceRepositoryGroup), err
}

// Update takes the representation of a sourceRepositoryGroup and updates it. Returns the server's representation of the sourceRepositoryGroup, and an error, if there is any.
func (c *FakeSourceRepositoryGroups) Update(sourceRepositoryGroup *jenkinsiov1.SourceRepositoryGroup) (result *jenkinsiov1.SourceRepositoryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sourcerepositorygroupsResource, c.ns, sourceRepositoryGroup), &jenkinsiov1.SourceRepositoryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.SourceRepositoryGroup), err
}

// Delete takes name of the sourceRepositoryGroup and deletes it. Returns an error if one occurs.
func (c *FakeSourceRepositoryGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sourcerepositorygroupsResource, c.ns, name), &jenkinsiov1.SourceRepositoryGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSourceRepositoryGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sourcerepositorygroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.SourceRepositoryGroupList{})
	return err
}

// Patch applies the patch and returns the patched sourceRepositoryGroup.
func (c *FakeSourceRepositoryGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkinsiov1.SourceRepositoryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sourcerepositorygroupsResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.SourceRepositoryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.SourceRepositoryGroup), err
}
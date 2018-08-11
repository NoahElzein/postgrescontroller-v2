// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	postgrescontrollerv1 "github.com/demo/postgrescontroller/pkg/apis/postgrescontroller/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePostgreses implements PostgresInterface
type FakePostgreses struct {
	Fake *FakePostgrescontrollerV1
	ns   string
}

var postgresesResource = schema.GroupVersionResource{Group: "postgrescontroller.kubeplus", Version: "v1", Resource: "postgreses"}

var postgresesKind = schema.GroupVersionKind{Group: "postgrescontroller.kubeplus", Version: "v1", Kind: "Postgres"}

// Get takes name of the postgres, and returns the corresponding postgres object, and an error if there is any.
func (c *FakePostgreses) Get(name string, options v1.GetOptions) (result *postgrescontrollerv1.Postgres, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postgresesResource, c.ns, name), &postgrescontrollerv1.Postgres{})

	if obj == nil {
		return nil, err
	}
	return obj.(*postgrescontrollerv1.Postgres), err
}

// List takes label and field selectors, and returns the list of Postgreses that match those selectors.
func (c *FakePostgreses) List(opts v1.ListOptions) (result *postgrescontrollerv1.PostgresList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postgresesResource, postgresesKind, c.ns, opts), &postgrescontrollerv1.PostgresList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &postgrescontrollerv1.PostgresList{}
	for _, item := range obj.(*postgrescontrollerv1.PostgresList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postgreses.
func (c *FakePostgreses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postgresesResource, c.ns, opts))

}

// Create takes the representation of a postgres and creates it.  Returns the server's representation of the postgres, and an error, if there is any.
func (c *FakePostgreses) Create(postgres *postgrescontrollerv1.Postgres) (result *postgrescontrollerv1.Postgres, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postgresesResource, c.ns, postgres), &postgrescontrollerv1.Postgres{})

	if obj == nil {
		return nil, err
	}
	return obj.(*postgrescontrollerv1.Postgres), err
}

// Update takes the representation of a postgres and updates it. Returns the server's representation of the postgres, and an error, if there is any.
func (c *FakePostgreses) Update(postgres *postgrescontrollerv1.Postgres) (result *postgrescontrollerv1.Postgres, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postgresesResource, c.ns, postgres), &postgrescontrollerv1.Postgres{})

	if obj == nil {
		return nil, err
	}
	return obj.(*postgrescontrollerv1.Postgres), err
}

// Delete takes name of the postgres and deletes it. Returns an error if one occurs.
func (c *FakePostgreses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(postgresesResource, c.ns, name), &postgrescontrollerv1.Postgres{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostgreses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postgresesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &postgrescontrollerv1.PostgresList{})
	return err
}

// Patch applies the patch and returns the patched postgres.
func (c *FakePostgreses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *postgrescontrollerv1.Postgres, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postgresesResource, c.ns, name, data, subresources...), &postgrescontrollerv1.Postgres{})

	if obj == nil {
		return nil, err
	}
	return obj.(*postgrescontrollerv1.Postgres), err
}

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/demo/postgrescontroller/pkg/client/clientset/versioned/typed/postgrescontroller/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePostgrescontrollerV1 struct {
	*testing.Fake
}

func (c *FakePostgrescontrollerV1) Postgreses(namespace string) v1.PostgresInterface {
	return &FakePostgreses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePostgrescontrollerV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

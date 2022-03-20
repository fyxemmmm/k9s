package k9s

import (
	appsv1 "github.com/fyxemmmm/k9s/client-go/k9s/typed/apps/v1"
	"github.com/fyxemmmm/k9s/client-go/rest"
)

func NewForConfig(c *rest.Config) (*ClientSet, error) {
	var cs ClientSet
	var err error

	cs.appsV1, err = appsv1.NewForConfig(c)
	if err != nil {
		panic(err)
	}

	return &cs, nil
}

type ClientSet struct {
	appsV1 *appsv1.AppsV1Client
}

func (c *ClientSet) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}

package v1

import (
	v1 "github.com/fyxemmmm/k9s/client-go/k9s/api/apps/v1"
	"github.com/fyxemmmm/k9s/client-go/rest"
)

type DeploymentsGetter interface {
	Deployments() DeploymentInterface
}

type DeploymentInterface interface {
	Get() *v1.Deployment
}

type deployments struct {
	client rest.Interface
}

func newDeployments(c *AppsV1Client) *deployments {
	return &deployments{client: c.RestClient()}
}

func (c *deployments) Get() *v1.Deployment {
	return nil
}
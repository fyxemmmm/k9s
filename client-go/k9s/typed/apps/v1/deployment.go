package v1

import v1 "github.com/fyxemmmm/k9s/client-go/k9s/api/apps/v1"

type DeploymentsGetter interface {
	Deployments() DeploymentInterface
}

type DeploymentInterface interface {
	Get() *v1.Deployment
}

type deployments struct {
	
}

func newDeployments() *deployments {
	return &deployments{}
}

func (c *deployments) Get() *v1.Deployment {
	return nil
}
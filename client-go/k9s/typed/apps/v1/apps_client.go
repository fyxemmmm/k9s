package v1

import "github.com/fyxemmmm/k9s/client-go/rest"

type AppsV1Interface interface {
	RestClient() rest.Interface
	DeploymentsGetter
}

type AppsV1Client struct {
	restClient rest.Interface
}

func NewForConfig(config *rest.Config) (*AppsV1Client, error) {
	return &AppsV1Client{}, nil
}

func (c *AppsV1Client) Deployments() DeploymentInterface {
	return newDeployments(c)
}

func (c *AppsV1Client) RestClient() rest.Interface {
	return c.restClient
}

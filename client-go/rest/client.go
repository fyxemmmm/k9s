package rest

type Interface interface {
	Get() *Request
}

type RESTClient struct {
	
}

func (c *RESTClient) Get() *Request {
	return nil
}
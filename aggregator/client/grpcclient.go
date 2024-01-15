package client

type GRPCClient struct {
	Endpoint string
}

func NewGRPCClient(endpoint string) *GRPCClient {
	return &GRPCClient{
		Endpoint: endpoint,
	}
}

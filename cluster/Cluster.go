package cluster

type Endpoint struct {
	Host      string
	Port      int32
	Weight    int8
	Status    int8
	StartTime int64
}

type Invocation struct {
	MethodName string
	Args       []interface{}
	Result     interface{}
	ErrorInfo  error
}

type RpcRequest struct {
	RequestId      int64
	ServiceName    string
	InvocationInfo Invocation
}

type Cluster interface {
	getAllEndpoints() []Endpoint

	getActiveEndpoints() []Endpoint
}

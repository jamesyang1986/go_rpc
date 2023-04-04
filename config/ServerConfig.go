package config

type ServerConfig struct {
	ListenPort int32
	Services   []ServiceConfig
}

type MethodInfo struct {
	ExecuteTimeout int32
	MethodName     string
	CallArgs       []interface{}
	ReturnValue    interface{}
	ErrorInfo      error
}

type ServiceConfig struct {
	Timeout       int32
	CallInterface interface{}
	MethodList    []MethodInfo
}

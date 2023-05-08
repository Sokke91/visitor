package main

type RpcServer struct{}

type Credentials struct {
	Username string
	Password string
}

func (r *RpcServer) Login(credentials Credentials, result *string) error {
	*result = "token kdald"
	return nil
}

func (r *RpcServer) CheckToken(token string, result *bool) error {
	*result = true
	return nil
}

func (r *RpcServer) CurrentUser(token string) error {
	return nil
}

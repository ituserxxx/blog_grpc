package svc

import (
	"brl/http/internal/config"
	"brl/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"

)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

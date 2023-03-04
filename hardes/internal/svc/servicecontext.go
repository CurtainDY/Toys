package svc

import (
	"hardes/internal/config"
	"hardes/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config

	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}

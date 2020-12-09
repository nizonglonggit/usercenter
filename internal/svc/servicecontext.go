package svc

import (
	"github.com/nizonglonggit/usercenter/internal/config"
	"github.com/nizonglonggit/usercenter/model"
)

type ServiceContext struct {
	c config.Config
	Model *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
	}
}

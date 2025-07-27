package services

import "github.com/gazprom-el-monitoring/ims/internal/repositories"

type Services struct {
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{}
}

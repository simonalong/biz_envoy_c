package router

import (
	"biz-c-service/controller"
)

type Option func()

var options []Option

func include() {
	// cf相关接口
	options = append(options, controller.CfController)
}

func Register() {
	include()
	for _, opt := range options {
		opt()
	}
}

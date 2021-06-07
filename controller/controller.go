package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	DefineRoutes(router *gin.RouterGroup)
}


func RegisterController(controller Controller,router *gin.RouterGroup) {
	controller.DefineRoutes(router)
}

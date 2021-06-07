package controller

import "github.com/gin-gonic/gin"

type ResourceController interface {
	DefineRoutes(router *gin.RouterGroup)
	List(context *gin.Context)
	Find(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

func AddDefaultResourceRoutes(rc ResourceController, router *gin.RouterGroup) {
	router.GET("", rc.List)
	router.GET("/:id", rc.Find)
	router.POST("", rc.Create)
	router.PATCH("/:id", rc.Update)
	router.DELETE("/:id", rc.Delete)
}


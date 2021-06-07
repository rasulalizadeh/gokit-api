package tests

import (
	"log"
	"net/http"
	"rasulalizadeh/gokit-api/controller"
	"rasulalizadeh/gokit-api/router"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TestMongoConnect(t *testing.T) {

	router := gin.Default()

	// Disable CORS
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8090"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		MaxAge:           12 * time.Hour,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
	})
	router.Use(corsMiddleware)

	apiGroup := router.Group("api")
	apiV1 := apiGroup.Group("v1")

	controller.RegisterController(TestController{}, apiV1.Group("status"))

	// Register Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run("localhost:8090")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type TestController struct {
}

func (c TestController) DefineRoutes(router *gin.RouterGroup) {
	//store := persistence.NewRedisCache("localhost:6379", "", 10 * time.Second)
	//router.GET("/:id", cache.CachePage(store, 3 * time.Hour, c.Find))
	router.GET("", c.Status)
}

func (c TestController) Status(context *gin.Context) {
	context.JSON(http.StatusOK, router.NewSuccessResponse("Test Server Staus \"OK\"", map[string]string{
		"status":     "ok",
		"controller": "test",
	}))
}

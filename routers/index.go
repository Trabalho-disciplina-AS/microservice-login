package routers

import (
	"goseed/controllers"

	"goseed/middlewares"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.Authentication())
	// authGroup.Use(middlewares.CORSMiddleware())
	authGroup.GET("/profile", authController.Profile)

}

// InitRoute ..
func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	setAuthRoute(router)
	return router
}

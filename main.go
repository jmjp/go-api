package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"un1versum.com/clutchers/bootstrap"
	"un1versum.com/clutchers/controllers"
)

func main() {
	bootstrap.Initialize()

	router := gin.Default()
	_ = router.SetTrustedProxies(nil)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "route not found"})
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.FindUsers)
			users.POST("/", controllers.CreateUser)
			users.POST("/body", controllers.BodyCreateUser)
		}
	}

	router.Run()
}

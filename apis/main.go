package main

import (
	"apis/controllers"
	"apis/db"
	"apis/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"OK": 200,
		})
	})

	todoApis := r.Group("/todo")
	todoApis.Use(middleware.AuthorizeMiddleware())
	todoApis.GET("", controllers.GetTodoController)
	todoApis.POST("", controllers.AddTodoController)
	todoApis.PUT("/:id", controllers.UpdateTodoController)
	todoApis.DELETE("/:id", controllers.DeleteTodoController)

	r.Run(":8090")
}

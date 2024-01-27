package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vidiana0101/task-5-pbi-bptns-VIDIANA/controllers/usercontrollers"
	"github.com/vidiana0101/task-5-pbi-bptns-VIDIANA/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/user", usercontrollers.Index)
	r.GET("/api/user/:id", usercontrollers.Show)
	r.POST("/api/user", usercontrollers.Create)
	r.PUT("/api/user/:id", usercontrollers.Update)
	r.DELETE("/api/user", usercontrollers.Delete)

	r.Run()
}

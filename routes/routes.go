package routes

import (
	"api-golang-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.ShowAllStudents)
	r.Run()
}

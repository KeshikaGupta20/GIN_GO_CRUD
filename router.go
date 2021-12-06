package router

import (
	"github.com/gin-gonic/gin"

	"keshika/controller"
	
)

func RegisterRoutes(c *gin.Engine) {

	router := gin.Default()
	router.POST("/createuser", controller.CreateUser)
	router.GET("/readusers", controller.GetUser)
	router.PUT("/deleteuser", controller.DeleteUser)
	router.DELETE("/updateuser", controller.GetUserbyid)
}

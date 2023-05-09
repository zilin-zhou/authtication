package api

import (
	"application/controller"

	"github.com/gin-gonic/gin"
)

func InitAPI(engine *gin.Engine) {
	api := engine.Group("api")
	{
		api.POST("/register", controller.RegisterController{}.Register)
		api.POST("/login", controller.LoginControll{}.Login)
		// api.POST("/searchwitness", controller.SearchWitnessController{}.SearchWitness)
		// api.POST("/deleteuser", controller.DeleteUserController{}.DeleteUser)
	}
}

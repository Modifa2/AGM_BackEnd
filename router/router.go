package router

import (
	c "gopls-workspace/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	user := r.Group("api/user/")
	{
		user.POST("Register", c.Register)
		user.POST("CheckInUser", c.CheckInUser)
		user.POST("GetAllUsers", c.GetAllUsers)
		user.POST("ReturnUser", c.ReturnUser)
		user.POST("CheckNewUser", c.CheckNewUser)
		user.POST("GetAssociations", c.GetAssociations)
		user.POST("MarkAsDisplayed", c.MarkAsDisplayed)
	}
}	
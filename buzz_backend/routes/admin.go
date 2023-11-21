package routes

import (
	"github.com/athunlal/controls"
	"github.com/gin-gonic/gin"
)

func AdminRouts(c *gin.Engine) {
	admin := c.Group("/admin")
	{
		//Admin rounts
		admin.POST("/login", controls.AdminLogin)
		admin.POST("/signup", controls.AdminSignup)
		admin.GET("/logout", controls.AdminSignout)
		// admin.GET("/profile", middlereware.AdminAuth, controls.AdminProfile)
		admin.GET("/adminvalidate", controls.ValidateAdmin)

	}

}

package main

import (
	"github.com/athunlal/config"
	"github.com/athunlal/routes"
	"gorm.io/gorm"

	"github.com/athunlal/initializer"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func init() {
	initializer.LoadEnv()
	var err error
	config.DB, err = config.DBconnect()
	if err != nil {
		panic(err)
	}

	R.LoadHTMLGlob("templates/*.html")
}

var R = gin.Default()

func main() {

	gin.SetMode(gin.ReleaseMode)

	routes.AdminRouts(R)
	routes.UserRoutes(R)

	R.Run()
}

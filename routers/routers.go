package routers

import (
	v1 "passport.xinfos.com/api/v1"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	g.LoadHTMLGlob("web/templates/*")

	g.Static("/static", "./web/static")

	v1Api := g.Group("/v1")
	{
		v1Api.POST("/user/get", v1.GetUserInfoByID)
		v1Api.POST("/user/create", v1.CreateUser)
	}

	wp := g.Group("/")
	{
		wp.GET("/login", v1.ServiceLogin)
	}

	return g
}

package routers

import (
	v1 "passport.xinfos.com/api/v1"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	v1Api := g.Group("/v1")
	{
		v1Api.POST("/user/get", v1.GetUserInfoByID)
		v1Api.POST("/user/create", v1.CreateUser)
	}

	return g
}

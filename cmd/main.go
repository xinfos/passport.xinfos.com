package main

import (
	"passport.xinfos.com/driver"
	"passport.xinfos.com/pkg/logger"
	"passport.xinfos.com/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	driver.InitDB()
	driver.InitRedis()

	logger.InitLogger()

	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g = routers.Load(g)
	// ginpprof.Wrap(g)

	if err := g.Run(":8082"); err != nil {
		logger.F.Error("service start fail: " + err.Error())
	}
}

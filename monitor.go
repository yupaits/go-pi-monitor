package go_pi_monitor

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Config *AppConfig

func Run() {
	Config = initConfig()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/static", Config.Http.StaticPath)
	r.StaticFile("/favicon.ico", Config.Http.Favicon)
	r.LoadHTMLGlob(Config.Http.HtmlPathPattern)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/info", monitorInfoHandler)

	log.Printf("monitor启动成功, 端口: %s", Config.Http.Port)
	log.Fatal(r.Run(":" + Config.Http.Port))
}

func monitorInfoHandler(c *gin.Context) {
	//TODO 获取监控数据
}

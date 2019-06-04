package go_pi_monitor

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/static", "./ui/dist/static")
	r.StaticFile("/favicon.ico", "./ui/dist/favicon.ico")
	r.LoadHTMLGlob("./ui/dist/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/info", monitorInfoHandler)

	log.Printf("monitor启动成功, 端口: %s", "80")
	log.Fatal(r.Run(":80"))
}

func monitorInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, getMetrics())
}

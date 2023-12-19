package main

import (
	"plumpalbert/plumpifier/lib"
	"plumpalbert/plumpifier/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config := lib.ReadConfig("config.json")

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		routes.RouteCallback(c, config)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

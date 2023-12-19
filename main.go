package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/containrrr/shoutrrr"
	"github.com/gin-gonic/gin"
	"plumpalbert/plumpifier/lib"
)

func main() {
	config := lib.ReadConfig("config.json")

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		sendNotification(c, config)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func sendNotification(c *gin.Context, config lib.Config) {
	decoder := json.NewDecoder(c.Request.Body)

	var body lib.MovieDownloaded
	err := decoder.Decode(&body)

	if err != nil {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "could not decode request body",
		})
		return
	}

	sender, err := shoutrrr.CreateSender(config.Urls...)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not instantiate sender",
		})
		return
	}

	var tmp bytes.Buffer

	if body.EventType == "Download" {
		msg, err := template.New("msg").Parse(config.Radarr.Download)
		msg.Execute(&tmp, body)

		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "could not create template message",
			})
			return
		}
	}

	sender.Send(tmp.String(), nil)
}

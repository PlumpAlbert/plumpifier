package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/containrrr/shoutrrr"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", sendNotification)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func sendNotification(c *gin.Context) {
	urls := []string{
	}

	decoder := json.NewDecoder(c.Request.Body)

	var body MovieDownloaded
	err := decoder.Decode(&body)

	if err != nil {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "could not decode request body",
		})
		return
	}

	sender, err := shoutrrr.CreateSender(urls...)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not instantiate sender",
		})
		return
	}

	msg, err := template.New("msg").Parse("{{ .Movie.Title }} ({{ .Movie.Year }}) downloaded [{{ .File.Quality }}]")
	var tmp bytes.Buffer
	msg.Execute(&tmp, body)

	sender.Send(tmp.String(), nil)
}

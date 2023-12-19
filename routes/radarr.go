package routes

import (
	"io"
	"log"
	"net/http"
	"plumpalbert/plumpifier/lib"
	"plumpalbert/plumpifier/lib/radarr"

	"github.com/gin-gonic/gin"
)

func RouteCallback(ctx *gin.Context, config lib.Config) {
	buffer, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("[RADARR] Could not read request body")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not read request body",
			"error":   err,
		})
		return
	}

	webhookPayload, err := lib.DecodeBody[radarr.WebhookPayload](buffer)
	if err != nil {
		log.Println("[RADARR] Could not parse payload", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not parse payload",
			"error":   err,
		})
		return
	}

	switch webhookPayload.EventType {
	case "Download":
		{
			body, err := lib.DecodeBody[radarr.MovieDownloaded](buffer)
			if err != nil {
				log.Println("[RADARR] err:", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "could not parse download payload",
					"error":   err,
				})
				return
			}
			log.Println("[RADARR] body:", body)

			message, err := lib.PrepareTemplate(config.Radarr.Download, body)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "could not prepare template",
					"error":   err,
				})
				return
			}

			lib.SendNotification(message, config.Urls...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
			})
			return
		}
	}
}

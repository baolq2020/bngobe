package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../services"
	"encoding/base64"
	"io/ioutil"
)

func DetectImage(c *gin.Context) {

	file, _ , _ := c.Request.FormFile("file")

	content, _ := ioutil.ReadAll(file)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	
	mimeType := http.DetectContentType(content)

	switch mimeType {
		case "image/png":
			encoded = "data:image/png;base64," + encoded
		case "image/jpeg":
			encoded = "data:image/jpeg;base64," + encoded
		default:
			encoded = encoded
	}

	services.SendImagesToKafka(encoded)


	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
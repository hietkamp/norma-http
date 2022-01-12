package handlers

import "github.com/gin-gonic/gin"

func HandleMessage(c *gin.Context) {
	c.Writer.WriteHeader(200)
}

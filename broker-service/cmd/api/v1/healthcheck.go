package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAppStatus(c *gin.Context) {
	env := gin.H{
		"status": "available",
		"system_info": map[string]string{
			"environment": "Local",
			"version":     "1.0.0",
		},
	}
	c.IndentedJSON(http.StatusOK, env)
}

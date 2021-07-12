package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		var apikey string
		if strings.Contains(c.FullPath(), "/users") {
			apikey = "dooqo4idb84cekgo7t44w"
		} else if strings.Contains(c.FullPath(), "/content") {
			apikey = "eidfr8ca629sor6k9ukg7"
		}

		if c.GetHeader("apikey") == apikey {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid key",
			})
		}
	}
}

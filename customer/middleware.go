package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthorizationToken check token2019
func AuthorizationToken(c *gin.Context) {
	authKey := c.GetHeader("Authorization")
	if authKey != "token2019" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
}

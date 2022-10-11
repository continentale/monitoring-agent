package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthorizationV2(c *gin.Context) {
	// check if token is needed
	secret := c.GetHeader("secret")
	if viper.GetBool("server.useSecret") && viper.GetString("server.secret") != secret {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// check if endpoint is enabled
	endpoint := c.FullPath()[8:]

	if !viper.GetBool(endpoint + ".enabled") {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

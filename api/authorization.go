package api

import (
	"net/http"

	"github.com/continentale/monitoring-agent/config"
	"github.com/gin-gonic/gin"
)

func AuthorizationV2(c *gin.Context) {
	// check if token is needed
	secret := c.GetHeader("secret")
	if config.ConfigStruct.Server.UseSecret && config.ConfigStruct.Server.Secret != secret {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// check if endpoint is enabled
	endpoint := c.FullPath()[8:]

	switch endpoint {
	case "version":
		if !config.ConfigStruct.Endpoints.Version.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "mem":
		if !config.ConfigStruct.Endpoints.Mem.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "procs":
		if !config.ConfigStruct.Endpoints.Procs.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "disks":
		if !config.ConfigStruct.Endpoints.Disks.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "load":
		if !config.ConfigStruct.Endpoints.Load.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "time":
		if !config.ConfigStruct.Endpoints.Time.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "cpus":
		if !config.ConfigStruct.Endpoints.Cpus.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "file":
		if !config.ConfigStruct.Endpoints.File.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	case "exec":
		if !config.ConfigStruct.Endpoints.Exec.Enabled {
			c.AbortWithStatus(http.StatusForbidden)
		}
	default:
		c.Data(http.StatusForbidden, "text/html", []byte("Unsupported endpoint"))
	}
}

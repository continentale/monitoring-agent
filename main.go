/**
 * @package   monitoring-agent
 * @copyright monitoring-agent contributors
 * @license   GNU Affero General Public License (https://www.gnu.org/licenses/agpl-3.0.de.html)
 * @authors   https://github.com/continentale/monitoring-agent/graphs/contributors
 * @todo lots of documentation
 *
 *
 * Monitoring Agent with REST-API for Linux, Windows and osX
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/continentale/monitoring-agent/api"
	"github.com/continentale/monitoring-agent/config"
	"github.com/continentale/monitoring-agent/paths"
	"github.com/continentale/monitoring-agent/types"
)

var (
	// VERSION is The specific Version for the API
	VERSION = "0.0.1"
	// GITCOMMIT is the commit for the build
	GITCOMMIT = "HEAD"
	// BUILDDATE is the date from the build
	BUILDDATE = ""
)

func main() {
	paths.InitCommon()
	paths.InitOSSpecific()
	config.InitConfig()

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	v2 := router.Group("/api/v2")

	v2.Use(api.AuthorizationV2)

	// TODO: transfer it to the API package with the global vars from main
	v2.GET("version", func(c *gin.Context) {
		versions := types.Version{
			Commit:  GITCOMMIT,
			Version: VERSION,
			Date:    BUILDDATE,
		}

		c.JSON(http.StatusOK, versions)
	})

	v2.GET("mem", api.GetMemory)
	v2.GET("procs", api.GetProcs)
	v2.GET("disks", api.GetDisk)
	v2.GET("load", api.GetLoad)
	v2.GET("time", api.GetTime)
	v2.GET("cpus", api.GetCPU)

	v2.GET("file", api.ShowFile)
	v2.GET("exec", api.ExecCommand)

	log.Println("Running Version:", VERSION, "with commit tag", GITCOMMIT, "build on", BUILDDATE)
	log.Println("Running Timeouts:", viper.GetDuration("timeouts"))

	listenAdress := config.ConfigStruct.Server.Address
	if listenAdress == "" {
		listenAdress, _ = os.Hostname()
	}

	s := &http.Server{
		Addr:           ":20480",
		Handler:        router,
		ReadTimeout:    time.Duration(config.ConfigStruct.Server.Timeouts) * time.Second,
		WriteTimeout:   time.Duration(config.ConfigStruct.Server.Timeouts) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if config.ConfigStruct.Server.Protocol == "https" {
		log.Println("Server listening on", "https://"+fmt.Sprintf("%s:%d", listenAdress, config.ConfigStruct.Server.Port))
		s.ListenAndServeTLS(config.ConfigStruct.Server.Certificate, config.ConfigStruct.Server.Key)
	} else {
		log.Println("Server listening on", "http://"+fmt.Sprintf("%s:%d", listenAdress, config.ConfigStruct.Server.Port))
		s.ListenAndServe()
	}
}

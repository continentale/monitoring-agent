package api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/continentale/monitoring-agent/paths"
	"github.com/gin-gonic/gin"
)

var (
	memRegexp = regexp.MustCompile(`\((.*)\)`)
)

func GetMemory(c *gin.Context) {
	file, _ := os.Open(paths.ProcFilePath("meminfo"))

	defer file.Close()

	mapping, _ := parseMemInfo(file)

	fmt.Println(mapping)

	c.JSON(http.StatusOK, mapping)
}

func parseMemInfo(r io.Reader) (map[string]float64, error) {
	memMapping := map[string]float64{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		keyValues := strings.Fields(line)

		// remove whitespace prefix
		memKey := keyValues[0][:len(keyValues[0])-1]

		// rename () to _
		memKey = memRegexp.ReplaceAllString(memKey, "_$1")
		memValue, _ := strconv.ParseFloat(keyValues[1], 64)

		switch len(keyValues) {
		case 2:
		case 3:
			switch keyValues[2] {
			case "kB":
				memValue = memValue * (1 << 10)
			case "mB":
				memValue = memValue * (1 << 20)
			case "gB":
				memValue = memValue * (1 << 30)
			}
		}

		memMapping[memKey] = memValue
	}

	return memMapping, nil
}

package types

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

type FilterEntryOperator string

const (
	FILTER_ENTRY_EXACT FilterEntryOperator = ":"
	FILTER_ENTRY_REGEX FilterEntryOperator = ":*"
)

type FilterEntry struct {
	Operator FilterEntryOperator
	Value    string
}

type Filter struct {
}

func ApplyFilter(data []byte, filter string, endpoint string) []byte {
	filterParts := strings.Split(filter, "||")
	gjsonData := gjson.ParseBytes(data)
	result := make([]interface{}, 0)
	// iterate over the output from the endpoint
	for _, value := range gjsonData.Array() {
		// iterate over all given filtered parts
		for _, filterPart := range filterParts {
			// is regex search
			if strings.Contains(filterPart, string(FILTER_ENTRY_REGEX)) {
				fieldName := filterPart[:strings.IndexAny(filterPart, string(FILTER_ENTRY_REGEX))]
				searchValue := filterPart[strings.IndexAny(filterPart, string(FILTER_ENTRY_REGEX))+2:]
				pattern := regexp.MustCompile(searchValue)

				if pattern.MatchString(value.Get(fieldName).String()) {
					parseEndpoint(endpoint, value.Raw, &result)
				}
			} else if strings.Contains(filterPart, string(FILTER_ENTRY_EXACT)) {
				fieldName := filterPart[:strings.IndexAny(filterPart, string(FILTER_ENTRY_EXACT))]
				searchValue := filterPart[strings.IndexAny(filterPart, string(FILTER_ENTRY_EXACT))+1:]

				if value.Get(fieldName).String() == searchValue {
					parseEndpoint(endpoint, value.Raw, &result)
				}
			}
		}
	}

	dataResult, _ := json.Marshal(result)
	return dataResult
}

func parseEndpoint(endpoint string, value string, result *[]interface{}) {
	switch endpoint {
	case "disks":
		*result = append(*result, parseDisk(value))
	case "procs":
		*result = append(*result, parseProcs(value))
	}
}

func parseDisk(data string) Disks {
	var typeValue Disks
	json.Unmarshal([]byte(data), &typeValue)
	return typeValue
}

func parseProcs(data string) Procs {
	var typeValue Procs
	json.Unmarshal([]byte(data), &typeValue)
	return typeValue
}

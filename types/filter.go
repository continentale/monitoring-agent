package types

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

type FilterEntryOperator string

const (
	FILTER_ENTRY_EXACT        FilterEntryOperator = ":"
	FILTER_ENTRY_REGEX        FilterEntryOperator = ":*"
	FILTER_ENTRY_NEGATE_REGEX FilterEntryOperator = ":#"
	FILTER_ENTRY_NOT          FilterEntryOperator = ":!"
)

func ApplyFilter(data []byte, filter []string, endpoint string) []byte {
	gjsonData := gjson.ParseBytes(data)
	result := make([]interface{}, 0)

	// iterate over the output from the endpoint
	for _, value := range gjsonData.Array() {
		isAllMatching := true
		isLastFilter := false
		// iterate over all given filtered parts
		for i, filterPart := range filter {
			// check if the given filter is the last one to guarantee that all filters are applied as true
			if i == len(filter)-1 {
				isLastFilter = true
			}

			// is regex search
			if strings.Contains(filterPart, string(FILTER_ENTRY_REGEX)) {
				fieldName := filterPart[:strings.IndexAny(filterPart, string(FILTER_ENTRY_REGEX))]
				searchValue := filterPart[strings.IndexAny(filterPart, string(FILTER_ENTRY_REGEX))+2:]
				pattern := regexp.MustCompile(searchValue)
				if value.Get(fieldName).IsArray() {
					if !matchRegexpArray(pattern, value.Get(fieldName).Array()) {
						isAllMatching = false
					}
				} else {
					if !pattern.MatchString(value.Get(fieldName).Raw) {
						isAllMatching = false
					}
				}
			} else if strings.Contains(filterPart, string(FILTER_ENTRY_NEGATE_REGEX)) {
				fieldName := filterPart[:strings.IndexAny(filterPart, string(FILTER_ENTRY_NEGATE_REGEX))]
				searchValue := filterPart[strings.IndexAny(filterPart, string(FILTER_ENTRY_NEGATE_REGEX))+2:]
				pattern := regexp.MustCompile(searchValue)

				if value.Get(fieldName).IsArray() {
					if matchRegexpArray(pattern, value.Get(fieldName).Array()) {
						isAllMatching = false
					}
				} else {
					if pattern.MatchString(value.Get(fieldName).String()) {
						isAllMatching = false
					}
				}
			} else if strings.Contains(filterPart, string(FILTER_ENTRY_NOT)) {

				fieldName := filterPart[:strings.IndexAny(filterPart, string(FILTER_ENTRY_NOT))]
				searchValue := filterPart[strings.IndexAny(filterPart, string(FILTER_ENTRY_NOT))+2:]
				if value.Get(fieldName).IsArray() {
					if isInArray(searchValue, value.Get(fieldName).Array()) {
						isAllMatching = false
					}
				} else {
					if value.Get(fieldName).String() != searchValue {
						isAllMatching = false
					}
				}
			} else { // match exact, but must be stand at the bottom
				fieldName := filterPart[:strings.IndexAny(filterPart, string(FILTER_ENTRY_EXACT))]
				searchValue := filterPart[strings.IndexAny(filterPart, string(FILTER_ENTRY_EXACT))+1:]

				if value.Get(fieldName).IsArray() {
					if !isInArray(searchValue, value.Get(fieldName).Array()) {
						isAllMatching = false
					}
				} else {
					if value.Get(fieldName).String() != searchValue {
						isAllMatching = false
					}
				}
			}

			// apply filter only if its the last filter and all previous ones are matching
			if isAllMatching && isLastFilter {
				parseEndpoint(endpoint, value.Raw, &result)
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

func isInArray(search string, field []gjson.Result) bool {
	for _, v := range field {
		if search == v.String() {
			return true
		}
	}
	return false
}

func matchRegexpArray(search *regexp.Regexp, field []gjson.Result) bool {
	for _, v := range field {
		if search.MatchString(v.String()) {
			return true
		}
	}
	return false
}

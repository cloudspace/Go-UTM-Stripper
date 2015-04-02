package main // import "github.com/cloudspace/Go-UTM-Stripper"

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"strings"
)

// Command-line flags.
var (
	urlStringToStrip = flag.String("url", "", "URL to Strip")
)

func main() {
	flag.Parse()

	urlToStrip, err := url.Parse(*urlStringToStrip)
	if err != nil {
		fmt.Println(getJSONError(err))
	}

	params := urlToStrip.Query()

	for eachKey := range params {
		if len(eachKey) < 4 {
			continue
		}
		if strings.EqualFold(eachKey[:4], "utm_") {
			params.Del(eachKey)
		}
	}

	urlToStrip.RawQuery = params.Encode()
	result := make(map[string]interface{}, 0)
	result["result"] = urlToStrip.String()

	fmt.Println(asJSON(result))
}

func getJSONError(myError error) string {

	errorJSON := make(map[string]interface{})
	errorJSON["error"] = myError.Error()
	jsonData, err := json.Marshal(errorJSON)
	if err != nil {
		return "{\"error\": \"There was an error generatoring the error.. goodluck\"}"
	}
	return string(jsonData)
}

func asJSON(anything interface{}) string {

	jsonData, err := json.Marshal(anything)
	if err != nil {
		return getJSONError(err)
	}
	return string(jsonData)
}

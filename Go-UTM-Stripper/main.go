package main // import "github.com/cloudspace/Go-UTM-Stripper"

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(errorStringAsJSON(fmt.Sprintf("Must have 1 argument: it has %v arguments", len(os.Args)-1)))
		return
	}

	urlStringToStrip := os.Args[1]
	urlToStrip, err := url.Parse(urlStringToStrip)
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

func errorStringAsJSON(errorString string) string {

	return "{\"error\": \"" + errorString + "\"}"
}

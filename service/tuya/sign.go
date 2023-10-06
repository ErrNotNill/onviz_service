package tuya

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func Sign() {
	httpMethod := http.MethodGet // Replace with your HTTP method
	contentSHA256 := "..."       // Replace with the SHA-256 hash of the request payload
	headers := map[string]string{
		"Authorization": "Bearer " + Token,
		// Add other headers as needed
	}
	apiURL := "https://api.tuya.com/v1.0/your/endpoint" // Replace with your API endpoint

	stringToSign := GenerateStringToSign(httpMethod, contentSHA256, headers, apiURL)
	fmt.Println("String to Sign:", stringToSign)
}

func GenerateStringToSign(httpMethod, contentSHA256 string, headers map[string]string, apiURL string) string {
	var headerKeys []string
	for key := range headers {
		headerKeys = append(headerKeys, key)
	}
	sort.Strings(headerKeys)

	headerStrings := []string{}
	for _, key := range headerKeys {
		headerStrings = append(headerStrings, key+": "+headers[key])
	}

	urlInfo, _ := url.Parse(apiURL)

	stringToSign := httpMethod + "\n" +
		contentSHA256 + "\n" +
		strings.Join(headerStrings, "\n") + "\n" +
		urlInfo.Host + urlInfo.Path

	return stringToSign
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// Construct the SMS using https://golang.org/pkg/net/url/#Values
	msgData := url.Values{}
	msgData.Set("api_key", "")    // Put Nexmo api_key here
	msgData.Set("api_secret", "") // Put Nexmo api_secret here
	msgData.Set("to", "")         // Number format for this field is 447700900000
	msgData.Set("from", "NEXMO")
	msgData.Set("text", "Hello from Nexmo")
	msgDataReader := *strings.NewReader(msgData.Encode())
	// Set up HTTP client and HTTP POST request
	urlStr := "https://rest.nexmo.com/sms/json"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Make the request
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data)
		}
	} else {
		fmt.Println(resp.Status)
	}
}

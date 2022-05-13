package request

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// Makes HTTP GET request with requestBody 
func MakeGetRequest(url string, requestBody []byte) []byte {
	return makeRequest(url, requestBody, "GET")
}

// Makes HTTP POST request with requestBody 
func MakePostRequest(url string, requestBody []byte) []byte {
	return makeRequest(url, requestBody, "POST")
}

func makeRequest(url string, requestBody []byte, method string) []byte {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatal(err)
	}

    request.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    
	response, err := client.Do(request)

    if err != nil {
        log.Fatal(err)
    }

    defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return responseBody;
}
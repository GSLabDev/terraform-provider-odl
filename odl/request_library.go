package odl

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// PostRequest ... post Request to odl
func (c *Config) PostRequest(requestPath string, body map[string]interface{}) (*http.Response, error) {
	var bodyBuffer *bytes.Buffer
	if body != nil {
		buff, _ := json.Marshal(&body)
		var jsonStr = []byte(buff)
		bodyBuffer = bytes.NewBuffer(jsonStr)
	}
	request, err := http.NewRequest("POST", c.URL+requestPath, bodyBuffer)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return nil, fmt.Errorf("[ERROR] Error in creating request %s", err.Error())
	}
	return c.getResponse(request)
}

// GetRequest ... get Request to odl
func (c *Config) GetRequest(requestPath string) (*http.Response, error) {
	request, err := http.NewRequest("GET", c.URL+requestPath, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return nil, fmt.Errorf("[ERROR] Error in creating request %s", err.Error())
	}
	return c.getResponse(request)
}

func (c *Config) getResponse(request *http.Request) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	request.SetBasicAuth(c.Username, c.Password)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: tr}
	return client.Do(request)
}

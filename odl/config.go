package odl

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Config ... ODL configuration details
type Config struct {
	ServerIP string
	Port     int
	Username string
	Password string
	URL      string
}

//checkConnection ... checkConnection with server.
func (c *Config) checkConnection() (*Config, error) {
	c.URL = "http://" + c.ServerIP + ":" + strconv.Itoa(c.Port) + "/"
	log.Printf("[DEBUG] Checking url")
	_, err := url.Parse(c.URL)
	if err != nil {
		log.Println("[Error] URL is not in correct format")
		return nil, err
	}
	log.Printf("[DEBUG] Creating request")
	request, err := http.NewRequest("GET", c.URL+"auth/v1/users", nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	request.SetBasicAuth(c.Username, c.Password)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: tr}
	log.Printf("[DEBUG] Checking users")
	resp, err := client.Do(request)
	if err != nil {
		log.Println(" [ERROR] Connecting to server ", err)
		return nil, fmt.Errorf("[ERROR] Server IP or port is incorrect")

	}
	if resp.Status == "200 OK" {
		return c, nil
	}
	return nil, fmt.Errorf("[ERROR] Incorrect Username or Password %s", err.Error())
}

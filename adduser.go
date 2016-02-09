package cloudmqtt

import (
	"bytes"
	"fmt"
	"net/http"
)

func (client *Client) AddUser(username, password string) (err error) {
	url := "https://api.cloudmqtt.com/user"

	jsonBody := []byte(`{"username":"` + username + `", "password":"` + password + `"}`)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(client.CloudMQTTUsername, client.CloudMQTTPassword)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf(resp.Status)
	}

	return err
}

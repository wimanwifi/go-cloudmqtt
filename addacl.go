package cloudmqtt

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func (client *Client) AddACL(username string, topic string, read bool, write bool) (err error) {
	url := "https://api.cloudmqtt.com/acl"

	jsonBody := []byte(`{"username":"` + username + `", "topic":"` + topic + `", "read": ` + strconv.FormatBool(read) + `, "write": ` + strconv.FormatBool(write) + `}`)
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

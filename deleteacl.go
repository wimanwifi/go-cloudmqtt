package cloudmqtt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) DeleteACL(username, topic string) error {
	url := "https://api.cloudmqtt.com/acl"

	s := UserACL{
		Username: username,
		Topic:    topic,
	}
	body := new(bytes.Buffer)
	enc := json.NewEncoder(body)
	if err := enc.Encode(s); err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", url, body)
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

	return nil
}

package cloudmqtt

import (
	"fmt"
	"net/http"
)

func (client *Client) DeleteUser(username string) (err error) {
	url := "https://api.cloudmqtt.com/user"

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", url, username), nil)
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

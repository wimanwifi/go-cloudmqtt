package cloudmqtt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (client *Client) ListUsers() (userList UserList, err error) {
	url := "https://api.cloudmqtt.com/user"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return userList, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(client.CloudMQTTUsername, client.CloudMQTTPassword)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return userList, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return userList, err
	}

	err = json.Unmarshal(body, &userList)
	return userList, err
}

package cloudmqtt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (client Client) ListACL() (ACLList, error) {
	url := "https://api.cloudmqtt.com/acl"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(client.CloudMQTTUsername, client.CloudMQTTPassword)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var aclList ACLList
	if err := json.Unmarshal(body, &aclList); err != nil {
		return nil, err
	}

	return aclList, nil
}

package cloudmqtt

type Client struct {
	CloudMQTTUsername string
	CloudMQTTPassword string
}

func New(username string, password string) *Client {
	client := new(Client)
	client.CloudMQTTUsername = username
	client.CloudMQTTPassword = password
	return client
}

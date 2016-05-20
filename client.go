package cloudmqtt

type Client struct {
	CloudMQTTUsername string
	CloudMQTTPassword string
}

func New(username, password string) Client {
	return Client{
		CloudMQTTUsername: username,
		CloudMQTTPassword: password,
	}
}

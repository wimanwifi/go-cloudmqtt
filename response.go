package cloudmqtt

type UserList []User

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ACLList []ACL

type ACL struct {
	Topic    string `json:"topic"`
	Read     bool   `json:"read"`
	Write    bool   `json:"write"`
	Username string `json:"username"`
}

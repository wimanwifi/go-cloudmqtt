package cloudmqtt

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserList []User

type UserACL struct {
	Username string `json:"username"`
	Topic    string `json:"topic"`
}

type ACLList []ACL

type ACL struct {
	Topic    string `json:"topic"`
	Read     bool   `json:"read"`
	Write    bool   `json:"write"`
	Username string `json:"username"`
}

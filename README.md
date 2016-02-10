[![GoDoc](https://godoc.org/github.com/wimanwifi/go-cloudmqtt?status.svg)](https://godoc.org/github.com/wimanwifi/go-cloudmqtt)

# Pure Go CloudMQTT API Client

go-cloudmqtt is intended to be used as a Go package. It does not include a command to run the interpreter.

## Installation

To start using the library, run:
```sh
go get github.com/wimanwifi/go-cloudmqtt
```
## Example

A simple example that loads & runs functions:
```go
package main

import (
	"fmt"

	"github.com/wimanwifi/go-cloudmqtt"
)

func main() {
	// Initialize with CloudMQTT username and password
	api := cloudmqtt.New("CLOUDMQTT_USERNAME", "CLOUDMQTT_PASSWORD")

	// Add user with username and password as string
	if err := api.AddUser("username", "password"); err == nil {
		fmt.Println("User added successfully")
	}

	// Add ACL with username, password as string and read, write as bool
	if err := api.AddACL("username", "topic", true, true); err == nil {
		fmt.Println("ACL added successfully")
	}

	// List users
	fmt.Printf("\nUsers:\n")
	if users, err := api.ListUsers(); err == nil {
		for _, user := range users {
			fmt.Printf("%s => %s\n", user.Username, user.Password)
		}
	}

	// List ACL
	fmt.Printf("\nACL:\n")
	if aclList, err := api.ListACL(); err == nil {
		for _, acl := range aclList {
			fmt.Printf("%s => %s [r: %t / w: %t]\n", acl.Username, acl.Topic, acl.Read, acl.Write)
		}
	}
	fmt.Printf("\n")

	// Delete ACL matching username and topic as string
	if err := api.DeleteACL("username", "topic"); err == nil {
		fmt.Println("ACL deleted successfully")
	}

	// Delete user matching username as string
	if err := api.DeleteUser("username"); err == nil {
		fmt.Println("User deleted successfully")
	}
}
```

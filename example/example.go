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
	if acllist, err := api.ListACL(); err == nil {
		for _, acl := range acllist {
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

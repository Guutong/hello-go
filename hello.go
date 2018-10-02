package main

import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println("Hello, ", getUsers())
}

func getUsers() string {
	u, err := user.Current()
	if err != nil {
		return "World"
	}
	return u.Name
}

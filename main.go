package main

import (
	"fmt"
)

func main() {
	userID := "123"
	user, address, err := searchUserAndAddressByUserID(userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Address: %+v\n", address)
}

type (
	user struct {
		id     string
		name   string
		status string
	}
	address struct {
		country       string
		city          string
		StreetAddress string
	}
)

func searchUserAndAddressByUserID(userID string) (*user, *address, error) {
	user, err := getUser(userID)
	if err != nil {
		return nil, nil, err
	}
	if user.status != "active" {
		// This should be return nil, nil, ErrNotFound.
		return nil, nil, nil
	}

	address, err := getAddress()
	if err != nil {
		return nil, nil, err
	}

	return user, address, nil
}

func getUser(userID string) (*user, error) {
	return &user{
		id:     userID,
		name:   "john",
		status: "active",
	}, nil
}

func getAddress() (*address, error) {
	return &address{
		country:       "Japan",
		city:          "Tokyo",
		StreetAddress: "1-1",
	}, nil
}


package main

import (
	"fmt"
)

func main() {
	u := user{name: "John"}
	fmt.Printf("u.greet(): %v\n", u.greet())
	u.update("Doe")
	fmt.Printf("u.greet(): %v\n", u.greet())
}

type user struct {
	name string
}

func (u user) greet() string {
	return fmt.Sprintf("I'm %s", u.name)
}

func (u *user) update(after string) {
	u.name = after
}

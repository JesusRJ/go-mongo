package tests_test

import (
	"strconv"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

type Config struct {
	ID      any
	Pets    int
	Address string
}

func GetUser(name string, config Config) *User {
	user := User{
		Name:    name,
		Address: "5th Avenue, number 123",
	}

	user.ID = config.ID
	user.Address = config.Address

	for i := 0; i < config.Pets; i++ {
		user.Pets = append(user.Pets, &Pet{Name: name + "_pet_" + strconv.Itoa(i+1)})
	}

	return &user
}

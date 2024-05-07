package tests_test

import (
	"fmt"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

type Config struct {
	ID      any
	Pets    int
	Phones  int
	Address *Address
	Company *Company
}

func GetUser(name string, config Config) *User {
	user := User{
		Name:  name,
		Phone: PhoneNumbers(config.Phones),
		Pets:  Pets(config.Pets),
	}

	user.ID = config.ID

	if config.Address != nil {
		user.Address = config.Address
	}

	if config.Company != nil {
		user.Company = config.Company
	}

	return &user
}

func PhoneNumbers(c int) (phones []*Phone) {
	if c == 0 {
		return nil
	}
	for x := range c {
		phones = append(phones, &Phone{Number: fmt.Sprintf("(00) 0000-000%v", x)})
	}
	return phones
}

func Pets(c int) (pets []*Pet) {
	if c == 0 {
		return nil
	}
	for x := range c {
		pets = append(pets, &Pet{Name: fmt.Sprintf("pet_%v", x)})
	}
	return pets
}

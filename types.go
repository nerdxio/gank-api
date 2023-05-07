package main

import "math/rand"

type Account struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

// constractor
func NewAccount(firstname, lastName string) *Account {
	return &Account{
		Id:        rand.Intn(10000),
		FirstName: firstname,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
	}
}

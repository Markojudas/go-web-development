package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

//function that stores the user
func StoreUsers(m map[string]User) {
	fp, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	json.NewEncoder(fp).Encode(m)
}

//function that loads the user
func LoadUsers() map[string]User {
	m := make(map[string]User)

	fp, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return m //returns empty list
	}
	defer fp.Close()

	err = json.NewDecoder(fp).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}

	return m
}

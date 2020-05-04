package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func userInfo(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	user := &User{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
func main() {
	user, err := userInfo("manutvm")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("%+v\n", user)
}

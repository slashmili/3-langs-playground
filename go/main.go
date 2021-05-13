package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}
type Users struct {
	Collection []User
}

func main() {

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/stargazers", "rust-lang-nursery", "rust-cookbook"))
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	users := make([]User, 0)
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatalln(err)
	}
	filtered_users := make([]User, 0)
	for _, u := range users {
		if strings.HasPrefix(u.Login, "s") {
			filtered_users = append(filtered_users, u)
		}
	}
	fmt.Println(filtered_users)
}

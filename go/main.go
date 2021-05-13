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

func FetchUser(user_name string, input chan UserResult) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", user_name))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		input <- UserResult{Error: err}
	}
	input <- UserResult{UserInfo: user}
}

type UserResult struct {
	UserInfo User
	Error    error
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

	input := make(chan UserResult)

	for _, u := range users {
		if strings.HasPrefix(u.Login, "s") {
			go FetchUser(u.Login, input)
		}
	}

	filtered_users := make([]User, 0)

	for res := range input {
		if res.Error == nil {
			filtered_users = append(filtered_users, res.UserInfo)
			fmt.Println(res)
		}
	}
	fmt.Println(filtered_users)
	// Problems:
	// never reaches this part of the code
	// I have to use WaitingGroup and create another channel to
	// receive the message and decrease WaitingGroup counter.
}

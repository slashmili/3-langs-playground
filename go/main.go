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
	filter_func := func(u User) bool { return !strings.HasPrefix(u.Login, "f") }
	filtered_users := filter(users, filter_func)
	fmt.Println(filtered_users)
}

func filter(ss []User, test func(User) bool) (ret []User) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}

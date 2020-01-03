package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)


type Post struct {
	UserID int `json: "userId"`
	ID int `json: "id"`
	Title string `json: "title"`
	Body string `json: "body"`
}

func main() {
	client := &http.Client{}

	url := "http://jsonplaceholder.typicode.com/posts"

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error())
	}

	var posts [] Post

	err = json.NewDecoder(resp.Body).Decode(&posts)

	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(posts); i++ {
		fmt.Println(posts[i].ID)
		fmt.Println(posts[i].Title)
		fmt.Println(posts[i].Body)
	}

}
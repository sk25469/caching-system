package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Uncomment this line when testing locally
const baseUrl = "http://localhost:3000"

func main() {
	// id, path, caching
	id := flag.Int("id", 1, "the id of todo or posts")
	path := flag.String("path", "todos", "the path namely todos or posts")
	caching := flag.String("caching", "", "whether caching enabled or disabled")

	// Parse command-line arguments
	flag.Parse()

	var url string
	if *caching == "" {
		if *path == "posts" {
			url = fmt.Sprintf(baseUrl+"/posts/%d", *id)
		} else if *path == "todos" {
			url = fmt.Sprintf(baseUrl+"/todos/%d", *id)
		}

	} else {
		if *path == "posts" {
			url = fmt.Sprintf(baseUrl+"/caching/posts=%v", *caching)
		} else if *path == "todos" {
			url = fmt.Sprintf(baseUrl+"/caching/todos=%v", *caching)
		} else if *path == "all" {
			url = fmt.Sprintf(baseUrl+"/caching=%v", *caching)
		}
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch data from API: " + err.Error())
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read response body: ", err.Error())
		return
	}

	fmt.Println(string(body))

}

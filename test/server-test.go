package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sk25469/momoney-backend-assignment/utils"
)

// Uncomment this line when testing from local server
// const baseUrl = "http://localhost:3000"

// Use this url when testing from deployment
const baseUrl = "http://20.207.85.42:3000"

func main() {
	// id gives the id for todos or posts
	id := flag.Int("id", 1, "the id of todos or posts")

	// path can be todos or posts
	path := flag.String("path", "todos", "the path namely todos or posts")

	// caching can take 2 values - true or false
	caching := flag.String("caching", "", "whether caching enabled or disabled")

	// Parse command-line arguments
	flag.Parse()

	// check if id is not an integer
	if utils.CheckVariableType(*id) != "int" {
		log.Fatal("id can only be an integer")
		return
	}

	// check if path only has string and takes posts and todos values
	if err := utils.CheckTypeAndValues(*path, "string", []string{"posts", "todos", "all"}); err != nil {
		log.Fatal("path can only takes values posts, todos and all")
		return
	}

	// check if caching only has string and takes true and false values
	if err := utils.CheckTypeAndValues(*caching, "string", []string{"true", "false", ""}); err != nil {
		log.Fatal("caching can only takes values true and false")
		return
	}

	var url string

	// when caching is empty, the requests are directly cached, with respective paths
	if *caching == "" {
		if *path == "posts" {
			url = fmt.Sprintf(baseUrl+"/posts/%d", *id)
		} else if *path == "todos" {
			url = fmt.Sprintf(baseUrl+"/todos/%d", *id)
		} else {
			log.Fatal("can't use all tag with caching tag")
			return
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

	// Makes a get request for the url
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch data from API: " + err.Error())
		return
	}

	// Make sure to close the response body before the program ends
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read response body: ", err.Error())
		return
	}

	fmt.Println(string(body))

}

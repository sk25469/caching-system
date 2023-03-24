package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sk25469/momoney-backend-assignment/models"
	"github.com/sk25469/momoney-backend-assignment/utils"
)

// PostCache stores the cached requests for all the posts
var postCache map[int]models.Post = make(map[int]models.Post)

// Todo Cache stores the cached requests for all the todos
var todoCache map[int]models.Todo = make(map[int]models.Todo)

// Everytime a request goes through the middleware, it will check if the id
// already exist in their respective todos, if they exist, then it returns the
// cached response already present in the cache.
// If the id doesn't exist, it sends the requests for that id, gets the response and
// stores inside the cache
func InitMiddleWare(idParam int, requestType utils.Request) (models.Entity, error) {

	if requestType == utils.Post {
		// if the id already exist in cache
		if _, ok := postCache[idParam]; ok {
			log.Printf("RETRIEVING POSTS ---> Id: [%v] already present in cache, returning data from cache", idParam)
			return postCache[idParam], nil
		} else {
			log.Printf("RETRIEVING POSTS ---> Id: [%v] not present in cache, sending request", idParam)

			posts, err := SendPostsRequest(idParam, utils.PostUrl)
			if err != nil {
				log.Printf("Error occured while getting posts: [%v}", err)
				return models.Post{}, err
			}
			log.Printf("RETRIEVING POSTS ---> Id: [%v] retrieved data from request, storing in cache", idParam)

			postCache[idParam] = posts
			return posts, nil

		}
	} else {
		if _, ok := todoCache[idParam]; ok {
			log.Printf("RETRIEVING TODOS ---> Id: [%v] already present in cache, returning data from cache", idParam)
			return todoCache[idParam], nil
		} else {
			log.Printf("RETRIEVING TODOS ---> Id: [%v] not present in cache, sending request", idParam)

			todos, err := SendTodosRequest(idParam, utils.TodoUrl)
			if err != nil {
				log.Printf("Error occured while getting todos: [%v}", err)
				return models.Post{}, err
			}
			log.Printf("RETRIEVING TODOS ---> Id: [%v] retrieved data from request, storing in cache", idParam)

			todoCache[idParam] = todos
			return todos, nil
		}
	}
}

func SendPostsRequest(idParam int, requestUrl string) (models.Post, error) {
	url := fmt.Sprintf(requestUrl+"%d", idParam)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch data from API: %v\n", err)
		return models.Post{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return models.Post{}, err
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("Failed to unmarshal post response body: %v\n", err)
		return models.Post{}, err
	}
	return post, nil
}

func SendTodosRequest(idParam int, requestUrl string) (models.Todo, error) {
	url := fmt.Sprintf(requestUrl+"%d", idParam)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch data from API: %v\n", err)
		return models.Todo{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return models.Todo{}, err
	}

	var todo models.Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Printf("Failed to unmarshal todo response body: %v\n", err)
		return models.Todo{}, err
	}
	return todo, nil
}

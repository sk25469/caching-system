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

// TodoCache stores the cached requests for all the todos
var todoCache map[int]models.Todo = make(map[int]models.Todo)

// Everytime a request goes through the middleware, it will check if the id
// already exist in their respective todos, if they exist, then it returns the
// cached response already present in the cache.
// If the id doesn't exist, it sends the requests for that id, gets the response and
// stores inside the cache
func InitMiddleWare(idParam int, requestType utils.Request, cachingEnabled bool) (models.Entity, error) {
	switch requestType {
	case utils.Post:
		if !cachingEnabled {
			posts, err := HandlePostsRequestWhenCachingDisabled(idParam)
			if err != nil {
				log.Printf("Error occured while handling disabled caching for posts: [%v]", err)
				return models.Post{}, err
			}
			return posts, nil
		}
		// if the id already exist in postCache
		if _, ok := postCache[idParam]; ok {
			log.Printf("RETRIEVING POSTS ---> Id: [%v] already present in cache, returning data from cache", idParam)
			return postCache[idParam], nil
		} else {
			posts, err := HandlePostsRequestForFirstTime(idParam)
			if err != nil {
				log.Printf("Error occured while getting posts: [%v}", err)
				return models.Post{}, err
			}
			return posts, nil

		}
	case utils.Todo:
		if !cachingEnabled {
			todos, err := HandlePostsRequestWhenCachingDisabled(idParam)
			if err != nil {
				log.Printf("Error occured while handling disabled caching for todos: [%v]", err)
				return models.Todo{}, err
			}
			return todos, nil
		}

		// if the id already exist in todoCache
		if _, ok := todoCache[idParam]; ok {
			log.Printf("RETRIEVING TODOS ---> Id: [%v] already present in cache, returning data from cache", idParam)
			return todoCache[idParam], nil
		} else {
			todos, err := HandleTodosRequestForFirstTime(idParam)
			if err != nil {
				log.Printf("Error occured while getting todos: [%v}", err)
				return models.Todo{}, err
			}
			return todos, nil
		}
	default:
		return models.Post{}, nil
	}
}

// When a particular id is requested for the first time and caching is not disabled, this function runs
// for the posts, requests the data and stores in postCache
func HandlePostsRequestForFirstTime(idParam int) (models.Post, error) {
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

// When a particular id is requested for the first time and caching is not disabled, this function runs
// for the todos, requests the data and stores in todoCache
func HandleTodosRequestForFirstTime(idParam int) (models.Todo, error) {
	log.Printf("RETRIEVING TODOS ---> Id: [%v] not present in cache, sending request", idParam)

	todos, err := SendTodosRequest(idParam, utils.TodoUrl)
	if err != nil {
		log.Printf("Error occured while getting todos: [%v}", err)
		return models.Todo{}, err
	}
	log.Printf("RETRIEVING TODOS ---> Id: [%v] retrieved data from request, storing in cache", idParam)

	todoCache[idParam] = todos
	return todos, nil
}

// When caching is disabled, this directly requests the posts data and does not store in cache
func HandlePostsRequestWhenCachingDisabled(idParam int) (models.Post, error) {
	log.Printf("RETRIEVING POSTS ---> CACHING DISABLED :: Id: [%v] sending request", idParam)

	posts, err := SendPostsRequest(idParam, utils.PostUrl)
	if err != nil {
		log.Printf("Error occured while getting posts: [%v}", err)
		return models.Post{}, err
	}
	log.Printf("RETRIEVING POSTS ---> Id: [%v] retrieved data from request", idParam)

	return posts, nil
}

// When caching is disabled, this directly requests the todos data and does not store in cache
func HandleTodosRequestWhenCachingDisabled(idParam int) (models.Todo, error) {
	log.Printf("RETRIEVING TODOS ---> CACHING DISABLED :: Id: [%v] sending request", idParam)

	todos, err := SendTodosRequest(idParam, utils.TodoUrl)
	if err != nil {
		log.Printf("Error occured while getting todos: [%v}", err)
		return models.Todo{}, err
	}
	log.Printf("RETRIEVING TODOS ---> Id: [%v] retrieved data from request", idParam)

	return todos, nil
}

// Sends and retrieves the new requests for posts
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

// Sends and retrives the new request for todos
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

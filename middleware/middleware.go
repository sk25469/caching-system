package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/sk25469/momoney-backend-assignment/models"
	"github.com/sk25469/momoney-backend-assignment/utils"
)

// PostCache stores the cached requests for all the posts
var cache map[int]interface{} = make(map[int]interface{})

var mut sync.Mutex

// Everytime a request goes through the middleware, it will check if the id
// already exist in their respective todos, if they exist, then it returns the
// cached response already present in the cache.
// If the id doesn't exist, it sends the requests for that id, gets the response and
// stores inside the cache
func InitMiddleWare(idParam int, requestType utils.Request, cachingEnabled bool) (interface{}, error) {
	mut.Lock()
	defer mut.Unlock()

	var logVal string
	if requestType == utils.Post {
		logVal = "POST"
	} else {
		logVal = "TODOS"
	}

	if !cachingEnabled {
		val, err := HandleRequestWhenCachingDisabled(idParam, requestType)
		if err != nil {
			log.Printf("Error occured while handling disabled caching for %v: [%v]", err, logVal)
			return models.Post{}, err
		}
		return val, nil
	}
	if _, ok := cache[idParam]; ok {
		log.Printf("RETRIEVING %v ---> Id: [%v] already present in cache, returning data from cache", logVal, idParam)
		return cache[idParam], nil
	} else {
		posts, err := HandleRequestForFirstTime(idParam, requestType)
		if err != nil {
			log.Printf("Error occured while getting %v: [%v}", err, logVal)
			return models.Post{}, err
		}
		return posts, nil
	}
}

// When a particular id is requested for the first time and caching is not disabled, this function runs
// for the todos, requests the data and stores in todoCache
func HandleRequestForFirstTime(idParam int, requestType utils.Request) (interface{}, error) {
	var logVal string
	if requestType == utils.Post {
		logVal = "POST"
	} else {
		logVal = "TODOS"
	}

	log.Printf("RETRIEVING %v ---> Id: [%v] not present in cache, sending request", logVal, idParam)

	result, err := SendRequest(idParam, requestType)
	if err != nil {
		log.Printf("Error occured while getting %v: [%v}", err, logVal)
		return models.Todo{}, err
	}
	log.Printf("RETRIEVING %v ---> Id: [%v] retrieved data from request, storing in cache", logVal, idParam)

	cache[idParam] = result
	return result, nil
}

// When caching is disabled, this directly requests the posts data and does not store in cache
func HandleRequestWhenCachingDisabled(idParam int, requestType utils.Request) (interface{}, error) {
	var logVal string
	if requestType == utils.Post {
		logVal = "POST"
	} else {
		logVal = "TODOS"
	}

	log.Printf("RETRIEVING %v ---> CACHING DISABLED :: Id: [%v] sending request", idParam, logVal)

	val, err := SendRequest(idParam, requestType)
	if err != nil {
		log.Printf("Error occured while getting %v: [%v}", err, logVal)
		return models.Post{}, err
	}
	log.Printf("RETRIEVING %v ---> Id: [%v] retrieved data from request", logVal, idParam)

	return val, nil
}

// Sends and retrieves the new requests for posts
func SendRequest(idParam int, requestType utils.Request) (interface{}, error) {

	var url, logVal string
	var returnVal interface{}
	if requestType == utils.Post {
		url = fmt.Sprintf(utils.PostUrl+"%d", idParam)
		returnVal = models.Post{}
		logVal = "POST"
	} else {
		url = fmt.Sprintf(utils.TodoUrl+"%d", idParam)
		returnVal = models.Todo{}
		logVal = "TODOS"
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch data from API: %v\n", err)
		return models.Post{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return models.Post{}, err
	}

	err = json.Unmarshal(body, &returnVal)
	if err != nil {
		fmt.Printf("Failed to unmarshal %v response body: %v\n", logVal, err)
		return models.Post{}, err
	}
	return returnVal, nil
}

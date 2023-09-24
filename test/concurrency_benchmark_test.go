package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func MakeRequest(requestId string) error {
	defer wg.Done()
	url := "http://localhost:3000/posts/" + requestId
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch data from API: " + err.Error())
		return err
	}

	// Make sure to close the response body before the program ends
	defer response.Body.Close()
	_, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read response body: ", err.Error())
		return err
	}

	// fmt.Println(string(body))
	return nil
}

func StartTest() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go MakeRequest(strconv.Itoa(i % 50))
	}
	wg.Wait()
}

func BenchmarkStartTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartTest()
	}
}

func main() {
	// StartTest() // You can run this for testing outside of benchmarking
	// Comment out the above line when running benchmarks
}

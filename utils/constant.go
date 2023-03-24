package utils

// Named enum to differentiate between requests to posts and todos
type Request int

const (
	Post Request = 0
	Todo Request = 1
)

const (
	PostUrl string = "https://jsonplaceholder.typicode.com/posts/"
	TodoUrl string = "https://jsonplaceholder.typicode.com/todos/"
)

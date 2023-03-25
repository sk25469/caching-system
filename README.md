# MoMoney Backend Assignment

This project provides an API for retrieving posts and todos from a database. It also includes caching functionality that can be toggled for individual routes or globally.

## Getting Started

### Prerequisites

* Go 1.16 or later

### Installation

To run the server locally -

* Clone this repository to your local machine.
* Run `go mod download` to download the project dependencies.
* Run `go run main.go` to start the server.

By default, the local server will run on port 3000.

### Usage

* Access the deployed API at <http://20.207.85.42:3000>.
* Use the available routes to interact with the API.

### API

* The API is hosted on Azure Container Instance and can be accessed using the base URL <http://20.207.85.42:3000>.
* The following routes are available:

| Endpoint        | Methods           | Description  |
| --------------- |:-----------------:| ------------:|
| `/posts/:id`      | GET | Retrieves a post with the given id |
| `/todos/:id`      | GET      |   Retrieves a todo with the given id |
| `/caching/posts=:flag` | GET      |    Toggles caching for all requests to posts |
| `/caching/todos=:flag`      | GET | Toggles caching for all requests to todos |
| `/caching=:flag` | GET      |    Toggles caching globally for all requests |

* `:flag` can be set to `true` or `false` to toggle caching on or off.
* `:id` is an integer

### Testing

To test the API using CLI, go to the test directory. There are multiple testing configurations available:

* To retrieve a post or todo, use the following command:

`go run server-test.go -id=<id-of-request> -path=<path-to-request>`

Replace `id-of-request` with the ID of the post or todo you want to retrieve and
`path-to-request` with either posts or todos.

* To toggle caching for a particular route, use the following command:

`go run server-test.go -caching=<flag> -path=<path-to-toggle>`

Replace `flag` with either `true` or `false`, depending on whether you want to enable or disable caching, and replace `path-to-toggle` with either `posts`, `todos`, or `all`, depending on which routes you want to toggle caching for.

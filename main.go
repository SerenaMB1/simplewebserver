package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writter http.ResponseWriter, request *http.Request) {

	})

}

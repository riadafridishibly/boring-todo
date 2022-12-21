package main

import (
	"log"
	"net/http"

	"github.com/riadafridishibly/svelte-todo/api"
)

func main() {
	todoApi, err := api.NewTodoAPI("test.db")
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8989", todoApi.Routes())
}

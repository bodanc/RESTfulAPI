package main

import (
	"fmt"
	"net/http"
	"os"
	"restapi/handlers"
)

func main() {
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("192.168.2.3:25000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

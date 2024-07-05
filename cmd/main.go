package main

import (
	"fmt"
	"net/http"
	"snipets/api"
)

const PORT = "8000"

func main() {
	http.HandleFunc("/snipets", api.Snipets)
	http.HandleFunc("/users", api.Users)
	http.HandleFunc("/auth", api.Auth)

	fmt.Printf("\nAPI server running at http://localhost:%s\n", PORT)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		panic(err)
	}
}

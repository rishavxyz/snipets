package main

import (
	"fmt"
	"net/http"
	"os"
	"snipets/api"

	"github.com/google/uuid"
)

const PORT = "8000"

func init() {
	mode := os.Getenv("MODE")

	// seed when not in production
	if mode != "production" {
		api.Snipets = append(api.Snipets,
			api.Snipet{
				Id:    uuid.New(),
				Code:  "const name = \"Rishav\"",
				Title: "This is JavaScript",
				Desc:  "",
				Lang:  "javascript",
				Theme: "nord",
			},
			api.Snipet{
				Id:    uuid.New(),
				Code:  ":(){:|:&};:",
				Title: "Fork bomb (Dos attack)",
				Desc:  "Do NOT run this code, it will crash your system",
				Lang:  "bash",
				Theme: "dracula",
			},
		)
	}

}

func main() {
	http.HandleFunc("/snipets", api.Main)
	http.HandleFunc("/auth", api.Auth)

	fmt.Printf("\nAPI server running at http://localhost:%s\n", PORT)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		panic(err)
	}
}

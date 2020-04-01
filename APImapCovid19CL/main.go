package main

import (
	"fmt"
	"net/http"
)

type application struct {
	server *http.ServeMux
}

func main() {
	app := application{server: http.NewServeMux()}
	app.Routes()

	fmt.Printf("Start Listen")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

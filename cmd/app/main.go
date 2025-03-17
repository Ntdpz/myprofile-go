package main

import (
	"fmt"
	"log"
	"net/http"

	"myprofile-go/internal/app/handler"

)

func main() {
	http.HandleFunc("/getMe", handler.GetMeHandler)
	http.HandleFunc("/", handler.StartHandler)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

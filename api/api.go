package api

import (
	"fmt"
	"log"
	"net/http"

	"01.alem.school/git/Azel/ascii-art-web-dockerize/handlers"
)

func Api() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomePage)
	mux.HandleFunc("/ascii-art", handlers.AsciiPage)

	fileServer := http.FileServer(http.Dir("./ui/style/"))
	mux.Handle("/style/", http.StripPrefix("/style", fileServer))

	fmt.Println("http://localhost:8080/")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

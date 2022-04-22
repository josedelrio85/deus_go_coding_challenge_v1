package main

import (
	"log"
	"net/http"

	deus "github.com/josedelrio85/deus_go_coding_challenge_v1/pkg"
)

func main() {
	log.Println("DEUS Go Coding Challenge v1 starting...")

	mux := http.NewServeMux()
	mux.Handle("/setevent", deus.SetEvent())
	mux.Handle("/getevent", deus.GetDistinctVisitors())

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

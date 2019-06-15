package main

import (
	"fmt"
	"log"

	/*"./handlers"
	"./server"*/

	"app/handlers"
	"app/server"

	"goji.io"
	"goji.io/pat"
)

var serverAddres = ":8000"
var user = "Diego"

func main() {
	fmt.Println("WE ARE LIVEEEE v2")
	mux := goji.NewMux()
	mux.HandleFunc(pat.Put("/organization-ms/folder"), handlers.Move)
	mux.HandleFunc(pat.Get("/organization-ms/logs/move"), handlers.LogMove)
	mux.HandleFunc(pat.Post("/organization-ms/folder"), handlers.CreateFolder)
	mux.HandleFunc(pat.Get("/organization-ms/logs/create"), handlers.LogCreateFolder)
	mux.HandleFunc(pat.Delete("/organization-ms/folder"), handlers.Delete)
	mux.HandleFunc(pat.Get("/organization-ms/logs/delete"), handlers.LogDelete)

	srv := server.New(mux, serverAddres)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server no funca :(  : %v", err)
	} else {

	}

}

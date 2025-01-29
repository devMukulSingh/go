package main

import (
	"fmt"
	"log"
	"net/http"
	"tutorial/internal/config"
)


func main()  {
	 cfg:= config.MustLoad()

	 router:= http.NewServeMux()

	 router.HandleFunc("GET /",func( w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello world"))
	 })

	 //setup server

	 server := http.Server{
		Addr: cfg.Address,
		Handler:router,
	 }

	 fmt.Printf("Server started at %s", cfg.Address);
	 err := server.ListenAndServe();
	if err!= nil{
		log.Fatal("Failed to start server")
	}

}


package main

import (
	"flag"
	"log"
	gohttp "net/http"

	supermart "github.com/alka/supermartchannel"
	"github.com/alka/supermartchannel/http"
	"github.com/alka/supermartchannel/store"
	"github.com/gorilla/mux"
)

var port = flag.String("port", "8080", "port to listen")

func main() {
	flag.Parse()
	router := mux.NewRouter()

	martStore := store.NewMartStore()
	itemStore := store.NewItemStore()
	defer itemStore.Close()
	defer martStore.Close()
	martService := supermart.NewSuperMartService(itemStore, martStore)
	martHandler := http.NewMartHandler(martService, martStore)
	martHandler.InstallRoutes(router)

	log.Println("starting http server, listening on port:", *port)
	if err := gohttp.ListenAndServe(":"+*port, router); err != nil {
		log.Fatalf("error in starting server: %v", err)
	}

}

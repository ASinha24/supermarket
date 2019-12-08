package main

import (
	"flag"
	"log"
	gohttp "net/http"

	"github.com/alka/supermart"
	"github.com/alka/supermart/http"
	"github.com/alka/supermart/store"
	"github.com/gorilla/mux"
)

var port = flag.String("port", "8080", "port to listen")

func main() {
	flag.Parse()
	router := mux.NewRouter()

	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)
	martHandler.InstallRoutes(router)

	log.Println("starting http server, listening on port:", *port)
	if err := gohttp.ListenAndServe(":"+*port, router); err != nil {
		log.Fatalf("error in starting server: %v", err)
	}

}

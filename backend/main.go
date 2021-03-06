package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"ajhaa.fi/elosimulator/routes"
	"github.com/julienschmidt/httprouter"
)

func apiRoot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Healthy!")
}

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	router := httprouter.New()
	router.GET("/", apiRoot)
	router.GET("/api", routes.Index)
	router.GET("/api/simulate", routes.Simulate)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", router))
}

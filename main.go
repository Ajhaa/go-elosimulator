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

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/simulate/:amount", routes.Simulate)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", router))
}

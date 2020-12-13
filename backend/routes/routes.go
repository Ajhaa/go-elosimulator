package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ajhaa.fi/elosimulator/simulator"
	"github.com/julienschmidt/httprouter"
)

// Index - get welcome message
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the elosimulator!\nRunning build %s\n", os.Getenv("COMMIT_SHA"))
}

// Simulate - Launch a simulation with N players
func Simulate(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	amount, err := strconv.Atoi(pm.ByName("amount"))
	if err != nil {
		fmt.Fprintf(w, "Not a number: %s", pm.ByName("amount"))
	}

	players := simulator.Simulate(amount)

	response, _ := json.Marshal(players)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
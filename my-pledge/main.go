package main

import (
	"github.com/gorilla/mux"
	"log"
	"martin.com/pleage/api/controllers"
	"martin.com/pleage/schedule"
	"net/http"
)

func main() {
	schedule.TestSchedule()
	r := mux.NewRouter()
	r.HandleFunc("/api/pledge", controllers.PledgeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

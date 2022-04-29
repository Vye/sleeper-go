package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sleep/{seconds}", SleepHandler)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if len(host) == 0 {
		host = "0.0.0.0"
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	log.Println("Listening on: ", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func SleepHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seconds, err := strconv.Atoi(vars["seconds"])
	if err != nil {
		fmt.Fprint(w, "Could not convert to int")
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Fprintf(w, "woke after %s seconds", vars["seconds"])
}


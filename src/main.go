package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	startTime = time.Now()
)

func main() {
	http.HandleFunc("/stat", getStatValue)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}

func getStatValue(writer http.ResponseWriter, request *http.Request) {
	statType := request.URL.Query().Get("stat")
	log.Printf("Getting stat %s", statType)
	if statType == "time" {
		t := time.Since(startTime).Round(time.Second)
		formattedTime := fmt.Sprintf("%s", t)
		fmt.Fprint(writer, formattedTime)
	} else if statType == "random" {
		fmt.Fprint(writer, rand.Int())
	} else {
		fmt.Fprint(writer, "Unknown stat, please check the http://localhost:8080")
	}
}

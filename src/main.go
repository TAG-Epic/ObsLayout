package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
		t := time.Now()
		formattedTime := fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
		fmt.Fprint(writer, formattedTime)
	} else {
		fmt.Fprint(writer, "Unknown stat, please check the http://localhost:8080")
	}
}

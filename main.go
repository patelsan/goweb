package main

import (
  "net/http"
  "log"
  "fmt"
  "time"
)

func tellTime(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, time.Now().Format("2006012150405"))
}

func main(){
  	http.HandleFunc("/", tellTime)
	fmt.Println("Listening on port 2020") 
	err := http.ListenAndServe(":2020", nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

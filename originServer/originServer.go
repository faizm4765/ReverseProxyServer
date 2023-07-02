package main

import (
	"fmt"
	"net/http"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am grooooot!")
	fmt.Println("Origin Server received request at ", time.Now())
	w.Header().Set("Content-Type", "text")
	responseText := "I am grooooot!"
	contentLength, err := w.Write([]byte(responseText))
	if err != nil {
		fmt.Println("Error in writing response body: ", err)
		return
	}

	fmt.Println("Content Length in Body", contentLength)
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

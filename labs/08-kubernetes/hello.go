package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Path:\t%s", r.URL.Path[1:])
	if err != nil {
		return
	}
	//for i,thing := range r.URL.Path{
	//	fmt.Fprintf(w, "[%d]:\t%s\n",i, thing)
	//}
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

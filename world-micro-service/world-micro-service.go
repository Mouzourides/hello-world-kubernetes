package main

import "fmt"
import "net/http"

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "world!")
}

func main() {
	http.HandleFunc("/world", world)
	http.ListenAndServe(":8081", nil)
}

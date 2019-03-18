package main

import ("fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello ")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://localhost:8081/world")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "Hello ", string(body))
}


func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello-world", helloWorld)
	http.ListenAndServe(":8080", nil)
}

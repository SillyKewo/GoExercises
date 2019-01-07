package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "text/plain")
	if err != nil {
		e := "could not read body: " + err.Error()
		http.Error(w, e, http.StatusInternalServerError)
		return
	}
	name := string(b)
	if name == "" {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye, web!")

}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Fprintf(w, "Hello %q\n", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/bye/", byeHandler)
	r.HandleFunc("/hello/{name}", nameHandler).Methods("GET")

	http.Handle("/", r)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

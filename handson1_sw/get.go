package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// try changing the value of this url
	res, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == 404 {
		fmt.Println(res.Status)
	}
	defer res.Body.Close()

	// if res.StatusCode == http.StatusOK {
	// 	body, _ := ioutil.ReadAll(res.Body)
	// 	str := string(body)

	// 	fmt.Println(str)
	// }

	doGet2()

}

func doGet() {
	req, err := http.NewRequest("GET", "https://golang.org", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	fmt.Println(res.Status)

}

func doGet2() {
	req, err := http.NewRequest("GET", "https://http-methods.appspot.com/Hungary/", nil)

	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	q := req.URL.Query()

	q.Set("v", "true")

	req.URL.RawQuery = q.Encode()


	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	
	
	if res.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		str := string(body)

		fmt.Println(str)
	}
}

func doPut() {
	s := strings.NewReader("Hello world!")

	req, err := http.NewRequest("PUT", "https://http-methods.appspot.com/kevinjonjensen/Message", s)

	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}


	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	fmt.Println(res.Status)


}

package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	from       string
	to         string
	serviceUrl string
)

func main() {
	from = os.Getenv("from")
	to = os.Getenv("to")
	serviceUrl = os.Getenv("service_url")
	fmt.Printf("env:  from target is: %s , to target is: %s , servicUrl is: %s \n", from, to, serviceUrl)

	fmt.Println("redirector is  up ... :D")
	err := http.ListenAndServe(":9091", http.HandlerFunc(redirect))
	if err != nil {
		fmt.Println("something is wrong :(")
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request host : " + req.URL.Host + " path : " + req.URL.Path)

	target := from + req.URL.Path

	if req.URL.Path == serviceUrl {
		target = to + req.URL.Path
		fmt.Println("redirect to target : " + target)
	} else {
		fmt.Println("redirect to target : " + target)
	}
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

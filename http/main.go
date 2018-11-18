package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, URL path %s!", r.URL.Path[1:])
	data, err := ioutil.ReadFile("site/index.html")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("File:\n", string(data))
	fmt.Fprintf(w, string(data))

}

func andreouhandler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("site/andreou/andreou.html")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("File:\n", string(data))
	fmt.Fprintf(w, string(data))

}

func moralishandler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("site/moralis/moralis.html")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("File:\n", string(data))
	fmt.Fprintf(w, string(data))

}

func vordonihandler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("site/vordoni/vordoni.html")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("File:\n", string(data))
	fmt.Fprintf(w, string(data))

}

func main() {
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/andreou/andreou.html", andreouhandler)
	http.HandleFunc("/moralis/moralis.html", moralishandler)
	http.HandleFunc("/vordoni/vordoni.html", vordonihandler)
	http.Handle("/site/", http.StripPrefix("/site/", http.FileServer(http.Dir("site"))))

	log.Fatal(http.ListenAndServe(":8086", nil))
}

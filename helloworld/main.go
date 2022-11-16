package main

import (
	"go/format"
	"net/http"
)



func main() {
	http.HandleFunc("/hello", sayHello)

	http.ListenAndServe(addr ":8080", handler nil)
}

func sayHello(w http.ResponseWriter,r *http.Request) {

    fmt.Printf(format "Request got...")
	w.Write([]byte("<h1>Hello</h1>"))
}
func getName(w.http.ResponseWriter) {

}
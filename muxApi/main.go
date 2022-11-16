package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type info struct {
	Name   string
	Email  string
	Native string
}

type infoinfo []info

func allinfo(w http.ResponseWriter, r *http.Request) {
	a := infoinfo{
		info{Name: "Ahin Vinith", Email: "ahinfencer07@gmail.com", Native: "Thottavaram"},
	}
	fmt.Println("Endpoint hit:All Infos endpoint")
	fmt.Fprintf(w, "%s\n", a)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/names", allinfo).Methods("GET")
	return r

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/a", allinfo)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: handleRequest(),
	}
	err := httpServer.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("unexpected: %v", err)
		os.Exit(1)
	}
}

package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
		router.HandleFunc("/", HelloWorld).Methods("GET")
    log.Fatal(http.ListenAndServe(":8765", router))
}

func HelloWorld(writer http.ResponseWriter, request *http.Request){
  writer.Write([]byte("Hello world"))
}

package main

import (
    "log"
    "net/http"
    "os"
    "os/exec"
)

type fortuneHandler struct {

}

func (f fortuneHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
  fortune, err := exec.Command("fortune").Output()
  if err != nil {
    writer.Write([]byte(err.Error()))
  }
  writer.Write(fortune)
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
	http.Handle("/fortune", fortuneHandler{})
  log.Fatal(http.ListenAndServe(":" + port, nil))
}

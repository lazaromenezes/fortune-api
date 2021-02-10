package main

import (
    "log"
    "net/http"
    "os"
    "os/exec"
)

func handleFortune(writer http.ResponseWriter, request *http.Request){
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
  http.HandleFunc("/fortune", handleFortune)
  log.Fatal(http.ListenAndServe(":" + port, nil))
}

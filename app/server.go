package main

import (
	"fmt"
	"net/http"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  homePage := http.FileServer(http.Dir("./src"))
  http.Handle("/", homePage)

  fmt.Println("Starting server at localhost:5099...")
  startServer := http.ListenAndServe(":5099", nil)

  check(startServer)
}

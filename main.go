package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":8080", nil)
  fmt.Println("Serving in container port 8080..")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello from frontend - staging")
}

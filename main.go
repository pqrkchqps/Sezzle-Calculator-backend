//Tutorial used
//https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import(
  "fmt"
  "log"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, Welcome to the Homepage")
  fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
  handleRequests()
}

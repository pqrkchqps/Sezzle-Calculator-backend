//Tutorial used
//https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import(
  "fmt"
  "log"
  "net/http"
  "strings"
  "strconv"
)

type Calculation struct {
  op string
  x float64
  y float64
  calc float64
}

func addition(w http.ResponseWriter, r *http.Request){
  u := strings.Split(r.URL.Path, "/")
  x, err := strconv.ParseFloat(u[2], 64)
  if err != nil {
    panic(err)
  }
  y, err := strconv.ParseFloat(u[3], 64)
  if err != nil {
    panic(err)
  }
  calc := x + y
  returnVal := Calculation{"+", x, y, calc}
  fmt.Fprintf(w, "Hello, Welcome %s", returnVal)
  fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
  http.HandleFunc("/+/", addition)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
  handleRequests()
}

//Tutorial used
//https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import(
  "fmt"
  "log"
  "net/http"
  "strings"
  "strconv"
  "encoding/json"
)

type Calculation struct {
  Op string
  X float64
  Y float64
  Calc float64
}

func addition(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint hit: addition")
  u := strings.Split(r.URL.Path, "/")
  x, err := strconv.ParseFloat(u[3], 64)
  if err != nil {
    panic(err)
  }
  y, err := strconv.ParseFloat(u[4], 64)
  if err != nil {
    panic(err)
  }
  calc := x + y
  returnVal := &Calculation{Op: "+", X: x, Y: y, Calc: calc}
  fmt.Println(returnVal)
  if err := json.NewEncoder(w).Encode(returnVal); err != nil {
    panic(err)
  }
}

func handleRequests() {
  http.HandleFunc("/calculate/+/", addition)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
  handleRequests()
}

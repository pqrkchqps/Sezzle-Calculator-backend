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
  "os"
)

type Calculation struct {
  Op string
  X float64
  Y float64
  Calc float64
}

var calculations []Calculation

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
  calculations = append(calculations, *returnVal)
}

func subtraction(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint hit: subtraction")
  u := strings.Split(r.URL.Path, "/")
  x, err := strconv.ParseFloat(u[3], 64)
  if err != nil {
    panic(err)
  }
  y, err := strconv.ParseFloat(u[4], 64)
  if err != nil {
    panic(err)
  }
  calc := x - y
  returnVal := &Calculation{Op: "-", X: x, Y: y, Calc: calc}
  fmt.Println(returnVal)
  if err := json.NewEncoder(w).Encode(returnVal); err != nil {
    panic(err)
  }
  calculations = append(calculations, *returnVal)
}

func multiplication(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint hit: multiplication")
  u := strings.Split(r.URL.Path, "/")
  x, err := strconv.ParseFloat(u[3], 64)
  if err != nil {
    panic(err)
  }
  y, err := strconv.ParseFloat(u[4], 64)
  if err != nil {
    panic(err)
  }
  calc := x * y
  returnVal := &Calculation{Op: "*", X: x, Y: y, Calc: calc}
  fmt.Println(returnVal)
  if err := json.NewEncoder(w).Encode(returnVal); err != nil {
    panic(err)
  }
  calculations = append(calculations, *returnVal)
}

func division(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint hit: division")
  u := strings.Split(r.URL.Path, "/")
  x, err := strconv.ParseFloat(u[3], 64)
  if err != nil {
    panic(err)
  }
  y, err := strconv.ParseFloat(u[4], 64)
  if err != nil {
    panic(err)
  }
  calc := x / y
  returnVal := &Calculation{Op: "d", X: x, Y: y, Calc: calc}
  fmt.Println(returnVal)
  if err := json.NewEncoder(w).Encode(returnVal); err != nil {
    panic(err)
  }
  calculations = append(calculations, *returnVal)
}

func getCalculations(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint hit: calulations")
  idx := len(calculations)-10
  if (idx < 0){
    idx = 0
  }
  calculations = calculations[idx:]
  fmt.Println(calculations)
  (w).Header().Set("Access-Control-Allow-Origin", "https://arcane-sierra-66995.herokuapp.com")

  if err := json.NewEncoder(w).Encode(calculations); err != nil {
    panic(err)
  }
}

func resetCalculations(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint hit: calulations/reset")
  calculations = calculations[:0]
}

func handleRequests() {
  http.HandleFunc("/calculate/+/", addition)
  http.HandleFunc("/calculate/-/", subtraction)
  http.HandleFunc("/calculate/*/", multiplication)
  http.HandleFunc("/calculate/d/", division)
  http.HandleFunc("/calculations/", getCalculations)
  http.HandleFunc("/calculations/reset", resetCalculations)


  port := os.Getenv("PORT")
  if port == "" {
    port = ":8081"
  } else {
    port = ":" + port
  }
  log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
  handleRequests()
}

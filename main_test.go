//Tutorial Used
//https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6

package main
import (
    "testing"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "fmt"
)


var Router *mux.Router = mux.NewRouter()

func TestAddition(t *testing.T) {
  res, _ := http.Get("http://localhost:8081/calculations/reset")

  res, _ = http.Get("http://localhost:8081/calculate/+/2/3")
  fmt.Println(res)
  checkResponseCode(t, http.StatusOK, res.StatusCode)
  var calc Calculation;
	json.NewDecoder(res.Body).Decode(&calc)
  if calc.Op != "+" {
      t.Errorf("Expected +. Got %s", calc.Op)
  }
  if calc.X != 2 {
      t.Errorf("Expected 2. Got %f", calc.X)
  }
  if calc.Y != 3 {
      t.Errorf("Expected 3. Got %f", calc.Y)
  }
  if calc.Calc != 5 {
      t.Errorf("Expected 5. Got %f", calc.Calc)
  }
}


func TestSubtraction(t *testing.T) {
  res, _ := http.Get("http://localhost:8081/calculate/-/2/3")
  fmt.Println(res)
  checkResponseCode(t, http.StatusOK, res.StatusCode)
  var calc Calculation;
	json.NewDecoder(res.Body).Decode(&calc)
  if calc.Op != "-" {
      t.Errorf("Expected -. Got %s", calc.Op)
  }
  if calc.X != 2 {
      t.Errorf("Expected 2. Got %f", calc.X)
  }
  if calc.Y != 3 {
      t.Errorf("Expected 3. Got %f", calc.Y)
  }
  if calc.Calc != -1 {
      t.Errorf("Expected -1. Got %f", calc.Calc)
  }
}

func TestMultiplcation(t *testing.T) {
  res, _ := http.Get("http://localhost:8081/calculate/*/2/3")
  fmt.Println(res)
  checkResponseCode(t, http.StatusOK, res.StatusCode)
  var calc Calculation;
	json.NewDecoder(res.Body).Decode(&calc)
  if calc.Op != "*" {
      t.Errorf("Expected -. Got %s", calc.Op)
  }
  if calc.X != 2 {
      t.Errorf("Expected 2. Got %f", calc.X)
  }
  if calc.Y != 3 {
      t.Errorf("Expected 3. Got %f", calc.Y)
  }
  if calc.Calc != 6 {
      t.Errorf("Expected 6. Got %f", calc.Calc)
  }
}

func TestDivision(t *testing.T) {
  res, _ := http.Get("http://localhost:8081/calculate/d/6/3")
  fmt.Println(res)
  checkResponseCode(t, http.StatusOK, res.StatusCode)
  var calc Calculation;
	json.NewDecoder(res.Body).Decode(&calc)
  if calc.Op != "d" {
      t.Errorf("Expected -. Got %s", calc.Op)
  }
  if calc.X != 6 {
      t.Errorf("Expected 6. Got %f", calc.X)
  }
  if calc.Y != 3 {
      t.Errorf("Expected 3. Got %f", calc.Y)
  }
  if calc.Calc != 2 {
      t.Errorf("Expected 2. Got %f", calc.Calc)
  }
}

func TestGetCalculations(t *testing.T) {
  res, _ := http.Get("http://localhost:8081/calculations")
  var calcs []Calculation;
  json.NewDecoder(res.Body).Decode(&calcs)
  if calcs[0].Op != "+" {
      t.Errorf("Expected +. Got %s", calcs[0].Op)
  }
  if calcs[1].Op != "-" {
      t.Errorf("Expected -. Got %s", calcs[1].Op)
  }
  if calcs[2].Op != "*" {
      t.Errorf("Expected *. Got %s", calcs[2].Op)
  }
  if calcs[3].Op != "d" {
      t.Errorf("Expected d. Got %s", calcs[3].Op)
  }
  res, _ = http.Get("http://localhost:8081/calculate/+/2/3")
  res, _ = http.Get("http://localhost:8081/calculate/-/2/3")
  res, _ = http.Get("http://localhost:8081/calculate/*/2/3")
  res, _ = http.Get("http://localhost:8081/calculate/d/6/3")

  res, _ = http.Get("http://localhost:8081/calculate/+/2/3")
  res, _ = http.Get("http://localhost:8081/calculate/-/2/3")
  res, _ = http.Get("http://localhost:8081/calculate/*/2/3")
  res, _ = http.Get("http://localhost:8081/calculate/d/6/3")

  res, _ = http.Get("http://localhost:8081/calculations")
  json.NewDecoder(res.Body).Decode(&calcs)
  if calcs[0].Op != "*" {
      t.Errorf("Expected *. Got %s", calcs[2].Op)
  }
  if calcs[1].Op != "d" {
      t.Errorf("Expected d. Got %s", calcs[3].Op)
  }
  if calcs[2].Op != "+" {
      t.Errorf("Expected +. Got %s", calcs[0].Op)
  }
  if calcs[3].Op != "-" {
      t.Errorf("Expected -. Got %s", calcs[1].Op)
  }
  if calcs[4].Op != "*" {
      t.Errorf("Expected *. Got %s", calcs[2].Op)
  }
  if calcs[5].Op != "d" {
      t.Errorf("Expected d. Got %s", calcs[3].Op)
  }
  if calcs[6].Op != "+" {
      t.Errorf("Expected +. Got %s", calcs[0].Op)
  }
  if calcs[7].Op != "-" {
      t.Errorf("Expected -. Got %s", calcs[1].Op)
  }
  if calcs[8].Op != "*" {
      t.Errorf("Expected *. Got %s", calcs[2].Op)
  }
  if calcs[9].Op != "d" {
      t.Errorf("Expected d. Got %s", calcs[3].Op)
  }
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

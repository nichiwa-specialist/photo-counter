package picture_counter

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
}

func init() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/count", count)
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, " + r.FormValue("name"))
}

func count(w http.ResponseWriter, r *http.Request) {
    marks := [...]Mark { Mark{123, 456, 15}, Mark{321, 654, 11} }
    data, _ := json.Marshal(marks)
    fmt.Fprint(w, string(data))
}

type Mark struct {
    X int `json:"x"`
    Y int `json:"y"`
    Size int `json:"size"`
}

package picture_counter

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, " + r.FormValue("name"))
}

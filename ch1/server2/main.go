// Server2 is a minimal "echo" and counter srever
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex
var count int

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)

    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count ++
    mu.Unlock()
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}

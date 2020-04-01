package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
fs := http.FileServer(http.Dir("."))
http.Handle("/files/", http.StripPrefix("/files", fs))
http.HandleFunc("/dump", func(w http.ResponseWriter, r *http.Request) {
  dump, _ := httputil.DumpRequest(r, true)
fmt.Fprintf(w, "%s", string(dump))
})
log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import "github.com/gorilla/mux"
import  "log"
import  "net/http"
import  "time"
import  "io"


func handler(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello, World!\n")
}

// Route declararation
func router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/", handler)
  return r
}

func main() {
  router := router()
  srv := &http.Server{
    Handler: router,
    Addr: "127.0.0.1:9100",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  log.Fatal(srv.ListenAndServe())
}

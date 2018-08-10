package main

import (
  	"github.com/gorilla/mux"
    "net/http"
    "log"
)

func main()  {
  router := mux.NewRouter()
  router.HandleFunc("/", dolarHandler)

  http.Handle("/", router)
  log.Fatal(http.ListenAndServe(":8080", router))

}

func dolarHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("I am alive !\n"))
}

package main

import (
  	"github.com/gorilla/mux"
    "net/http"
    "log"
    "strconv"

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
  asgari := getAsgari()
  dolar := getDolar()
  asgariDolar:= asgari / dolar
  a:=strconv.FormatFloat(asgari,'f', -1, 64)
  d:=strconv.FormatFloat(dolar,'f', -1, 64)
  ad:=strconv.FormatFloat(asgariDolar,'f', -1, 64)
  w.Write([]byte("asgari ucret : "+a+"\n"))
  w.Write([]byte("dolar : "+d+"\n"))
  w.Write([]byte("asgari ucret kac dolar : "+ad+"\n"))
}

func getAsgari() (asgari float64)  {
  asgari=1725.08
  return
}

func getDolar() (dolar float64) {
  dolar = 6.46
  return
}

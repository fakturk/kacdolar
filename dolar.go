package main

import (
  	"github.com/gorilla/mux"
    "net/http"
    "log"
    "strconv"
    "encoding/json"
	  "io/ioutil"
    "fmt"

)

type Doviz []struct {
	Selling    float64 `json:"selling"`
	UpdateDate int     `json:"update_date"`
	Currency   int     `json:"currency"`
	Buying     float64 `json:"buying"`
	ChangeRate float64 `json:"change_rate"`
	Name       string  `json:"name"`
	FullName   string  `json:"full_name"`
	Code       string  `json:"code"`
}

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

    res, _ := http.Get("https://www.doviz.com/api/v1/currencies/all/latest")

  	temp, _ := ioutil.ReadAll(res.Body)

  	var doviz Doviz
  	err := json.Unmarshal(temp, &doviz)
  	if err != nil {
  		fmt.Println("There was an error:", err)
  	}
  dolar = doviz[0].Selling

  return
}

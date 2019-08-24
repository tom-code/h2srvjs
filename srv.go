package main

import (
  "io/ioutil"
  "log"
  "net/http"

  "github.com/robertkrimen/otto"
  "golang.org/x/net/http2"
  "golang.org/x/net/http2/h2c"
)

var (
  vm *otto.Otto
)

func server() {

  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    uri := r.RequestURI
    log.Println(uri)

    body, _ := ioutil.ReadAll(r.Body)
    r.Body.Close()
    
    _, err := vm.Call("handler", nil, r.URL, string(body), w)
    if err != nil {
      log.Println(err)
    }
  })

  h2s := &http2.Server{}
  h1s := &http.Server{
    Addr:    ":8082",
    Handler: h2c.NewHandler(handler, h2s),
  }

  h1s.ListenAndServe()
}

func main() {
  vm = otto.New()
  script, err := ioutil.ReadFile("script.js")
  if err != nil {
    panic(err)
  }

  _, err = vm.Run(script)
  if err != nil {
    panic(err)
  }

  server()
}

package main

import (
  "io/ioutil"
  "log"
  "net/http"
  "sync"
  "flag"

  "github.com/robertkrimen/otto"
  "golang.org/x/net/http2"
  "golang.org/x/net/http2/h2c"
)

var (
  vm *otto.Otto
  mutex = &sync.Mutex{}
)

func server(bindaddr string) {

  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    uri := r.RequestURI
    log.Println(uri)

    body, _ := ioutil.ReadAll(r.Body)
    r.Body.Close()

    mutex.Lock()
    _, err := vm.Call("handler", nil, r.URL, string(body), w)
    mutex.Unlock()
    if err != nil {
      log.Println(err)
    }
  })

  h2s := &http2.Server{}
  h1s := &http.Server{
    Addr:    bindaddr,
    Handler: h2c.NewHandler(handler, h2s),
  }

  log.Printf("server binding to %s\n", bindaddr)
  err := h1s.ListenAndServe()
  if err != nil {
    panic(err)
  }
}

func main() {
  bindptr := flag.String("bind", ":8082", "bind address")
  flag.Parse()

  vm = otto.New()
  script, err := ioutil.ReadFile("script.js")
  if err != nil {
    panic(err)
  }

  _, err = vm.Run(script)
  if err != nil {
    panic(err)
  }

  server(*bindptr)
}

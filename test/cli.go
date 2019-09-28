package main

import (
  "crypto/tls"
  "fmt"
  "io/ioutil"
  "net"
  "net/http"
  "strings"

  "golang.org/x/net/http2"
)

func createHTTP2Client() *http.Client {
  client := &http.Client{}
  client.Transport = &http2.Transport{
    AllowHTTP: true,
    DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
      return net.Dial(netw, addr)
    }}
  return client
}

func main() {
  c := createHTTP2Client()
  resp, err := c.Get("http://localhost:8082/tra/da/1")
  if err != nil {
    panic(err)
  }
  b, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(b))

  req, _ := http.NewRequest("POST", "http://localhost:8082/tra/da/2", strings.NewReader("{\"aa\":\"dd\"}"))
  resp, err = c.Do(req)
  
  b, _ = ioutil.ReadAll(resp.Body)
  fmt.Println(string(b))
}

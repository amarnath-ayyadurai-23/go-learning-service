package main

import (
  "fmt"
  "net/http"
  "io/ioutil"

)

func TestGET(url string) {

  //url := uri
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}

func main(){
	TestGET("http://127.0.0.1:3000/hello")
  TestGET("http://127.0.0.1:3000/")
  TestGET("http://127.0.0.1:3000/api/v1")
}
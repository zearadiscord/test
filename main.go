package main

import (
  "fmt"
  "net/http"
  "sync"
)
var wg sync.WaitGroup

func sub() {
  defer wg.Done()
  a,_ := http.Get("http://ddosfilter.net/")
  fmt.Println("test",a.Status)
}

func main() {
  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go sub()
    defer sub()
  }
}

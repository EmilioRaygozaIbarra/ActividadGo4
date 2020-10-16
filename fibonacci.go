package main

import "fmt"

func fibonacci(n int) int{
  if n==1 || n==0{
    return n
  }else{
    return fibonacci(n-1)+fibonacci(n-2)
  }
}

func main() {
  var n int
  fmt.Scanln(&n)
  for i:=0;i<n;i++{
    fmt.Println(fibonacci(i))
  }
}
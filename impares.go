package main

import "fmt"

func generadorImpares() func() int{
  i:=int(1)
  return func() int{
    x := i
    i += 2
    return x
  }
}

func main() {
  impar:=generadorImpares()
  fmt.Println(impar())
  fmt.Println(impar())
  fmt.Println(impar())
  fmt.Println(impar())
  fmt.Println(impar())
  fmt.Println(impar())
}
package main

import "fmt"

func numeroMasGrande(args ...int) int{
  aux:=0
  for _, v := range(args){
    if aux<v{
      aux=v
    }
  }
  return aux
}

func main() {
  n:=numeroMasGrande(2,8,4,3,5,20,4,19)
  fmt.Println(n)
}
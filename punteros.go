package main

import "fmt"

func intercambia(a, b *int){
  aux:=0
  aux=*b
  *b=*a
  *a=aux
}

func main() {
  var a,b int
  fmt.Println("a: ")
  fmt.Scan(&a)
  fmt.Println("b: ")
  fmt.Scan(&b)
  intercambia(&a,&b)
  fmt.Println("a: ",a)
  fmt.Println("b: ",b)

}
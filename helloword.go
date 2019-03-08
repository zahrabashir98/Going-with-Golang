package main

import "fmt"

func main() {

  x := "Hello World"

  x = "first "
  fmt.Println(x)
  x = x + "second"
  fmt.Println(x)

  x  = "hello"
  var y string = "world"
  fmt.Println(x == y)

  fmt.Println("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)


}
package main

import "fmt"

func main() {

  // A slice is a segment of an array. Like arrays slices are indexable and have a length
  var x []float64
  x = make([]float64, 5)
  fmt.Println(x)

  // add limitation for array
  y := make([]float64, 5, 10)
  fmt.Println(y)

  arr := [5]float64{1,2,3,4,5}
  z := arr[0:5]
  fmt.Println(arr)
  fmt.Println(z)
}

package main

import "fmt"

func main() {

    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
      }
    min := x[0] ;
    
    for i := 0; i < len(x); i++ {
        if i==len(x)-1{
            break}
        if x[i]>x[i+1]{
            min = x[i+1]
        }
    }
    fmt.Println(min)

      


}

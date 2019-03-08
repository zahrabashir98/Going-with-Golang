package main

import "fmt"

func main() {

    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
    slice3 := []int{5,6,7}
    slice4 := make([]int, 5)
    copy(slice4, slice3)
    fmt.Println(slice3, slice4)

}

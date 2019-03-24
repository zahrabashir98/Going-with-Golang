package main

import "fmt"

func main() {

    // Runtime Error: map should be intialized before
    // var x map[string]int
    // x["key"] = 10
    // fmt.Println(x)

    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])
    fmt.Println(x)
    delete(x, "key")
    fmt.Println(x)
    // returns 0, false for undefined keys (false is flag to know wether it was successful or not)
    name, ok := x["Un"]
    fmt.Println(name, ok)

    if name, ok := x["s"]; ok {
        fmt.Println(name, ok)
      }
      
      


}

package main

import "fmt"

func main() {

    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])
    fmt.Println(x)
    delete(x, "key")
    fmt.Println(x)
    // returns 0, false for undefined keys (false is flag to know wether it was successful or not)
    name, ok := x["Un"]
    fmt.Println(name, ok)

    // ?
    if name, ok := x["s"]; ok {
        fmt.Println(name, ok)
      }

    // map of maps
    elements := map[string]map[string]string{
    "H": map[string]string{
        "name":"Hydrogen",
        "state":"gas",
    },
    "He": map[string]string{
        "name":"Helium",
        "state":"gas",
    }}
    fmt.Println(elements["H"]["name"])

    v := [6]string{"a","b","c","d","e","f"}
    fmt.Println(v[2:5])
    
      


}

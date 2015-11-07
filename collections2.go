package main

import "fmt"

func main() {

    // explicit length
    slice := make([]string, 3)
    slice[0] = "assigning"
    slice[1] = "some"
    slice[2] = "variables"

    fmt.Printf("%q\n", slice)
    fmt.Printf("Length is %d", len(slice))
}
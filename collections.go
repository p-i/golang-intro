package main

import "fmt"

func main() {

    // implicit length
    array := [...]string{"hello", "world!", "let's", "put", "some", "more", "words", "here"}
    fmt.Printf("%q\n", array)
    fmt.Println(array[0], array[1])

    slice := array[0:1]
    fmt.Printf("%q\n", slice)

    slice = array[0:2]
    fmt.Printf("%q\n", slice)

    slice = append(slice, "!!!")
    fmt.Printf("%q\n", slice)
}

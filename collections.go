package main

import "fmt"

func main() {

    a := [...]string{"hello", "world!", "let's", "put", "some", "more", "words", "here"}  //implicit length

    fmt.Printf("%q\n", a)
    fmt.Println(a[0], a[1])

    slice := a[0:1]
    fmt.Printf("%q\n", slice)

    slice = a[0:2]
    fmt.Printf("%q\n", slice)

    slice = append(slice, "!!!")
    fmt.Printf("%q\n", slice)
}

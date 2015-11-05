package main

import "fmt"

func main() {
    action := func() {
        name := "The Ultimate Answer"
        value := 42
        fmt.Printf("%s: %d", name, value)
    }
    action()
}
package main

import "fmt"

func main() {

    action := func(value interface{}) {
        fmt.Printf("The Ultimate Answer: %v of type %T \n", value, value)
    }

    action(42)
    action(42.00)
    action("42")
    action(true)
}
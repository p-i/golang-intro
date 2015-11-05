package main

import (
    "math/rand"
    "fmt"
    "time"
)

func main() {
    message := "Random Value"
    rand.Seed(time.Now().UnixNano())
    fmt.Printf("%s: %d", message, rand.Int())
}
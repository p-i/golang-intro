package main

import "fmt"

func main() {
    cities := map[string]int{
        "New York":    8336697,
        "Los Angeles": 3857799,
        "Chicago":     2714856,
    }
    for key, value := range cities {
        fmt.Printf("%s has %d inhabitants\n", key, value)
    }
    fmt.Println()
    for _, value := range cities {
        fmt.Printf("some city has %d inhabitants\n", value)
    }
}
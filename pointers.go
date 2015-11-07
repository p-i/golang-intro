package main

import "fmt"

type User struct {
    Name, Location string
}

func (u *User) Greetings() string {
    return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
    *User
}

func NewPlayer(name, location string) *Player {
    return &Player{
        User:   &User{name, location},
    }
}

func main() {
    fmt.Println(NewPlayer("Matt", "LA").Greetings())
}
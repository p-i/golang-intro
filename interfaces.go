package main

import "fmt"

type User struct {
	FirstName, LastName string
}

type Namer interface {
	Name() string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.LastName, u.FirstName)
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"Guy", "Family"}
	fmt.Println(Greet(u))
}
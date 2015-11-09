package main
import "fmt"

type User struct {
    Name, Team string
}

func (u *User) Greetings() string {
    return fmt.Sprintf("Hi %s from team %s", u.Name, u.Team)
}

type Employee struct {
    *User
}

func NewEmployee(name, team string) *Employee {
    return &Employee{
        User:   &User{name, team},
    }
}

func main() {
    fmt.Println(NewEmployee("Matt", "BAs").Greetings())
}
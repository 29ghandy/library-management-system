package entities

import "fmt"

type User struct {
	id       int
	Name     string
	Password string
	Email    string
}

func (u User) Print() {
	fmt.Println("id:", u.id)
	fmt.Println("Name:", u.Name)
}

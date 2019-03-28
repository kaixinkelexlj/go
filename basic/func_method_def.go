package basic

import "fmt"

type User struct {
	Id   int
	Name string
}

func (u User) GetUserNameAndPrint() string {
	fmt.Printf("hello %v", u.Name)
	return u.Name
}

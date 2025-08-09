package users

import "slices"

type User interface {
	Age() int
}

type Employee struct {
	age int
}

type Customer struct {
	age int
}

func (u Employee) Age() int {
	return u.age
}

func (u Customer) Age() int {
	return u.age
}

func Elder(u ...User) int {
	elder := slices.MaxFunc(u, func(u1 User, u2 User) int {
		return u1.Age() - u2.Age()
	})
	return elder.Age()
}

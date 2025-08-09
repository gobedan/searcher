package users2

import (
	"reflect"
	"slices"
)

type User interface {
	Employee | Customer
}

type Employee struct {
	Age int
}

type Customer struct {
	Age int
}

func Elder[U User](u ...U) U {
	return slices.MaxFunc(u, func(u1 U, u2 U) int {
		user1 := reflect.ValueOf(u1)
		user2 := reflect.ValueOf(u2)
		return int(user1.FieldByName("Age").Int() - user2.FieldByName("Age").Int())
	})
}

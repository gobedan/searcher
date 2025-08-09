package users2

// "reflect"

type Employee struct {
	Age int
}

type Customer struct {
	Age int
}

func Elder(users ...any) any {
	maxAge := 0
	var elder any
	for _, user := range users {
		switch user := user.(type) {
		case Employee:
			if user.Age > maxAge {
				elder = user
				maxAge = user.Age
			}
		case Customer:
			if user.Age > maxAge {
				elder = user
				maxAge = user.Age
			}
		}
	}

	return elder
}

/* type User interface {
	Employee | Customer
} */
/* func Elder[U User](u ...U) U {
	return slices.MaxFunc(u, func(u1 U, u2 U) int {
		user1 := reflect.ValueOf(u1)
		user2 := reflect.ValueOf(u2)
		return int(user1.FieldByName("Age").Int() - user2.FieldByName("Age").Int())
	})
} */

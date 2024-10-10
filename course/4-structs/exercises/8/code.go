package main

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	limit := 100
	if membershipType == "premium" {
		limit = 1000
	}
	user := User{
		Membership: Membership{membershipType, limit},
		Name:       name,
	}
	return user
}

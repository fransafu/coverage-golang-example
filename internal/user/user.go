// Package user describe user entity attributes and manage user storage
package user

type User struct {
	ID        uint64
	FirstName string
	LastName  string
	Age       uint16
}

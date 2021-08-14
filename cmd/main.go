package main

import (
	"fmt"

	user "github.com/fransafu/coverage-golang-example/internal/user"
)

func main() {
	var userStorage user.UserStorage

	user1 := user.User{}
	user1.ID = 1
	user1.FirstName = "Francisco"
	user1.LastName = "Sanchez"
	user1.Age = 99

	user2 := user.User{}
	user2.ID = 2
	user2.FirstName = "Scott"
	user2.LastName = "Sanchez"
	user2.Age = 68

	userStorage.AddUser(user1)
	userStorage.AddUser(user2)

	fmt.Println("Search user by ID 1: ", userStorage.FindUserByID(1))
}

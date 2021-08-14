package user_test

import (
	"testing"

	user "github.com/fransafu/coverage-golang-example/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestUserStorage_SaveUser(t *testing.T) {
	var userStorage user.UserStorage

	user1 := user.User{}
	user1.ID = 1
	user1.FirstName = "Francisco"
	user1.LastName = "Sanchez"
	user1.Age = 99

	userStorage.AddUser(user1)

	assert.Equal(t, 1, userStorage.Count())
}

func TestUserStorage_SearchUser(t *testing.T) {
	var userStorage user.UserStorage

	user1 := user.User{}
	user1.ID = 1
	user1.FirstName = "Francisco"
	user1.LastName = "Sanchez"
	user1.Age = 99

	userStorage.AddUser(user1)

	assert.Equal(t, user1, userStorage.FindUserByID(user1.ID))
}

func TestUserStorage_EmptySearchUser(t *testing.T) {
	var userStorage user.UserStorage

	assert.Equal(t, 0, userStorage.Count())
	assert.Equal(t, user.User{}, userStorage.FindUserByID(1))
}

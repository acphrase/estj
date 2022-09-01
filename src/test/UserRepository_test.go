package test

import (
	"estj/src/dataaccesslayer/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_GetAllUser(t *testing.T) {
	// Start app.
	StartingTest()

	// Given
	userRepository := repository.GetUserRepository()

	// When
	users, _ := userRepository.GetAllUser()

	// Then
	assert.True(t, true, len(*users) > 0)

	// End app.
	EndTest()
}

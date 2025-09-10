package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	user, err := NewUser("John Doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.True(t, user.ValidatePassword("password123"))
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Jane Doe", "jane.doe@example.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, user.Password, "123456", "Password should be hashed")
}

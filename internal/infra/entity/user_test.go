package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser("daniel","osmar","daniel@gmail.com","123456","")
	assert.NotNil(t, user.ID)
	assert.NotEqual(t, "123456", user.Password)
	assert.Equal(t, "daniel@gmail.com",user.Email)
	assert.Equal(t, "daniel", user.Firstname)
	assert.Equal(t, "osmar", user.Lastname)
}
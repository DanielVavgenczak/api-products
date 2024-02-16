package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCategory(t *testing.T) {
	category := NewCategory("eletronic","f302fb78-4de5-495b-a047-0fbcd7089b66")
	assert.NotEmpty(t, category.ID)
	assert.Equal(t, "eletronic", category.Title)
}
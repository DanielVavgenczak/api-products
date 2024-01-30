package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCategory(t *testing.T) {
	category := NewCategory("eletronic")
	assert.NotEmpty(t, category.ID)
	assert.Equal(t, "eletronic", category.Title)
}
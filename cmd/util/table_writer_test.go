package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	t.Run("should return user information on valid request", func(t *testing.T) {
		var table = CreateTable()
		assert.NotNil(t, table)
	})
}

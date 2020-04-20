package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumberString(t *testing.T) {
	testMap := []int{5, 30, 93, 500000, 6}
	for _, length := range testMap {
		randomNumber := RandomNumberString(length)
		assert.Equal(t, length, len(randomNumber))
	}
}

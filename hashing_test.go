package util

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	var src = rand.NewSource(time.Now().UnixNano())
	rString := RandomString(32, src)
	assert.NotEmpty(t, rString)
	t.Log(rString)
}

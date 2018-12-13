package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testType struct {
	Duration int64
}

func TestFillStruct(t *testing.T) {
	testStruct := testType{}

	valuesToFill := map[string]string{
		"Duration": "1000",
	}

	err := FillStruct(&testStruct, valuesToFill)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, testStruct.Duration)
		assert.Equal(t, int64(1000), testStruct.Duration)
	}
}

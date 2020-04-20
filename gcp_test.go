package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMetadataServerToken(t *testing.T) {
	token, err := GetMetadataserverToken("accounts.paysenger.co")
	if assert.NoError(t, err) {
		t.Log(token)
	}
}

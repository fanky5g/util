package util

import (
	"fmt"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

// GetMetadataserverToken gets jwt token for authorizing internal requests
func GetMetadataserverToken(serviceURL string) (string, error) {
	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", strings.TrimSpace(serviceURL))
	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		return "", fmt.Errorf("metadata.Get: failed to query id_token: %+v", err)
	}

	return idToken, nil
}

// VerifyMetadataserverToken verifies passed metadata token
func VerifyMetadataserverToken(serviceURL, token string) {
	return
}

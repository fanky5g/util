package util

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCheckRecaptchaToken(t *testing.T) {
// 	pass, err := CheckRecaptchaToken("localhost", "6LdceoUUAAAAABn_I0c0gcWDrTPuxzqS0Hac42AE", "03AO9ZY1B-JCHTVDYEYMY20m8B6kF4sKRhW9wl61T8CVlDu-ISEtQ4n_EFShf0hnQ2jftIn1rYitNdyNJ8GEgY4i12UxXahNepaILqFsA986b0cVcN6eZ3C3_iKM-4KVV0hBkK59PeXdzKSuU_4Ic9kQGa6OclOpKSUErPMQVC0_Wra_7vQLGZ1g72WpYujXVBfhfQzszlsuDK3BLPQ8dq4H9YtULkFQB1X6tQ6ZHApNpN0uewW90hGy9N3fb_Jk6Y1qOXV5QFFb1m34ilZ-eWIvoMOj88nL2aJlQKJn0Ma_2fF8zHYnehoMuBet0dpe4HrwSP631t7Tmr")
// 	if assert.NoError(t, err) {
// 		t.Log(pass)
// 	}
// }

// func TestGetUserIpAddress(t *testing.T) {
// 	req, _ := http.NewRequest("POST", "http://127.0.0.1", nil)
// 	req.Header.Set("X-FORWARDED-FOR", "localhost,mickeyhost")
// 	ip, err := GetUserIpAddress(req)
// 	if assert.NoError(t, err) {
// 		if !assert.Equal(t, ip, "localhost") {
// 			t.Logf("Expected ip to be localhost but got: %v", ip)
// 			t.Fail()
// 		}
// 	}
// }

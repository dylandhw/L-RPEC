package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"time"
)

func SignRequest(r *http.Response, secretKey []byte) {
	now := time.Now()
	unixTime := now.Unix()

	// hmac hash
	h := hmac.New(sha256.New, secretKey)
}

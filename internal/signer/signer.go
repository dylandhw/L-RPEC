package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

func SignRequest(r *http.Request, secretKey []byte) {
	now := time.Now().Unix()
	timeString := strconv.FormatInt(now, 10)

	stringToBeHashed := r.Method + r.URL.Path + timeString
	// hmac hash
	hash := hmac.New(sha256.New, secretKey)

	hash.Write([]byte(stringToBeHashed))

	signature := hash.Sum(nil)

	r.Header.Add("X-Timestamp", timeString)
	r.Header.Add("X-Signature", hex.EncodeToString(signature))
}

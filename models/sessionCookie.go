package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)
/*
	cookie: username + hash(login time)
	session: username + login time
*/

type UserInfo struct {
	Username  string    `json:"userName"`
	SigninTime time.Time `json:"loginTime"`
}

type ResetInfo struct {
	Username string `json:"userName"`
	Email    string `json:"email"`
}

func GenerateHash(userInfo UserInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(userInfo.Username + userInfo.SigninTime.GoString()))
	return hex.EncodeToString(hasher.Sum(nil))
}

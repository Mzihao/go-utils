package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func main() {
	secret := []byte("the shared secret key here")
	message := []byte("the message to hash here")

	hash := hmac.New(sha256.New, secret)
	hash.Write(message)

	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))
	//4643978965ffcec6e6d73b36a39ae43ceb15f7ef8131b8307862ebc560e7f988

	// to base64
	base64.StdEncoding.EncodeToString(hash.Sum(nil))
	// RkOXiWX/zsbm1zs2o5rkPOsV9++BMbgweGLrxWDn+Yg=
}

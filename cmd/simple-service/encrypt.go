package main

import (
	"fmt"
	"crypto/cipher"
	"crypto/aes"
	"crypto/rand"
)


func main() {
	key := "keygopostmediumkeygopostmediumkm"
	encrypted, err := encrypt( []byte("Hellow World!"), key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Encrypted text: ", string(encrypted))
}

// encryption
func encrypt( data []byte, key string) (resp []byte, err error) {
	
	//string to []byte using aes encryption 
	// key := []byte("keygopostmediumkeygopostmediumkm") //32
	block ,err := aes.NewCipher([]byte(key))
	// fmt.Println(len(block))
	if err != nil {
		return resp, err
	}

	//gcm wrapper
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}

	//nounce creation
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return resp, err
	}
	// nonce := []byte("hellothisis")

	//Final encryption using Seal()
	dst := gcm.Seal(nil, nonce, data, []byte("test"))
	return dst, nil
}
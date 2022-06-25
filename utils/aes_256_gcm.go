package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"cryptopwd/entity"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

//NewGCMEncrypter aes-256-gcm 加密
func newGCMEncrypter(aesKey string, v []byte) (string, string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(aesKey)
	plaintext := v

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return fmt.Sprintf("%x", nonce), fmt.Sprintf("%x", ciphertext), nil

}

//NewGCMDecrypter aes-256-gcm 解密
func newGCMDecrypter(aesKey, nonceV, ciphertextV string) ([]byte, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(aesKey)
	nonce, _ := hex.DecodeString(nonceV)

	ciphertext, _ := hex.DecodeString(ciphertextV)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, err
}

func aesPwds(p entity.Pwds, key []byte) (entity.EncryptoPwds) {
	pwds := p.GetP()
	eps := entity.NewEncryptoPwds()
	for _, v := range *pwds {
		nonce, ciphertext, err := newGCMEncrypter(string(key), []byte(v.GetPasswd()))
		if err != nil {
			log.Fatalln(err.Error())
		}
		enpwd := entity.NewEncryptoPwd(v.GetName(), ciphertext, "aes", nonce)
		//fmt.Printf("enpwd: %v\n", enpwd)
		eps.AppendEP(*enpwd)
	}
	return *eps
}

func AesDeCrypto(ep entity.EncryptoPwds, sk string) *entity.Pwds {
	ps := entity.NewPwds()
	for _, v := range ep.GetP() {
		pwd := entity.NewPwd()
		pwd.SetName(v.GetName())
		nonce := v.GetNonce()
		enpasswd := v.GetEnPasswd()
		b, err := newGCMDecrypter(sk, nonce, enpasswd)
		if err != nil {
			log.Fatalln(err.Error())
		}
		pwd.SetPasswd(string(b))
		ps.AppendPwd(pwd)
	}
	return ps
}
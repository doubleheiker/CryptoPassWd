package utils

import (
	"crypto/rand"
	"cryptopwd/entity"
	"fmt"
	"log"
	"math/big"
)

// This function returns (secure) random bytes key
func generateBytes(n int) ([]byte, error) {
	var k []byte
	for i := 0; i < n; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(93))
		if err != nil {
			return nil, err
		}
		b.Add(b, big.NewInt(33))
		k = append(k, byte(b.Int64()))
	}
	return k, nil
}


func ChooseCryptoAlgorithm(p entity.Pwds) {
	var index int
	fmt.Println("请选择你想要使用的加密算法: ")
	fmt.Println("1.DES")
	fmt.Println("2.AES-256-GCM")
	fmt.Println("3.RSA")
	fmt.Scanln(&index)

	switch index {
	case 1:
	case 2:
		var c string
		fmt.Print("是否需要自设密钥？(y/n): ")
		fmt.Scanln(&c)
		for c != "y" && c != "n" {
			fmt.Println("请输入y或者n")
			fmt.Scanln(&c)
		}
		if c == "y" {
			// TODO
		} else {
			key, err := generateBytes(32)
			if err != nil {
				log.Fatalln(err.Error())
			}
			ep := aesPwds(p, key)
			WriteEnPwds(ep)
			fmt.Printf("加密完成，请务必保存好密钥: %v\n", string(key))
		}

	case 3:
	default:
		err := fmt.Errorf("%v", "请输入正确的序号")
		log.Fatalln(err.Error())
	}
}
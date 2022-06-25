package utils

import (
	"bufio"
	"cryptopwd/entity"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadPwd() entity.Pwds {
	var filepath string
	fmt.Println("输入文件的绝对路径(包括后缀): ")
	fmt.Scanln(&filepath)

	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	pwds := entity.NewPwds()
	s := bufio.NewScanner(f)
	for s.Scan() {
		arr := strings.Split(s.Text(), " ")
		pwds.AppendP(arr)
	}
	
	return *pwds
}

func ReadEnPwd() entity.EncryptoPwds {
	var filepath string
	fmt.Println("输入文件的绝对路径(包括后缀): ")
	fmt.Scanln(&filepath)

	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	ep := entity.NewEncryptoPwds()
	decoder := json.NewDecoder(f)
	err2 := decoder.Decode(&ep)

	if err2 != nil {
		log.Fatal(err.Error())
	}

	return *ep
}

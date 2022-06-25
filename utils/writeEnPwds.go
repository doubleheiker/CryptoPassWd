package utils

import (
	"cryptopwd/entity"
	"encoding/json"
	"log"
	"os"
)

func WriteEnPwds(ep entity.EncryptoPwds) {
	f, err := os.Create("./encryptoPasswd.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	e := json.NewEncoder(f)

	err2 := e.Encode(ep)

	if err2 != nil {
		log.Fatal(err.Error())
	}
}

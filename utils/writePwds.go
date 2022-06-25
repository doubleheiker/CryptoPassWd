package utils

import (
	"cryptopwd/entity"
	"log"
	"os"
)

func WritePwds(ps *entity.Pwds) {
	f, err := os.Create("./result")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	for _, v := range *ps.GetP() {
		s := v.GetName() + " " + v.GetPasswd()
		_, err2 := f.WriteString(s+"\n")
		if err2 != nil {
			log.Fatalln(err2.Error())
		}
	}
}
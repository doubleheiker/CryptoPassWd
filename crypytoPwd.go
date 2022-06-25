package main

import (
	"cryptopwd/utils"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Log: ")
	//log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetFlags(log.Ldate | log.Ltime)
}

/*
 * TODO List:
 * 1.输入单个网站和对应密码的加密存储
 * 2.读入密码本然后加密存储
 * 3.查询单个网站的对应的密码
 * 4.明文导出密码本
 * 5.选择加解密使用的加密算法
 */
func main() {
	fmt.Println("==================================")
	fmt.Println("选择你想要执行的操作: ")
	fmt.Println("1.新增单个密码口令加密存储")
	fmt.Println("2.读取明文密码列表加密存储")
	fmt.Println("3.查询指定网站的密码")
	fmt.Println("4.明文导出整个密码本")
	fmt.Print("选择你想要执行的操作: ")

	var index int
	_, err := fmt.Scanln(&index)
	if err != nil {
		return
	}
	switch index {
	case 1:

	case 2:
		p := utils.ReadPwd()
		utils.ChooseCryptoAlgorithm(p)
	case 3:
	case 4:
		ep := utils.ReadEnPwd()
		fmt.Print("请输入密钥: ")
		var sk string
		_, err := fmt.Scanln(&sk)
		if err != nil {
			return
		}
		p := utils.AesDeCrypto(ep, sk)
		utils.WritePwds(p)
	default:
		err := fmt.Errorf("%v", "请输入正确的序号")
		log.Fatalln(err.Error())
	}
}

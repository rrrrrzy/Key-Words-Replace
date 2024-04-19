package main

import (
	"fmt"
	"os"
	"strings"
)

func cmd_line_model() {
	var filename, old_content, new_content string
	fmt.Println("欢迎使用 KWR, 请按照提示输入内容：")

	fmt.Printf("请输入要需要替换的文件完整名称\n例如：您可以输入 example.txt \n>>> ")
	fmt.Scanf("%s", &filename)
	// fmt.Print(filename)
	ok, err := check_filename(filename)
	if !ok || err != nil {
		panic(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		panic("open file err: " + err.Error())
	}

	fmt.Print("请输入需要替换的关键词\n例如：您可以输入 old_name\n>>> ")
	fmt.Scanf("%s\n", &old_content)

	fmt.Print("请输入期望的关键词\n例如：您可以输入 new_name\n>>> ")
	fmt.Scanf("%s\n", &new_content)
	if new_content == "" {
		fmt.Printf("注意：即将从 '%s' 中删去关键词 '%s'，是否确认[y/n]\n>>> ", filename, old_content)
		var sure string
		fmt.Scanf("%s", &sure)
		fmt.Print(sure)
		if strings.Compare(sure, "y") != 0 && strings.Compare(sure, "Y") != 0 {
			fmt.Print("Process end!")
			os.Exit(0)
		}
	}
	fmt.Println("正在调用服务...")

	kwr_core(file, old_content, new_content)
}

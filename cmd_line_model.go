package main

// import (
// 	"fmt"
// 	"os"
// )

// func cmd_line_model() {
// 	var filename, old_content, new_content string
// 	fmt.Println("欢迎使用 KWR, 请按照提示输入内容：")

// 	fmt.Print("请输入要需要替换的文件完整名称\n例如：您可以输入 example.txt \n>>> ")
// 	fmt.Scanf("%s\n", filename)

// 	ok, err := check_filename(filename)
// 	if !ok || err != nil {
// 		panic(err)
// 	}
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		panic("open file err: " + err.Error())
// 	}
// 	file_info, err := file.Stat()
// 	if err != nil {
// 		panic("get file status err: " + err.Error())
// 	}

// 	fmt.Print("请输入需要替换的关键词\n例如：您可以输入 old_name\n>>> ")
// 	fmt.Scanf("%s\n", old_content)

// 	fmt.Print("请输入期望的关键词\n例如：您可以输入 new_name\n>>> ")
// 	fmt.Scanf("%s\n", new_content)
// 	// if len(new_content) == 0 {
// 	// 	fmt.Print("")
// 	//	未来做空词替换时，可以加以判断
// 	// }
// 	fmt.Println("正在调用服务...")

// 	kwr_core(file, file_info.Size(), old_content, new_content)
// }

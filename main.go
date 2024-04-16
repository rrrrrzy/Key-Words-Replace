package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func kwr_core(file *os.File, size int64, old_name, new_name string) {
	content := make([]byte, size)

	_, err := file.Read(content)
	if err != nil {
		fmt.Print("read file fail!" + err.Error())
	}

	s_content := string(content)
	new_content := strings.ReplaceAll(s_content, old_name, new_name)
	new_b := []byte(new_content)
	new_file, err := os.Create("new_novel.txt")
	if err != nil {
		fmt.Print("create file error: ", err.Error())
	}
	new_file.Write(new_b)

	fmt.Print("Success! New file has been created.")

}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Wrong PARAM Number!\nIt should be '$ kwr file.txt old_name new_name' ")
	}
	re, err := regexp.Compile(".txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	ok := re.MatchString(os.Args[1])
	if !ok {
		log.Fatal("Wrong type of filename! It should be like 'filename.txt'")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("open file err: " + err.Error())
	}

	stat, err := file.Stat()
	if err != nil {
		fmt.Print("get file status err: " + err.Error())
	}

	size := stat.Size()
	kwr_core(file, size, os.Args[2], os.Args[3])
}

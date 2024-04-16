package main

import (
	"fmt"
	"os"
	"strings"
)

func filecut(file *os.File, size int64) {
	fmt.Print("The size of file is ", size, "byte, file cut process is about to start...\n")

	content := make([]byte, size)

	_, err := file.Read(content)
	if err != nil {
		fmt.Print("read file fail!" + err.Error())
	}

	s_content := string(content)
	new_content := strings.ReplaceAll(s_content, "读博阁", "书包")
	new_b := []byte(new_content)
	new_file, err := os.Create("new_novel.txt")
	if err != nil {
		fmt.Print("create file error: ", err.Error())
	}
	new_file.Write(new_b)
	fmt.Print("Success! New file has been created.")

}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print("open file err: " + err.Error())
	}

	stat, err := file.Stat()
	if err != nil {
		fmt.Print("get file status err: " + err.Error())
	}
	size := stat.Size()
	if size < 10000 {
		fmt.Printf("file size is: %v byte, do not need cut~~\n", size)
		file.Close()
		return
	} else {
		fmt.Print("file size is larger than 100kB, wait for cut!\n")
		filecut(file, size)
	}
}

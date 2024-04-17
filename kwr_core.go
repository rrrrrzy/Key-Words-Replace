package main

import (
	"fmt"
	"os"
	"strings"
)

func kwr_core(file *os.File, size int64, old_name, new_name string) {
	content := make([]byte, size)

	_, err := file.Read(content)
	if err != nil {
		panic("read file fail!" + err.Error())
	}

	s_content := string(content)
	new_content := strings.ReplaceAll(s_content, old_name, new_name)
	new_b := []byte(new_content)
	new_file, err := os.Create("new_file.txt")
	if err != nil {
		panic("create file error: " + err.Error())
	}
	new_file.Write(new_b)

	fmt.Print("Success! New file has been created in name \"new_file.txt\".")
	// os.Exit(0)
}

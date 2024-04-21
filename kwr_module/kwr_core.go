package kwr_module

import (
	"fmt"
	"os"
	"strings"
)

func kwr_core(file *os.File, old_name, new_name string) {
	status, err := file.Stat()
	if err != nil {
		panic("get file status err: " + err.Error())
	}
	size := status.Size()
	content := make([]byte, size)

	_, err = file.Read(content)
	if err != nil {
		panic("read file fail!" + err.Error())
	}
	file.Close()

	s_content := string(content)
	new_content := strings.ReplaceAll(s_content, old_name, new_name)
	new_b := []byte(new_content)
	new_file, err := os.Create("new_" + status.Name())
	if err != nil {
		panic("create file error: " + err.Error())
	}
	new_file.Write(new_b)
	new_file.Close()

	fmt.Printf("Success! New file has been created in name \"new_%s\" .\n", status.Name())
}
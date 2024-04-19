package main

import (
	"os"
	// "strings"
)

func main() {
	if len(os.Args) <= 1 {
		cmd_line_model()
		return
	}
	if len(os.Args) == 3 {
		return
	}
	if len(os.Args) != 4 {
		panic("Wrong PARAM Number!\nIt should be '$ kwr file.txt old_content new_content' ")
	}
	param_model()

}

package main

import (
	. "kwr/kwr_module"
	"os"
)

// "strings"

func main() {
	if len(os.Args) <= 1 {
		Cmd_line_model()
		return
	}
	if len(os.Args) == 3 {
		return
	}
	if len(os.Args) != 4 {
		panic("Wrong PARAM Number!\nIt should be '$ kwr file.txt old_content new_content' ")
	}
	Param_model()

}

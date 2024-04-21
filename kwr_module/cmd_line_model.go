package kwr_module

import (
	"fmt"
	"os"
	"strings"
)

func Cmd_line_model() {
	var filename, old_content, new_content string
	fmt.Println("Welcome to use KWR, pls input as notice below")

	fmt.Printf("Pls input the filename where to replace the words\ne.g you can input example.txt \n>>> ")
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

	fmt.Print("Pls input the old_name(will be replaced)\n>>> ")
	fmt.Scanf("%s\n", &old_content)

	fmt.Print("Pls input the expect word to replace former\nTips: if you just want to delete the former word, press ENTER directly! \n>>> ")
	fmt.Scanf("%s\n", &new_content)
	if new_content == "" {
		fmt.Printf("ATTENTION：we'll delete '%s' from '%s'，are you sure that[y/n]\n>>> ", old_content, filename)
		var sure string
		fmt.Scanf("%s", &sure)
		fmt.Print(sure)
		if strings.Compare(sure, "y") != 0 && strings.Compare(sure, "Y") != 0 {
			fmt.Print("Process end!")
			os.Exit(0)
		}
	}
	fmt.Println("wait pls ...")

	kwr_core(file, old_content, new_content)
}

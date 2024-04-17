package main

import (
	"log"
	"os"
	// "strings"
)

func main() {
	// if len(os.Args) <= 1 {
	// 	cmd_line_model()
	// 	return
	// }
	//fmt.Print(os.Args)
	if len(os.Args) != 4 {
		log.Fatal("Wrong PARAM Number!\nIt should be '$ kwr file.txt old_content new_content' ")
	}
	param_model()
	// re, err := regexp.Compile(".txt")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// ok := re.MatchString(os.Args[1])
	// if !ok {
	// 	log.Fatal("Wrong type of filename! It should be like 'filename.txt'")
	// }

	// file, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal("open file err: " + err.Error())
	// }

	// stat, err := file.Stat()
	// if err != nil {
	// 	fmt.Print("get file status err: " + err.Error())
	// }

	// size := stat.Size()
	// kwr_core(file, size, os.Args[2], os.Args[3])

}

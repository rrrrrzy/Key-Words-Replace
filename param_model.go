package main

import (
	"fmt"
	"log"
	"os"
)

func param_model() {
	ok, err := check_filename(os.Args[1])
	if !ok || err != nil {
		panic(err)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("open file err: " + err.Error())
	}

	stat, err := file.Stat()
	if err != nil {
		fmt.Print("get file status err: " + err.Error())
	}

	kwr_core(file, stat.Size(), os.Args[2], os.Args[3])
}

package main

import (
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

	kwr_core(file, os.Args[2], os.Args[3])
}

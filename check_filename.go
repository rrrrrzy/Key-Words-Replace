package main

import (
	"fmt"
	"regexp"
)

func check_filename(filename string) (bool, error) {
	re, err := regexp.Compile(".txt")
	if err != nil {
		return false, err
	}
	ok := re.MatchString(filename)
	if !ok {
		return false, fmt.Errorf("wrong type of filename! It should be like 'filename.txt'")
	}
	return true, nil
}

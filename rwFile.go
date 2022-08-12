package main

import (
	"os"
)

func readfile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
func writefile(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err1 := file.WriteString(data)
	if err1 != nil {
		return err
	}
	file.Close()
	return nil
}

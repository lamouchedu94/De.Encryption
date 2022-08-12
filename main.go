package main

import (
	"fmt"
)

func main() {
	s := Settings{}
	err := s.arguments()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(s.Sentence)
	//for encryption
	var AsciiSentence []int

	switch {
	case !s.file:
		if s.Encryption != "" {
			AsciiSentence = stringToAscii(s.Encryption)
		}
		if s.Decryption != "" {
			AsciiSentence = stringToAscii(s.Decryption)
		}
	case s.file:
		var StringVal string
		switch {
		case s.Encryption != "":
			StringVal, err = readfile(s.Encryption)
			AsciiSentence = stringToAscii(StringVal)
		case s.Decryption != "":
			StringVal, err = readfile(s.Decryption)
			AsciiSentence = stringToAscii(StringVal)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

	}
	var endS string
	if s.Encryption != "" {
		_, endS = encryption(AsciiSentence, s.Key)

		if s.verbose {
			fmt.Println("Encrypted sentence :", endS)
		}
	}
	//For Decryption
	if s.Decryption != "" {
		endS = decryption(AsciiSentence, s.Key)
		if s.verbose {
			fmt.Println("Decrypted sentence :", endS)
		}
	}
	if s.file {
		writefile(s.path, endS)
	}
}

func stringToAscii(sentence string) []int {
	var AsciiSentence []int
	for _, val := range sentence {
		AsciiSentence = append(AsciiSentence, int(val))
	}
	return AsciiSentence
}

func encryption(AsciiSentence []int, key int) ([]int, string) {
	var modified []int
	var modifiedString string
	for _, val := range AsciiSentence {
		modified = append(modified, int(val)+key)
	}
	for _, val := range modified {
		modifiedString += string(val)
	}

	return modified, modifiedString
}
func decryption(AsciiSentence []int, key int) string {

	var modified []int
	var modifiedString string
	for _, val := range AsciiSentence {
		modified = append(modified, int(val)-key)
	}
	for _, val := range modified {
		modifiedString += string(val)
	}

	return modifiedString
}

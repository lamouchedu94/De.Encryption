package main

import (
	"errors"
	"flag"
	"os"
)

type Settings struct {
	Encryption string
	Decryption string
	Key        int
	file       bool
	path       string
	verbose    bool
}

func (s *Settings) arguments() error {
	flag.StringVar(&s.Encryption, "e", "", "encryption mode")
	flag.StringVar(&s.Decryption, "d", "", "decryption mode")
	flag.IntVar(&s.Key, "k", 0, "Key to use")
	flag.BoolVar(&s.verbose, "v", false, "display result")
	flag.Parse()
	var err error
	s.isdirectory()
	if s.Encryption == "" && s.Decryption == "" {
		err = errors.New("Encryption (-e) or Decryption (-d)")
		return err
	}

	if s.Key == 0 {
		err = errors.New("No key given or key = 0")
		return err
	}
	s.Key += 160
	return nil
}

func (s *Settings) isdirectory() error {
	var err error
	switch {

	case s.Encryption != "":
		_, err := os.Stat(s.Encryption)
		if err != nil {
			//is not file
			//fmt.Println("is sentence")
			s.file = false
			return err
		}
		s.path = s.Encryption
	case s.Decryption != "":
		_, err = os.Stat(s.Decryption)
		if err != nil {
			//is not file
			//fmt.Println("is sentence")
			s.file = false
			return err
		}
		s.path = s.Decryption
	}
	//fmt.Println("is file")
	s.file = true
	return nil
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once!")
	}
	names := strings.Split(v, ",")
	for _, item := range names{
		s.Names = append(s.Names,item)
	}

	return nil
}

func fun() {
	var myName NamesFlag
	minusK := flag.Int("k",0,"an int")
	minus0 := flag.String("o","Mihalis","the name")
	flag.Var(&myName,"names","comma-separated list")
	flag.Parse()

	fmt.Println("-k:",*minusK)
	fmt.Println("-o:",*minus0)

	for i, item := range myName.GetNames() {
		fmt.Println(i, item)
	}
	fmt.Println("Remaing command-line arugments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}

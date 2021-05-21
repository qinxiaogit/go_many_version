package main

import (
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

var ErrEmpty = errors.New("empty string")


type stringService struct {}

func (stringService)Uppercase(string2 string)(string,error){
	if string2==""{
		return "",ErrEmpty
	}
	return strings.ToUpper(string2),nil
}

func (stringService)Count(string2 string)int{
	return len(string2)
}

type ServiceMiddleware func(service StringService)StringService

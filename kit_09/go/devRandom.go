package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func generRandom(){
	f, err := os.Open("/dev/random")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	err = binary.Read(f, binary.LittleEndian, &seed)
	if err != nil {
		return 
	}
	fmt.Println("Seed:", seed)

}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var total struct {
	sync.Mutex
	value int
}

func work(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<rand.Int();i++ {
		total.value++
	}
}

func workPrint(wg  *sync.WaitGroup)  {
	defer wg.Done()
	for  {
		fmt.Println(total.value)
		time.Sleep(time.Second)
	}

}



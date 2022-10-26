package main

import (
	"fmt"
	"runtime"
)

func main() {
	//pr()
	//fmt.Println("main() ended!")
	/**
	myHeap := &heapFloat32{1.2, 2.1, 3.1, -100.1}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(float32(-100.2))
	myHeap.Push(float32(0.2))

	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)

	values := list.New()
	values.PushFront(1)
	values.PushBack("aaaa")

	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())

	x := a{100, 200.12, "Struct a"}
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)

	//fun()
	//for i := 0; i < 100; i++ {
	//	generRandom()
	//}
*/
	//exportMain()

	//for true {
	//	fmt.Println(".")
	//	time.Sleep(20*time.Second)
	//}

	fmt.Println(runtime.GOMAXPROCS(0))
}


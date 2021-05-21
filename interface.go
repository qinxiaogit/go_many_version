package main

import "fmt"

type Duck interface {
	Quack()   // 鸭子叫
	DuckGo()  // 鸭子走
}

type Chicken struct {
}

func (c Chicken) IsChicken()  {
	fmt.Println("我是小鸡")
}

func (c Chicken) Quack() {
	fmt.Println("嘎嘎")
}

func (c Chicken) DuckGo() {
	fmt.Println("大摇大摆的走")
}
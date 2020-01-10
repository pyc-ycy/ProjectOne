// IntelliJ IDEA
// ProjectOne
// 2020/1/10
// 10:02
// 彭友聪

package main

import "fmt"
import "../FuncSet"

var v1, v2 int

// 因式分解关键字的写法一般用于全局变量的声明
var (
	v3 int
	v4 bool
)

var v5, v6 int = 1, 2
var v7, v8 = 123, "Hello"

func main() {
	//fmt.Println("Hello world")
	//println(v1, v2, v3, v4, v5, v6, v7, v8)
	//variableNothing()
	//var b = true
	//fmt.Println(b)
	//var a string = "abc"
	//fmt.Println("hello world", a)
	//constVariable()
	//FuncSet.SimpleOne()
	//FuncSet.SimpleTwo()
	//FuncSet.SimpleThree()
	FuncSet.One()
}

func variableNothing() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	var intValue int
	// 只能在函数体中出现
	intValue, intValue1 := 2, 4
	str := "try a operator ':='"
	fmt.Printf("intValue:%v, intValue1:%v\n", intValue, intValue1)
	fmt.Println(str)
}

func constVariable() {
	const LENGTH = 10
	const WIDTH = 5
	var area int
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d\n", area)
	println(a, b, c)
}

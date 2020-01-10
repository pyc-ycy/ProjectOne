// IntelliJ IDEA
// ProjectOne
// 2020/1/10
// 11:37
// 彭友聪

package FuncSet

import (
	"fmt"
	"unsafe"
)

const(
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)
func SimpleOne(){
	fmt.Println(a,b,c)
}
func SimpleTwo(){
	const(
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}
const(
	i = 1<<iota
	j = 3<<iota
	k
	l
)
func SimpleThree(){
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}
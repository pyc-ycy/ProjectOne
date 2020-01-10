// IntelliJ IDEA
// ProjectOne
// 2020/1/10
// 15:52
// 彭友聪

package FuncSet

import "fmt"

func One() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("a+b:%d\n", c)
	c = a - b
	fmt.Printf("a-b:%d\n", c)
	c = a * b
	fmt.Printf("a*b:%d\n", c)
	c = a / b
	fmt.Printf("a/b:%d\n", c)
	c = a % b
	fmt.Println("a%b:", c)
	a++
	fmt.Printf("a++:%d\n", a)
	a = 21
	a--
	fmt.Printf("a--:%d\n", a)
}

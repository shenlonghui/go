package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Println("1%d", a)
	fmt.Println("2%d", b)

	max(&a, &b)
	fmt.Println("3%d啊实打实", a)
	fmt.Println("4%d阿瑟东", b)
}

func max(x, y *int) int {
	var shen int

	shen = *x
	*x = *y
	*y = shen
	return shen
}

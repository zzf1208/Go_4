package main

import "fmt"

func main() {
	n1 := 13 //二进制为1101
	n2 := 3  //二进制为11
	//两个对应的二进制位有1个为1就为1
	fmt.Println(n1 | n2)
}

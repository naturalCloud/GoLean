package main

import "fmt"

// 闭包函数
func adder() func(int) int {

	return func(i int) int {
		return i * (i + 1)
	}
}

//传统闭包

type model func(int) int

func (model) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func sum(basic int) model {

	H := basic * 5
	return func(i int) int {
		if basic*H > 9 {
			return 3 * 1
		}

		return i * 1
	}

}

func main() {
	a := adder()

	fmt.Println(a(45))

	fmt.Println(sum(9)(5))

}

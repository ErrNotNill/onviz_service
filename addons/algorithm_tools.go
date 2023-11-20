package addons

import (
	"fmt"
	"math/big"
)

func FindHigherNum() {
	n := 10000

	f := big.NewInt(1)

	for i := 2; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}

	fmt.Println(f)
}
func Factorial() {
	var arr = []int{2, 5, 8, 1, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	fmt.Println("len(arr)-1: ", len(arr)-2)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("arr", arr)
	}
}

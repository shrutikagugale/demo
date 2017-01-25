// demo2
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	const n = 500000000
	fmt.Println(math.Sin(n))
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}

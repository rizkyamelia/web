package main

import "fmt"

func main() {
	var jumlah int
	var a int
	var b int
	var x [36]int
	a = 0
	b = 4000000
	jumlah = 0

	for i := a; i <= 35; i++ {
		if i < 2 {
			x[i] = i + 1
		} else {
			x[i] = x[i-2] + x[i-1]
		}
	}

	for i := a; i <= 35; i++ {
		if x[i] <= b && x[i]%2 == 0 {
			jumlah += x[i]
		}
	}

	fmt.Println(x)
	fmt.Printf("%d", jumlah)
}

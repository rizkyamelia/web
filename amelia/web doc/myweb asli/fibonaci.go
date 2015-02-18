package main

import "fmt"

func main() {
	var x [70]int
	var jumlah int
	jumlah = 0

	for i := 0; i <= 35; i++ {
		if i < 2 {
			x[i] = i + 1
		} else {
			x[i] = x[i-2] + x[i-1]
		}
	}

	for i := 0; i <= 35; i++ {
		if x[i] <= 4000000 && x[i]%2 == 0 {
			jumlah += x[i]
		}
	}

	fmt.Println(x)
	// fmt.Printf("%d", jumlah)
}

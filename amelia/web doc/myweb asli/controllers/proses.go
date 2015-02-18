package controllers

import (
	"amelia/gomvc"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ProsesController struct {
	gomvc.Controller
}

func (c *ProsesController) Project1(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	val1 := req.FormValue("param01")
	val2 := req.FormValue("param02")
	hasil := 0

	a, _ := strconv.Atoi(val1)
	b, _ := strconv.Atoi(val2)
	str := ""

	for val3 := sqrt(b); b >= a; val3-- {
		if b%val3 == 0 && Prime(val3, a) {
			fmt.Println(val3)
			hasil = val3
			str = str + " " + strconv.Itoa(hasil)
			break
		}
	}

	d1 := []byte(str)
	err := ioutil.WriteFile("dat1.txt", d1, 0644)
	check(err)

	c.ServeJson(hasil)
}

func Prime(b int, a int) bool {
	if b <= a {
		return false
	}

	for val3 := sqrt(b); val3 >= a; val3-- {
		if val3 == a {
			return true
		}

		if b%val3 == 0 {
			return false
		}
	}
	return true
}

func sqrt(b int) int {
	return int(math.Sqrt(float64(b)))
}

// c.ServeJson(hasil)

// fmt.Println(a, b, val3)

// 	data := make(map[string]interface{})
// 	data["val1"] = a
// 	data["val2"] = b
// 	data["hasil"] = val3

// 	c.ServeJson(data)
// }

// func test() {
// 	fmt.Println("test")
// }

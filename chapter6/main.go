package main

import "fmt"

func main() {
	// x := [5]float64{10, 20, 30, 40, 50}

	x := [5]float64{10,
		20,
		30,
		40,
		50, // Bat buoc co dau ","
	}

	// cach 1 duyet mang
	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	// cach 2 duyet mang
	var total2 float64
	for _, value := range x {
		total2 += value
	}
	fmt.Println(total2 / float64(len(x)))

	var a []float64
	fmt.Println("a=", a)

	b := make([]float64, 5, 10)
	fmt.Println("b=", b)

	c := []float64{1, 2, 3, 4, 5}
	fmt.Println("c=", c)

	d := c[0:5]
	fmt.Println("d=", d)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println(slice1, slice2)

	e := make(map[string]int)
	e["key"] = 10
	fmt.Println(e["key"])

	f := make(map[int]string)
	f[10] = "key"
	fmt.Println(f[10])
}

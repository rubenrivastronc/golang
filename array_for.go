package main

import "fmt"

func main() {

	array := [3]int{1, 2}
	array[2] = 20

	var matriz [2][2]int

	matriz[0][1] = 1
	fmt.Println(matriz[0][1])

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

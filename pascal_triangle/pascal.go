package main

import "fmt"

func main() {
	//var s int
	// fmt.Scanln(&s)
	//a:=make([][]int,s)
	var a [6][6]int
	b := 1
	//var c int
	a[0][0] = b
	a[1][0] = b
	a[1][1] = b
	for i := 2; i < 6; i++ {
		for j := 0; j < i; j++ {
			a[i][0] = b
			a[i][j+1] = a[i-1][j] + a[i-1][j+1]
			//fmt.Print(a[i][j])
		}
		// fmt.Println()
	}
	for i := 0; i < 6; i++ {
		for k := 0; k < 6*2/2-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {

			fmt.Printf("%d ", a[i][j])

			//fmt.Print(" ")
		}
		fmt.Println()
	}
}

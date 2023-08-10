package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter the input string = ")
	fmt.Scanln(&input)
	input1 := "go"
	output := rearrangeString(input, input1)
	fmt.Println(output)
}

func rearrangeString(s, s1 string) int {
	count := 0
	slice := []int32(s)
	fmt.Println(slice)
	slice1 := []int32(s1)
	fmt.Println(slice1)

	slice2 := slice1[0]
	slice3 := slice1[1]

	for i := 0; i < len(slice)-1; i++ {
		s2 := slice[i]
		s3 := slice[i+1]
		if s2 == slice2 && s3 == slice3 {
			count++
		}
	}
	return count
}

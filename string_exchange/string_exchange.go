package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter the input string = ")
	fmt.Scanln(&input)
	output := rearrangeString(input)
	fmt.Println(output)
}

func rearrangeString(s string) string {
	if len(s)%2 != 0 {
		return "Input string length must be even."
	}
	//count :=0
	slice := []int32(s)
	fmt.Println(slice)
	for i := 2; i < len(slice)-1; i += 4 {
		//count++
		// if count == 3{
		slice[i], slice[i+1] = slice[i+1], slice[i]
		//i++
		// count=0
		// }
	}

	return string(slice)
}

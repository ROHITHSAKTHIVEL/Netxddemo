package main

import "fmt"

func main() {
	var num string
	var oddsum int32 = 0
	var evensum int32 = 0
	var totalsum int32 = 0
	var temp1 int32 = 0
	var temp2 int32 = 0
	var temp3 int32 = 0
	fmt.Scanln(&num)
	slice := []int32(num)

	if len(num) < 13 || len(num) > 16 {
		fmt.Println("Print the valid number")
	} else {
		for i := 0; i < len(slice); i++ {
			if slice[i] >= 48 && slice[i] <= 57 {
				slice[i] = slice[i] - 48
			}
		}
		if slice[0] == 3 && slice[1] == 7 || slice[0] == 4 || slice[0] == 5 || slice[0] == 6 {
			for j := len(slice) - 1; j >= 0; j-- {

				if j%2 != 0 {
					oddsum += slice[j]
				} else {
					val := slice[j]
					temp3 = val * 2

					if temp3 > 9 {
						temp1 = temp3 % 10
						temp3 = temp3 / 10
						temp2 = temp3 % 10
						temp1 = temp1 + temp2
						evensum += temp1
					} else {
						evensum += temp3
					}
				}
			}
			totalsum = oddsum + evensum
			if totalsum%10 == 0 {

				fmt.Println("ceadit card is valid", slice)
			} else {
				fmt.Println("ceadit card is NotValid", slice)
			}
		}
	}
}

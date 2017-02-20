package main

import "fmt"

func main() {
	var Y int
	fmt.Scan(&Y)

	switch {
	case Y < 1918:
		if Y%4 == 0 {
			fmt.Printf("12.09.%d", Y)
		} else {
			fmt.Printf("13.09.%d", Y)
		}
	case Y == 1918:
		fmt.Println("26.09.1918")
	case Y > 1918:
		if Y%4 == 0 && (Y%100 != 0 || Y%400 == 0) {
			fmt.Printf("12.09.%d", Y)
		} else {
			fmt.Printf("13.09.%d", Y)
		}
	}
}

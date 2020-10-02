package main

import (
	"fmt"
)

// func main() {
// 	// for init; condition; post {}
// 	for i := 0; i <= 100; i++ {
// 		fmt.Println(i)

// 	}
// }

// nested loop

// func main() {
// 	// for init; condition; post {}
// 	for i := 0; i <= 10; i++ {
// 		fmt.Printf("The outer loop: %d\n", i)
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("\t\t The inner loop: %d\n", j)

// 		}
// 	}
// }

// for statement with breaks and continue
// func main() {
// 	x := 1
// 	for {
// 		x++
// 		if x > 100 {
// 			break
// 		}

// 		if x%2 != 0 {
// 			continue
// 		}

// 		fmt.Println(x)
		
// 	}
// 	fmt.Println("done.")
// }

// if else statments
func main() {
	x := 434
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")
	} else if x == 42 {
		fmt.Println("our value was 42")
	} else if x == 43 {
		fmt.Println("our value was 43")
	} else {
		fmt.Println("our value was not 40, 41, 42, or 43")
	}
}
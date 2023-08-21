package main

import (
	"fmt"
)

/* Full_Pyramid
        *
      * * *
    * * * * *
  * * * * * * *
* * * * * * * * *

*/
func Full_Pyramid(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num+i-1; j++ {
			if j <= num-i {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}

/*   Inverted Pyramid
* * * * * * * * *
  * * * * * * *
    * * * * *
      * * *
        *
*/
func Inverted_Pyramid(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i+(num-i)*2; j++ {
			if i-1 >= j {
				fmt.Printf(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

}

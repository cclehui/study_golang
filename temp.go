package main

import (
	"fmt"
	//"os"
	//"bufio"
	//"strconv"
	//"io/ioutil"
)

var f = func(i int) {
	fmt.Print("xiiiix")
}

func main() {

	fmt.Printf("%04d\n", 1)
	return

	var f = func(i int) {
		fmt.Print(i)

		if i > 0 {
			f(i - 1)
		}
	}

	f(10)
	//fmt.Print("\nxxxx\n")

	//fmt.Println(rand.Int())

}

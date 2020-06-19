package main

import "fmt"

func testSlice(data []int) {

	fmt.Printf("%p, %v, %v, %v, %v\n", &data, data, "ttttttt", len(data), cap(data))

	data = append(data, 111111111)
	data[0] = 9999

	fmt.Printf("%p, %v, %v, %v, %v\n", &data, data, "ttttttt", len(data), cap(data))

}

func main() {

	data := make([]int, 1, 20)
	fmt.Printf("%p, %v, %v, %v, %v\n", &data, data, "mmmmmmmm", len(data), cap(data))

	testSlice(data)

	fmt.Printf("%p, %v, %v, %v, %v\n", &data, data, "mmmmmmmm", len(data), cap(data))
	return
}

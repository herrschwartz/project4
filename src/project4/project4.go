package main

import "fmt"

func main() {

	w := 100

	for i := 1; i < 20; i++ {
		if w == i {
			fmt.Println("terminating")
			break
		}
		fmt.Println(i)
		w = w - 10
	}
	fmt.Println("Goodnight!")

	var isEven bool = true
	var even []int
	var odd []int
	for i := 0; i < 30; i++ {
		if isEven == true {
			even = append(even, i)
			isEven = false

		} else {
			odd = append(odd, i)
			isEven = true
		}
	}

	var Formula map[string]string
	Formula = make(map[string]string)

	Formula["Rectangle"] = "Base * Height"
	Formula["Circle"] = "π * Radius^2"
	Formula["Triagle"] = "1/2 * Base * Height"

	fmt.Print("The formula for the area of a circle= ")
	fmt.Println(Formula["Circle"])
	fmt.Println(even)
	fmt.Println(odd)

}

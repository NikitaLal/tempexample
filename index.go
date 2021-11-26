package main

import "fmt"

func main() {
	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("square", 40))

}
func calculation(str string, y ...int) int {
	area := 1
	for _, val := range y {
		fmt.Println(val)
		if str == "Rectangle" {
			area *= val
		} else if str == "square" {
			area = val * val
		}
	}
	return area

}

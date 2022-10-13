package main

import (
	"fmt"
)

type Student struct {
	rollno int
	name   string
	marks  int
}

func main() {
	// const welcome string = "Hello world"
	// //welcome := "Hello world";
	// height := 10
	// if height > 10 {
	// 	println("greater than 10")
	// } else {
	// 	println("less than equal to 10")
	// }
	// println(welcome)

	// switch {
	// case height > 10:
	// 	println("greater than 10")
	// case height < 10:
	// 	println("less than 10")
	// default:
	// 	println("equal to 10")
	// }

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	println("os x")
	// case "linux":
	// 	println("linux machine")
	// case "windows":
	// 	println(os)
	// 	println("windows machine")
	// default:
	// 	println("something else")
	// }

	// planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "nepture"}
	// fmt.Println(planets)

	// var planentSlice []string
	// planentSlice = append(planentSlice, "earth")
	// fmt.Println(planentSlice)
	ages := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}
	temp("anu")
	add, sub := addSubtract(15, 10)
	println(add)
	println(sub)

	var s1 = Student{101, "sunil", 55}
	fmt.Println(s1)
}
func addSubtract(a, b int) (int, int) {
	return a + b, a - b
}
func temp(name string) {
	println("hello " + name)
}

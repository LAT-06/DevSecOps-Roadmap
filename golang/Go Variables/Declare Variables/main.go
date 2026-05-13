package main
import ("fmt")

var x int
var y int = 2
var z = 3

func main() {
	var student1 string = "LAT" //type is string
	var student2 = "Thinh Lam" //type is inferred
	x := 2 //type is inferred

	fmt.Println(student1) // Output: LAT
	fmt.Println(student2) // Output: Thinh Lam
	fmt.Println(x) // Output: 2

	var a string
	var b int
	var c bool

	fmt.Println(a) // Output: (empty string)
	fmt.Println(b) // Output: 0
	fmt.Println(c) // Output: false

	a = "Hello"
	fmt.Println(a) // Output: Hello

	var j, k, l int = 1, 2, 3
	fmt.Println(j, k, l) // Output: 1 2 3
	u, v, w := "U", "are", "gay"
	fmt.Println(u, v, w) // Output: U are gay
}
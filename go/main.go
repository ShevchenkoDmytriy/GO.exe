// package main

// import "fmt"

// func main() {
// fmt.Println("Hello World!");
// var a int = 10
// fmt.Println(a)
// var num float64 = 22.22
// fmt.Println(num)
// b := 22
// fmt.Println(b)
// name := "asfaf"
// fmt.Println(name)
// fmt.Println(len(name))
// name1 := "qfqfgo run main.go"
// fmt.Println("Whats your name?")
// fmt.Scan(&name1)
// fmt.Println("Hello " + name1)
// var age int
// fmt.Println("How old are you?")
// fmt.Scan(&age)
// fmt.Println("You " + fmt.Sprint(age))
// num := 0
// if num > 0 {
// 	fmt.Print(">")
// } else if num < 0 {
// 	fmt.Println("<")
// } else {
// 	fmt.Println("=")
// }
// var a, b, c float64
// fmt.Println("Enter a")
// fmt.Scan(&a)
// fmt.Println("Enter b")
// fmt.Scan(&b)
// fmt.Println("Enter c")
// fmt.Scan(&c)
// D := (b * b) - 4*a*c
// if D > 0 {
// 	x1 := (-b + math.Sqrt(D)) / 2 * a
// 	x2 := (-b - math.Sqrt(D)) / 2 * a
// 	fmt.Println(fmt.Sprint(x1) + " " + fmt.Sprint(x2))
// } else if D < 0 {
// 	fmt.Println("Error")
// } else {
// 	x1 := (-b) / (2 * a)
// 	fmt.Println(x1)
// }
// fmt.Scan(&a)
// for i := 1; i <= 100; i++ {
// 	if i%2 != 0 {
// 		continue
// 	}
// 	fmt.Println(i)
// }
// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// for _, row := range matrix {
// 	for _, col := range row {
// 		fmt.Printf("%d", col)
// 	}
// 	fmt.Printf("\n")
// }

// world := []string{"H", "E", "L", "L", "O", "W", "O", "R", "L", "D"}
// mas := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// for i := 0; i < len(world); i++ {
// 	for j := 0; j < len(mas); j++ {
// 		if world[i] == mas[j] {
// 			fmt.Printf(mas[j])
// 		}
// 	}
// }
//  num := 6
// switch {
//  case num%2 == 0:
// 	fmt.Println("/2")
// fallthtough
//  case num%2 != 0:
//  	fmt.Println("!/2")
//  }
// fmt.Println("+ - * / %")
// var action string
// fmt.Scan(&action)
// var a float64
// var b float64
// fmt.Println("Enter num ")
// fmt.Scan(&a)
// fmt.Println("Enter num ")
// fmt.Scan(&b)
// switch {
// case action == "+":
// 	fmt.Println(a + b)
// case action == "-":
// 	fmt.Println(a - b)
// case action == "/":
// 	fmt.Println(a / b)
// case action == "*":
// 	fmt.Println(a * b)
// }
// names := []string{"q", "b", "c"}
// for i := 0; i < len(names); i++ {
// 	fmt.Println(names[i])
// }
// var res float64 = 0
// num := []float64{1, 2, 3, 4, 5, 6}
// for i := 0; i < len(num); i++ {

// 	res += num[i]

// }
// fmt.Println(res)
// array := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// for i := 0; i < len(array); i++ {
// 	for j := 0; j < len(array[i]); j++ {
// 		fmt.Println(array[i][j])
// 	}
// }
// slice := []int{1, 2, 3, 4, 5, 6}
// slice = append(slice, 2)
// sort.Ints(slice)
// fmt.Println(slice)
// slice := []int{1, 2, 3, 4, 5, 6}
// for _, el := range slice {
// 	fmt.Printf("%d\t", el)
// }
// age := 10
// name := "Tom"
// // fmt.Println("age - " + fmt.Sprint(age))
// fmt.Printf("Age - %d.Name - %s", age, name)
// var money map[string]int = map[string]int{
// 	"dollar": 3000,
// 	"euro":   2000,
// }
// fmt.Printf(fmt.Sprint(money))

// }
package main

import "fmt"

// func Print() {
// 	fmt.Printf("1 2 3 4 5 6")
// }
// func sum(num int, num1 int) {
// 	fmt.Printf(fmt.Sprint(num + num1))
// }
// func sum(num int, num1 int) int {
// 	res := num + num1
// 	return res
// }
// func sum(num int, num1 int) (int, int, int, int) {
// 	sum := num + num1
// 	min := num - num1
// 	mul := num * num1
// 	dev := num / num1
// 	return sum, min, mul, dev
// }
// func squar(Function func(int) int) {
// 	fmt.Println(Function(25))
// }
func test(x string) func() {
	return func() {
		fmt.Println(x)
	}
}
func main() {
	// a := func() {
	// 	fmt.Println("1")
	// }
	// a()
	// res := func(x int) int {
	// 	return x * x
	// }
	// squar(res)
	// a := test("qfqf")
	// a()
	test("qfqf")()
}

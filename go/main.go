package main

import (
	"fmt"
	"os"
)

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
// func test(x string) func() {
// 	return func() {
// 		fmt.Println(x)
// 	}
// }
// func main() {
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
// test("qfqf")()
// }
// func change(str *string) {
// 	*str = "qwe"
// }
// func main() {
// 	s := "rty"
// 	fmt.Println(s)
// 	change(&s)
// 	fmt.Println(s)
// }
// type User struct {
// 	name    string
// 	age     int
// 	pasword string
// }

// func Print(u User) {
// 	fmt.Println(u)
// }
// func PrintName(u User) {
// 	fmt.Println(u.name)
// }
// func Switch(u *User, name string) {
// 	u.name = name
// }
// func main() {
// 	// var user User = User{"qfqf", 12, "qfqf"}
// 	user := User{"qfqf", 12, "qfqf"}
// 	fmt.Println(user)
// 	Print(user)
// 	PrintName(user)
// 	Switch(&user, "qqqq")
// 	PrintName(user)
// }
// type Number struct {
// 	num1 int
// 	num2 int
// }
// type NI interface {
// 	Sum()
// }

// func (n Number) Sum() {
// 	fmt.Println(n.num1 + n.num2)
// }

//	func main() {
//		var a NI
//		num := Number{5, 1}
//		a = num
//		a.Sum()
//	}
// func main() {
// 	data := []byte("Hello Lol")
// 	e := ioutil.WriteFile("text.txt", data, 0600)
// 	if e != nil {
// 		fmt.Println("Error")
// 	}
// 	file_data, err := ioutil.ReadFile("text.txt")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	fmt.Println(string(file_data))

//		// os.Remove("text.txt")
//	}
//
//	func main() {
//		ch := make(chan int)
//		ch2 := make(chan string)
//		go say("qwe", ch, ch2)
//		time.Sleep(2 * time.Second)
//		go say("qwe", ch, ch2)
//		time.Sleep(2 * time.Second)
//		num := <-ch
//		str := <-ch2
//		fmt.Println(num, str)
//	}
//
//	func say(str string, ch chan int, ch2 chan string) {
//		fmt.Println(str)
//		ch <- 4
//		ch2 <- "qwwe"
//	}
//
//	func main() {
//		ch := make(chan int)
//		go Print(ch)
//		for i := range ch {
//			fmt.Println(i)
//		}
//	}
func Print(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		ch <- i
	}
	close(ch)
}
func main() {
	file, err := os.Create("Text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := "Hello"
	file.WriteString(data)
	file1, _ := os.ReadFile("Text.txt")
	fmt.Println(string(file1))
	os.Remove("Text.txt")
}

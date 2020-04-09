package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"net/http"
)

// 8. Structures:
type Car struct {
	// color string
	// size string
	// price int
	color, size string
	price int
}

//// Structure methods: value reciever and pointer reciever
////// Value reciever
func (myCar Car) carColor () string {
	return "Mi car color is: " + myCar.color
}

////// Pointer reciever
func (myCar *Car) priceUp () {
	myCar.price++
}

// 9. Interfaces: list of methods that each structs implements
type Area interface {
	area() float64
}
////
type Circle struct {
	x, y, radius float64
}
////
type Rectangle struct {
	width, height float64
}
////
func (myCircle Circle) area() float64 {
	return math.Pi * myCircle.radius * myCircle.radius
}
////
func (myRectangle Rectangle) area() float64 {
	return myRectangle.width * myRectangle.height
}
////
func getArea (a Area) float64 {
	return a.area()
}

// 10. Web:
func index (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello there!</h1>")
}

func about (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About page</h1>")
}

func main() {

	// 1. Variables:
	fmt.Println("Hello, 世界")
	var test float64 = 1.4
	fmt.Printf("%f\n", test)
	fmt.Println(math.Floor(test))
	fmt.Println(math.Ceil(test))

	// 2. Arrays & Slices:
	myArray := [4] string {"array1", "array2", "array3"} // Array: fixed size
	fmt.Println(reflect.TypeOf(myArray))
	mySlice := [] string {"slice1", "slice2", "slice3"} // Slice: no fixed size
	fmt.Println(reflect.TypeOf(mySlice))
	fmt.Printf("%T\n", mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(mySlice[1:3])

	// 3. Conditionals:
	//// 3.1. if, else if, else:
	x := 5
	y := 3
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", y, x)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	//// 3.2. switch:
	color := "blue"
	switch color {
	case "blue":
		fmt.Println("Color is " + color)
	}

	// 4. Loops:
	for i:=1; i<=10; i++ {
		fmt.Println("Number " + strconv.Itoa(i))
	}
	// FizzBuzz game
	for i := 1; i < 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		}	else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 5. Maps: key-value pairs
	emails := make(map[string]string)
	// emails := map[string]string{}
	emails["Juan"] = "juan@mail.com"
	emails["Fabian"] = "fabian@mail.com"
	fmt.Println(emails)
	fmt.Println(emails["Juan"])
	fmt.Println(len(emails))
	delete(emails, "Juan")
	fmt.Println(emails)
	emails2 := map[string]string{"Juan":"juan@mail.com", "Fabian":"fabian@mail.com"}
	fmt.Println(emails2)

	// 6. Range: loop through arrays
	names := [] string {"name1", "name2", "name3"}
	for i, name := range names {
		fmt.Printf("Index: %d - Name: %s\n", i, name)
	}
	//// Not using index '_'
	for _, name := range names {
		fmt.Printf("Name: %s\n", name)
	}
	//// range in Maps
	for k, v := range emails2 {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
	//// range in Maps: only key
	for k := range emails2 {
		fmt.Printf("Key:. %s\n", k)
	}
	//// range in Maps: only value
	for _, v := range emails2 {
		fmt.Printf("Value: %s\n", v)
	}

	// 7. Pointers:
	a := 10
	b := &a // &a = memory address where value of 5 is stored
	fmt.Println("Value of a: " + strconv.Itoa(a))
	fmt.Print("Memory address of value 5: ")
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Printf("%T\n", b)
	//// Use * to read value from memory address
	fmt.Println(*b)
	fmt.Println(*&a)
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(*&b)
	fmt.Println(**&b)
	//// Change value with Pointers
	*b = 20
	fmt.Println(a)
	fmt.Println(b)

	// 8. Structures:
	car1 := Car {color:"Blue", size:"Big", price:100}
	car2 := Car {"Red", "Small", 300}
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Printf("%T\n", car1)
	fmt.Printf("%T\n", car2)
	fmt.Println(car1.color)
	fmt.Println(car2.price)
	car1.color = "White"
	fmt.Println(car1)
	fmt.Println(car1.carColor())
	fmt.Println(car1.price)
	car1.priceUp()
	fmt.Println(car1.price)

	// 9. Interfaces:
	circle1 := Circle{2, 3, 5}
	rectangle1 := Rectangle{5, 4}
	fmt.Printf("Circle area: %f\n", getArea(circle1))
	fmt.Printf("Circle area: %f\n", getArea(rectangle1))

	// 10. Web:
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting..")
	http.ListenAndServe(":3000", nil)

}

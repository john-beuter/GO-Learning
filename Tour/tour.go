package main

/**
Working through "A Tour of GO" and adding new code as I
progress through the program. Notes are in here as well as in Obsidian

**/

import (
	"fmt"
	"math"
)

type Person struct {
	name string
}

func savePerson(person *Person) { //If you modify a pointer to the book you modify the actuall address. Not a copy

}

// You cannot carriage return the braces :(
// Notice how the variable type is declared after the name
// Alternativly I could declare the function like this add(x, y int)
func add(x int, y int) int {
	return x + y
}

// Look how we can return mutliple values!
// We can return ANY number of values in GO
func swap(x, y string) (string, string) {
	return y, x
}

// Notice how in this if statement we can assign the variable value in the condition instead of having to do
// v := math.Pow(x,n)
// Then if v < lim we can do it all in one line
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

//Return values can be named and declared as variables at the top of the file
//Here we have x and y
//The return statement does not specify the values returneed and is called a naked retur

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 3

func main() {
	//fmt.Println("First GO program!!!")
	//fmt.Println("The time is ", time.Now()) //Notic how I have to use captials for Now and Print
	fmt.Println(add(i, j)) //Pretty cool use of a function in GO

	k := 3 //Inside a function you can set a variable like this
	//This can only be done in a function

	//Watch this
	//You can dynamically set the variable type!! Kinda like python
	var value, value2, value3 = true, "Hello", 1
	fmt.Println(k, value, value2, value3)

}

//Is this a comment? Yes, yes it is
//go run hello.go

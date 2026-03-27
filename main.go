package main

import "fmt"

func main() {
	// strings
	
    var nameOne string = "emy"
    var nameTwo = "blessing"
    var nameThree string
    nameFour := "peaches"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
//numbers int,uint,float32, float64
	var ageOne int = 20
    var ageTwo = 30
    ageThree := 40
	var ageFour uint = 54
	var ageFive float32 = 5.5

	fmt.Print("hello")
	fmt.Print("world")

	fmt.Print("hello \\n")
	fmt.Print("world \\n")

	fmt.Println("hello")
	fmt.Println("world")

    fmt.Println(ageOne, ageTwo, ageThree, ageFour, ageFive)

	//formatted string - %v = template littoral equvolent. 
	name := "Bob"
	age := 42

	fmt.Printf("my name is %v and I am %v years old", name, age)

	// use %q when you want quotes
	userName := "Bob"

	fmt.Printf("my name is %q", userName)

//%t for getting type of variable
fmt.Printf("This is a type of %t", age)

//arrays = [3] declares size of an array
//arrays can only have one type inside
var ages = [3]int{20,25,30}

fmt.Println(ages, len(ages))

//slices- like arrays but you don't have to specify size
var scores = []int{100,50,60}
//updates item at index 2
scores[2] = 25
scores = append(scores, 50)
fmt.Println(scores, len(scores))

//start at index 0 and continue to index 2
rangeOne := scores[0:3]
//slice index 2 through end
rangeTwo := scores[2:]
//slice from beginning up to but not including index 3
rangeThree := scores[:3]
fmt.Println(rangeOne)
fmt.Println(rangeTwo)
fmt.Println(rangeThree)

//loops
x := 0
for x < 5 {
	fmt.Println("the value of x is", x)
	x++
}

for i := 0; i <5; i++{
	fmt.Println("value of i is", i)
}

//len() same as array.length in js
var elves = []string{"feanor", "Elrond", "Thingel"}
for i := 0; i < len(elves); i++{
	fmt.Println(elves[i])
}

//range keyword is used to iterate over an array, slice,etc..
//with range you must use both index and value
for index, value := range elves {
		fmt.Printf("the position of %v is %v \\\\n", value, index)
	}

	//if you only want to use one use _
	for _, value := range elves {
		fmt.Printf("the value is %v \\\\n", value)
}
}
package main

import (
	 "fmt"
	//  "math" 
	 "time"
)


// this file is referring from the gobyexample.com
func main()	{
    satu := 1
	dua := 2
	word := "apple"

	fmt.Println("Satu: ", 1);
	fmt.Printf("Dua: %d %d %s \n", satu, dua, word ); 
	fmt.Println("Tiga Koma 5: ",3.5);
	fmt.Println("Filbert Christian Winch"[0:len("Filbert")]);
	fmt.Println(len("Filbert"));


	fmt.Println("\n=============================================");
	fmt.Println("Repetition");
	fmt.Printf("\n\n\n");

	/* for is Go's only looping construct. 
	Here are some basic types of for loops
	*/


	i := 1;
	for i <= 3{
		fmt.Println(i);
		i = i + 1;
	}

	for j := 0 ; j < 3;j++{
		fmt.Println(j)
	}

	for i := range 3{
		// another way of accomplishing the basic "do this n times, with default increment of 1"
		fmt.Println("range", i);
	}

	for {
		/* for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function */
		fmt.Println("loop");
		break
	}

	for n:= range 6{
		if n % 2 == 0{
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("\n==================================================")
	fmt.Println("selection");
	fmt.Printf("\n\n\n");

	if 7 % 2 == 0{
		fmt.Println("7 is even");
	} else {
		fmt.Println("7 is odd");
	}


	if 8 % 4 == 0 || 7 % 2 == 0{
		fmt.Println("either 8 or 7 are even");
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// There's no ternary if in Go, so you'll need to use 
	// a full if if statement even for basic conditions

	fmt.Println("\n==================================================")
	fmt.Println("Switch");
	fmt.Printf("\n\n\n");

	
	fmt.Print("write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one");
	case 2:
		fmt.Println("two");
	case 3:
		fmt.Println("three");
	}

	/* You can use commas to separate multiple expression
	 in the same case statement (which is nice) */

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend");
	default:
		fmt.Println("It's Weekday");
	}

	/* switch without an expression is an alternate way to express
	if/else logic. Case Expressions can be non-constants

	or in other words, u can just switch case normally as if u are using if/else
	but in specific cases (which is quite convenient for sure)
	*/

time := time.Now()
switch {
case time.Hour() < 12:
	fmt.Println("It's before noon");
default:
	fmt.Println("it's after noon");
}


    /* a type switch compares type instead of values.
	use this to discover type of an interface value. 
	In this example, the variable t will have the type 
	corresponding to its clause

	*/

whatAmI := func(i interface{}){
	switch time := i.(type) {
	case bool:
		fmt.Println("I'm A Bool");
	case int:
	    fmt.Println("I'm an Int");
	default:
		fmt.Printf("Don't know the type of %T\n", time)
}
}

whatAmI(true);
whatAmI(1);
whatAmI("hey");

fmt.Println("==================================================")
fmt.Println("Arrays");
fmt.Printf("\n\n\n");

  var names = []string{"Yu"} // this means that in the index 0 of names, you initialize the name yu
  fmt.Println(names)  
  }

// session 2 : Goroutines & Default
// here how's you can make a function with a return int bruh
func addNumber(a int,b int) int{
  return a+b;
}

// here how you can do it with loop also, this one returns void
func printNumbers(){
  fmt.Println("this one prints the default function")
  for i:=1 ; i <= 10 ; i++{
    fmt.Println(i)
    time.Sleep(time.Second) // some kind of run to await per second kinda thing, asynchronous thing
  }
  fmt.Println("end of print")
}

func printAddNumbers(){
  fmt.Println("\n\nthis one prints the one with the sum bruh")
  for i:= 6 ; i <= 10; i++{
    fmt.Println(addNumber(i,i))
    time.Sleep(time.Second) // some kind of run to await per second kinda thing, asynchronous thing
  }
}



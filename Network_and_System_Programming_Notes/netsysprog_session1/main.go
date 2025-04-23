package main


import (
  "fmt"
  "time"
)

func main(){
  fmt.Println("anak anjing");
  var t int = 2;
  fmt.Printf("ma dem dog %d\n", t); 
  f := 5;
  fmt.Printf("ma dem dog %d\n", f);


  //map : where there's a key and value of each, just like `json 
  var maps map[string]int = map[string]int{"count": 2, "Weight" : 10}
  fmt.Println(maps);

  // suppose we want to get a value from an attribute
  fmt.Println("count attribute value: ", maps["count"]);



  //struct : like c, schema, use pascal for this one, to differentiate
  type Person struct {
    name string // samain kaya mikir char name
    age int // samain kaya mikir int age
  }

  // Notes: don't confuses go with c, go is derived from c, remember?

  // struct definition implementation
  var person1 Person = Person{"Filbert", 18}
  fmt.Println(person1)


  //how to do the conditional in go 
  var test_input int;
  fmt.Print("scan the test_input variable: ");
  fmt.Scanf("%d\n", &test_input);

  if test_input == 5{
    fmt.Println("input = 5")
  } else if test_input == 5{
    fmt.Println("your input is greater than 5")
  }else{
    fmt.Println("broken")
  }


  // how to do the switch 
  switch test_input{
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    default:
      fmt.Println("default")
  }


  // session 2 : goroutines 
  // definition -> goroutines used to run code asynchronously
  // the "go" is the asynchronous


  // remember that asynchronous enables to run multithread so in this below case, they both will run
  // they requires that "awaits, in here we're gonna see if each functions are awaited
  // the expected result is that they will halt every seconds 
  // P.S: BECAUSE THEY USE TIME.SLEEP Per seconds
  go printNumbers()
  go printAddNumbers()


  // session 2 : defer
  // definition -> to run the code after all the code has already ran, LIFO -> Last In First Out
  // some how it's kinda await thingy
  fmt.Println("hallo")

  defer fmt.Println("saya") // with this the code's gonna be -> hallo, budi, saya, weird....

  fmt.Println("budi")

  defer fmt.Println("test") // with this the code's gonna be -> hallo, budi, test, saya, also weird....


  // so yep, what can we got from here ?
  // the defer use the LIFO concept 

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

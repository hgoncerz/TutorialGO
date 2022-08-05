package main

import "fmt"

func main(){
	age := 35
	name := "shaun"

	//Print
	fmt.Print("hello, ")
	fmt.Print("world! \n ")

	//Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is",name)

	//Printf (formated strings)
	fmt.Printf("my name is %v any my name is %v \n",age, name)
	fmt.Printf("my name is %q any my name is %q\n",age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %0.1f points \n", 225.55)
	//sprintf(save formatted strings)
	var str = fmt.Sprintf("my name is %v any my name is %v \n",age, name)
	fmt.Println("the saved formated string is:", str)
}
package main

import "fmt"

import "reflect"

func  main() {

	i := 1

	fmt.Println(i)
	
	fmt.Println(reflect.TypeOf(i))

	fmt.Println("First For loop")

	for i<=3 {
		fmt.Println("from inside loop " , i)
		i = i + 1
	}


	fmt.Println("Second For loop")

	for j :=7 ; j<=10 ; j++ {
		fmt.Println("from inside loop j = " , j)
	}

	for {

		fmt.Println("loop")
		break
	}
}
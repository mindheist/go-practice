package main

import "fmt"

func main(){

	m := make(map[string] int)

	// this looks like creating a map in which keys are 
	// strings and all the values are going to be integers

	m["k1"] = 7

	// python Dejavu !

	m["k2"] = 10

	fmt.Println("map:",m)

	// Accessing the values - just like python maps
	fmt.Println(m["k1"])

	fmt.Println("length :" , len(m))

	// pythonic - again

	delete(m,"k1")

	fmt.Println("map after deleting = " , m)

	// or we could create a dict and add the key:values

	n := map[string]int{"a":1 , "b":2 , "c" : 3}

	fmt.Println("map n ==" , n)

}
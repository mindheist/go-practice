package main

import "fmt"

func main(){

	s:= make ([]string , 3)

	fmt.Println("empty array: " , s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"


	fmt.Println("After adding a,b,c: " , s)

	fmt.Println("Index 2 of string = " ,s[2])

	fmt.Println("len of the string" , len(s))

	s = append(s,"d")

	s = append(s,"e")

	fmt.Println("after adding d and e",s)

	c:=make([]string , len(s))

	copy(c,s)

	fmt.Println("copy ==" , c)

	l1 := s[2:5]

	//python Dejavu

	fmt.Println("slice 1 = " , l1)

	l1 = s[:5]

	fmt.Println("slice 2 = " , l1)


}
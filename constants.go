package main 

import "fmt"
import "math"

const s string  = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000000

	fmt.Println ("n =" , n)

	const d = 3e20 /n

	fmt.Println("d=",d)

	var x int  = 1

	fmt.Println(x)

	x = x + 1

	fmt.Println(x)


	fmt.Println(math.Sin(n))



}
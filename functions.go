package main 
import "fmt"

func plus(a int, b int) int {
	return a+b 
}

func plusplus ( a,b,c int) int {

	return a+b+c
}

func main(){

	fmt.Println("First function ; 1+2 =", plus(1,2))

	fmt.Println("Second function ; 10 + 20 + 30 == " , plusplus(10,20,30))
}


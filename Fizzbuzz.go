package main


import "fmt"


func main(){

	var a int

	fmt.Scanf("%d", &a)

	if a%15==0 && a!=0{
		fmt.Println("FizzBuzz")
	}else if a%3==0 && a!=0 {
		fmt.Println("Fizz")
	}else if a%5==0 && a!=0 {
		fmt.Println("Buzz")
	}else {
		fmt.Println("No FizzBuzz")
	}

}

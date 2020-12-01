package main

import "fmt"

func main(){

	var a int
	i := 1
	sum := 0
	fmt.Scanf("%d", &a)

	for i<=a {
		if i%2==1 {

			sum += i
			i = i+2

		}


	}

	fmt.Println(sum)

}
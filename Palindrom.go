package main

import "fmt"

func main(){


	var a string
	i := 0
	check := 0
	fmt.Scanf("%s", &a)
	len := len(a)
	div := len/2
	arr := len-1
	for i<div {

		if a[i]!=a[arr]{
			check++
		}
		i++
		arr--
	}
	if check == 0 {
	
		fmt.Println("Palindrom")

	}else{
		fmt.Println("No palindrom")
	}

}



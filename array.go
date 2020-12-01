package main

import "fmt"

func main() {
	var a int
	fmt.Printf("Array qiymatlari soni: ")
	fmt.Scanf("%d", &a)
	marcuz := make([]string, a)
	i := 0
	for i<a {
		fmt.Scanf("%v", &marcuz[i])
		i++
	}

	i = 0
	fmt.Printf("Takrorlanishlar: ")
	for i<a {
		j := i+1 
		for j<a {
			if marcuz[i]==marcuz[j]{
				fmt.Print(marcuz[i])
			}
			j++
		}
		
		i++
	}
	fmt.Println("")

}
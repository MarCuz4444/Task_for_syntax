package main


import "fmt"

func main(){

var a int;
	fibonacci := []int{0,1}
	//b := 1
	i := 2
	fmt.Scanf("%d", &a)
	
	if(i<=a){
		
		for i<=a { 
			lenfib := len(fibonacci)
			first := fibonacci[lenfib-1]
			second := fibonacci[lenfib-2]
			newone := first+second
			fibonacci = append(fibonacci, newone)
			i++
		}

			fmt.Println(fibonacci)
		}else{
		if(a==0){

			fmt.Println(fibonacci[0])

		}else{

			fmt.Println(fibonacci)

		}
	}

}
package main

import "fmt"

var f1 = 0
func fib() func() int{
	f2 := 1
	return func() int{
		temp := f2
		f2 = f2+f1
		f1 = temp
		return f1
	}
}

func main(){
	f := fib()
	sum := 0
	for f1<4000000{
		term := f()
		if term%2==0{
			sum+=term
		}
	}
	fmt.Printf("%d\n",sum)
}

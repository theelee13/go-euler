package main

import "fmt"

func main(){
	second := Prime{1,nil}
	first := Prime{0,&second}
	last := &second
	for i:=2;i<10;i++{
		newLast := Prime{i,nil}
		last.next = &newLast
		last = &newLast
	}
	for member := &first;member!=nil;member=member.next{
		fmt.Println((*member).val)
	}
}

type Prime struct{
	val int
	next *Prime
}

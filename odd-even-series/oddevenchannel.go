package main

import "fmt"

func generateOddNumbers(n int, oddChannel chan<- int){
	for i:=2; i <= n; i+=2{
		oddChannel<-i
	}
	close(oddChannel)
}

func generateEvenNumbers(n int, evenChannel chan<- int){
	for i:=1; i<= n; i+=2{
		evenChannel<-i
	}
	close(evenChannel)
}

func printOddEvenSeries(n int){

	oddChannel := make(chan int)
	evenChannel := make(chan int)

	go generateOddNumbers(n, oddChannel)
	go generateEvenNumbers(n, evenChannel)

	for i := range oddChannel{
		fmt.Printf("%d ",i)
	}

	for i := range evenChannel{
		fmt.Printf("%d ",i)
	}
}
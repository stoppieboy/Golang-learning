package main

import(
	"fmt"
	"Errors/greetings"
	"log"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message,err := greetings.Hello("Shivam")
	if err != nil{
		log.Fatal(err)
	}
	
	fmt.Println(message)
}
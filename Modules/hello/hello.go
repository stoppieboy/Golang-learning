package main

import (
    "fmt"
    "Modules/greetings"
)

func main(){
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}

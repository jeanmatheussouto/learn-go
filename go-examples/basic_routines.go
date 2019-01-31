package main

import(
	"fmt"
	"time"
)

func sayHello(msg string, delay time.Duration){
	time.Sleep(delay)
	fmt.Println(msg)
}

func main(){
	go sayHello("hey", 1*time.Second)
	go sayHello("ho", 2*time.Second)
	go sayHello("let's go!!", 3*time.Second)
	time.Sleep(3*time.Second)
}


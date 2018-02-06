package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	inputReader:= bufio.NewReader(os.Stdin)
	fmt.Println("please you name:")
	input,err:=inputReader.ReadString('\n')
	if err!=nil{
		fmt.Printf("An error occurred %s",err)
		os.Exit(0)
	}

	name:=input[:len(input)-1]
	fmt.Printf("hello %s, what can I do help your",name)

	for{
		input,err =inputReader.ReadString('\n')
		if err!=nil{
			fmt.Printf("An error occurred %s",err)
			continue
		}
		input =input[:len(input)-1]
		input =  strings.ToLower(input)
		switch input {
		case "":
			continue
		case "bye","nothing":
			fmt.Println("bye")
			os.Exit(0)
		default:
			fmt.Printf("sorry, I don,t catch you ")
		}
	}

}

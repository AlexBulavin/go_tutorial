package main

import "fmt" //Импортируем пакет fmt - format

func main() {
	var hello = "Hello world"
	fmt.Println(hello)

	outString := fmt.Sprintf("our string: %T", hello)
	fmt.Println(outString)
	fmt.Println()
}

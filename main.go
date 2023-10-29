package main

import "fmt" //Импортируем пакет fmt - format

func main() {
	var hello string
	fmt.Print("hello = ")
	fmt.Println(hello)

	hello = "Hello world"
	fmt.Println(hello)

	outString := fmt.Sprintf("our string: %T", hello) //Вывели тип переменной hello
	fmt.Println(outString)
	fmt.Println()
}

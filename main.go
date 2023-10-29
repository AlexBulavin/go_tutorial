package main

import (
	"fmt"
	"unsafe"
) //Импортируем пакет fmt - format

func main() {
	var hello string
	fmt.Print("hello = ")
	fmt.Println(hello)

	hello = "Hello world"
	fmt.Println(hello)

	outString := fmt.Sprintf("our string: %T", hello) //Вывели тип переменной hello
	fmt.Println(outString)

	fmt.Printf("Type: %T Value %v\n", hello, hello) //Выводим тип и значение переменной

	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof(uint8(1)))
	fmt.Println(unsafe.Sizeof(int8(1)))
	fmt.Println(unsafe.Sizeof(int32(1)))
	fmt.Println(unsafe.Sizeof("hello"))
	fmt.Println(unsafe.Sizeof(true))

}

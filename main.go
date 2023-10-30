package main

import (
	"fmt"
	"math"
	"unsafe"
) //Импортируем пакет fmt - format и unsafe

var name string //Переменная объявлена на уровне пакета. Она будет видна из любой функции внутри пакета

const name1 bool = true

func main() {

	var hello string
	fmt.Print("hello = ")
	fmt.Println(hello)

	name := "Hi there" //Объявили и инициализировали локальную переменную name. И теперь внутри блока используется именно она

	fmt.Println(name)
	name = "123456"
	fmt.Println(name)

	hello = "Hello world"
	fmt.Println(hello)

	outString := fmt.Sprintf("our string: %T", hello) //Вывели тип переменной hello
	fmt.Println(outString)

	fmt.Printf("Type: %T Value %v\n", hello, hello) //Выводим тип и значение переменной

	fmt.Println(unsafe.Sizeof(1)) //Выводм размер аргумента в байтах
	fmt.Println(unsafe.Sizeof(uint8(1)))
	fmt.Println(unsafe.Sizeof(int8(1)))
	fmt.Println(unsafe.Sizeof(int32(1)))
	fmt.Println(unsafe.Sizeof("hello"))
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(name1))
	fmt.Println(name1)
	test1()
	test2()
	fmt.Println(SumMultAB(5, 100))
	fmt.Println(powerAndLog(10.0, 2.0))

	var mutliplier func(x, y int) int
	mutliplier = func(x, y int) int {
		return x * y
	}
	fmt.Println(mutliplier(10, 20))

	x, y := 8, 15
	divider := func(x, y int) float32 {
		if y != 0 {
			return float32(x) / float32(y)
		} else {
			return float32(0.0)
		}
	}
	fmt.Print("divider(x, y): ")
	fmt.Println(divider(x, y))

	first, second := 85, 12

	sumFunc := func(n, m int) int {
		return n + m
	}
	fmt.Printf("calculate(%v, %v, sumFunc) = %v\n", first, second, calculate(first, second, sumFunc))

	subtractFunc := func(n, m int) int {
		return n - m
	}
	fmt.Printf("calculate(%v, %v, subtractFunc) = %v\n", first, second, calculate(first, second, subtractFunc))

	divideBy2 := createDivider(2)
	divideBy18 := createDivider(18)

	fmt.Printf("divideBy2 :%v\n", divideBy2(100))
	fmt.Printf("divideBy18 :%v\n", divideBy18(180))

	var intPointer *int
	fmt.Printf("intPointer Type: %T Value: %#v \n", intPointer, intPointer)

	var temp int64 = 7
	fmt.Printf("temp type: %T value %#v \n", temp, temp)

	//Операция получения указателя:
	var pointerA *int64 = &temp
	fmt.Printf("pointerA Type: %T %#v %#v\n", pointerA, pointerA, *pointerA)

	type OurString string
	type OurInt int

	type Person struct {
		Name string
		Age  int
	}

	var customString OurString

	fmt.Printf("customString %T %v \n", customString, customString)

	customString = "Hello Go developers"

	fmt.Printf("customString %T %v \n", customString, customString)

	customInt := OurInt(5)

	fmt.Printf("customInt %T %v \n", customInt, customInt)
	fmt.Println(int(customInt))

	interfaceValues()

	//end of main()
}

//Урок 13 Интерфейсы
type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бежит", h.Name)
}

func (h Human) Swim() string {
	return fmt.Sprintf("Человек %s плавает", h.Name)
}

func (h Human) Fly() string {
	return fmt.Sprintf("Человек %s летает", h.Name)
}

func (h Human) writeCode() {
	fmt.Printf("Человек %s пишет код", h.Name)
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("runner Type: %T Value: %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("runner is nil")
	}

	var unnamedRunner *Human
	fmt.Printf("unnamedRunner Type: %T Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("runner = unnamedRunner Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}

	nameRunner := &Human{Name: "John"}
	fmt.Printf("nameRunner Type: %T Value: %#v\n", nameRunner, nameRunner)

	runner = nameRunner
	fmt.Printf("runner = nameRunner Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}

}

func createDivider(divider int /*2*/) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider /*2*/
	}
	return dividerFunc
}

func test1() {
	name = "test1"
	fmt.Println(name)
}

func test2() {
	var num1 uint = 15
	num2 := 22
	fmt.Println(num1 + uint(num2)) //Привели тип int к uint
}

func SumMultAB(a, b int) (int, int) {
	return a + b, a * b
}

func powerAndLog(a, b float64) (power float64, logarithm float64) {
	power = math.Pow(a, b)
	logarithm = math.Log(a)

	return
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

package main

import (
	"GoDesignPattern/Singleton/unsafe"
	"fmt"
)

func main() {
	obj1 := unsafe.GetInstance()
	obj1.Name = "obj1"
	obj2 := unsafe.GetInstance()
	obj2.Name = "obj2"
	fmt.Println("obj1: ", obj1, " &obj1:", &obj1.Name)
	fmt.Println("obj2: ", obj2, " &obj2:", &obj2.Name)
	fmt.Printf("%p %T\n", obj1, obj1)
	fmt.Printf("%p %T\n", obj2, obj2)
}

package main

import (
	"GoDesignPattern/Factory/pkg"
	"fmt"
)

func main()  {
	payment,err := pkg.GetPaymentMethod(pkg.Cash)
	if err != nil{
		fmt.Printf("A payment method of type 'Cash' must Exist")
	}
	msg := payment.Pay(8.9)
	fmt.Println(msg)
}
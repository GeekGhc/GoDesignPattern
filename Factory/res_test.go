package main

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T)  {
	payment,err := GetPaymentMethod(Cash)
	if err != nil{
		fmt.Printf("A payment method of type 'Cash' must Exist")
	}
	msg := payment.Pay(8.9)
	fmt.Println(msg)
}
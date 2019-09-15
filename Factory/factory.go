package main

import (
	"errors"
	"fmt"
)

const (
	Cash      = 1
	DebitCard = 2
)
// 抽象公共接口
type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPM struct {
}
type DebitCardPM struct {
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("payment method %d not defined!", m))
	}
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f piad using debit card", amount)
}

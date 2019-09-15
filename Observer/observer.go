package Observer

import (
	"fmt"
)

// 报社-客户
type Customer interface {
	update()
}

type CustomerA struct {
}

func (*CustomerA) update() {
	fmt.Println("客户A,收到了报纸...")
}

type CustomerB struct {
}

func (*CustomerB) update() {
	fmt.Println("客户B,收到了报纸...")
}

// 报社(观察者)
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}

func (n *NewsOffice) newspaperCome() {
	// 通知所有客户
	n.notifyAllCustomer()
}

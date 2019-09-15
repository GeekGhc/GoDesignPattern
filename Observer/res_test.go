package Observer

import "testing"

func TestObserver(t *testing.T){
	customerA := &CustomerA{}
	customerB := &CustomerB{}

	office := &NewsOffice{}
	// 客户订阅
	office.addCustomer(customerA)
	office.addCustomer(customerB)
	// 消息通知
	office.newspaperCome()
}
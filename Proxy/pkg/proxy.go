package pkg

import "fmt"

type Seller interface {
	Sell(name string)
}

// 火车站
type Station struct {
	Stock int // 库存
}

func (station *Station) Sell(name string) {
	if station.Stock > 0 {
		station.Stock--
		fmt.Printf("代理点中： %s 买了一张票，剩余%d\n", name, station.Stock)
	} else {
		fmt.Println("票已经售空")
	}
}

// 火车代理点
type StationProxy struct {
	Station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) Sell(name string) {
	if proxy.Station.Stock > 0 {
		proxy.Station.Stock--
		fmt.Printf("代理点： %s买了一张票，剩余: %d\n", name, proxy.Station.Stock)
	} else {
		fmt.Println("票已经售空")
	}
}

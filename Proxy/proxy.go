package main

import "fmt"

type Seller interface {
	Sell(name string)
}

// 火车站
type Station struct {
	stock int // 库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中： %s 买了一张票，剩余%d\n", name, station.stock)
	} else {
		fmt.Println("票已经售空")
	}
}

// 火车代理点
type StationProxy struct {
	Station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	if proxy.Station.stock > 0 {
		proxy.Station.stock--
		fmt.Printf("代理点： %s买了一张票，剩余: %d\n", name, proxy.Station.stock)
	} else {
		fmt.Println("票已经售空")
	}
}

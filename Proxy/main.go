package main

import (
	"GoDesignPattern/Proxy/pkg"
)

func main()  {
	station := &pkg.Station{3}
	proxy := &pkg.StationProxy{station}

	station.Sell("张三")
	proxy.Sell("小明")
	proxy.Sell("小红")
	proxy.Sell("小张")
}

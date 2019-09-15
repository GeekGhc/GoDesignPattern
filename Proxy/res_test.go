package main

import "testing"

func TestProxy(t *testing.T)  {
	station := &Station{3}
	proxy := &StationProxy{station}

	station.sell("张三")
	proxy.sell("小明")
	proxy.sell("小红")
	proxy.sell("小张")
}

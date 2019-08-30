package main

// 建立私有变量
var instance *singleton

// 结构体替代
type singleton struct{
	Name string
}

// 获取实例方法
func getInstance() *singleton{
	if instance == nil{
		instance = new(singleton)
	}
	return instance
}
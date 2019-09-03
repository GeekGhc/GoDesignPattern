package safe

import "sync"

// 建立私有变量
var instance *singleton
var once sync.Once

// 单例结构体Ò
type singleton struct {
	Name string
}

func GetInstance() *singleton{
	//只会执行一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
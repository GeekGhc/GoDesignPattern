package unsafe
// 此写法非线程安全

// 建立私有变量
var instance *singleton

// 单例结构体
type singleton struct{
	Name string
}

// 获取实例方法
func GetInstance() *singleton{
	if instance == nil{
		instance = new(singleton)
	}
	return instance
}
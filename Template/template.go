package Template

import "fmt"

type Cooker interface {
	open()
	fire()
	cooke()
	offFire()
	close()
}

//定义一个类似于抽象类
type CookMenu struct {
}

func (*CookMenu) open() {
	fmt.Println("open...")
}

func (*CookMenu) fire() {
	fmt.Println("fire...")
}

// 做菜则由具体的子类实现
func (*CookMenu) cooke() {

}

func (*CookMenu) offFire() {
	fmt.Println("offFire...")
}

func (*CookMenu) close() {
	fmt.Println("close...")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.offFire()
	cook.close()
}

type Tomato struct {
	CookMenu
}

func (*Tomato) cooke() {
	fmt.Println("西红柿...")
}

type Potato struct {
	CookMenu
}

func (*Potato) cooke() {
	fmt.Println("土豆...")
}

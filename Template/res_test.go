package Template

import "testing"

func TestTemplate(t *testing.T){
	// 做西红柿
	tomato := &Tomato{}
	doCook(tomato)

	//做土豆
	potato := &Potato{}
	doCook(potato)
}

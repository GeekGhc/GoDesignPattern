package Strategy

// 实现接口，则为一个策略
type IStrategy interface {
	do(int, int) int
}

// 加
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct{}

func (*reduce) do(a,b int) int{
	return a-b
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy){
	operator.strategy = strategy
}

// 调用策略的方法
func (operator *Operator) calculate(a,b int)int{
	return operator.strategy.do(a,b)
}
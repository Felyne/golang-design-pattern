package factorymethod

var operatorFactoryMap = map[string]OperatorFactory{
	"plus":  PlusOperatorFactory{},
	"minus": MinusOperatorFactory{},
}

// 工厂类的工厂+单例模式
func NewOperatorFactory(ot string) OperatorFactory {
	return operatorFactoryMap[ot]
}

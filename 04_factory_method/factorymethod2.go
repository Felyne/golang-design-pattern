package factorymethod

var operatorFactoryMap = map[string]OperatorFactory{
	"plus":  PlusOperatorFactory{},
	"minus": MinusOperatorFactory{},
}

// 工厂类的的简单工厂
func NewOperatorFactory(ot string) OperatorFactory {
	return operatorFactoryMap[ot]
}

package adapter

//
////Target 是适配的目标接口
//type Target interface {
//	Request() string
//}
//
////Adaptee 是被适配的目标接口
//type Adaptee interface {
//	SpecificRequest() string
//}
//
////NewAdaptee 是被适配接口的工厂函数
//func NewAdaptee() Adaptee {
//	return &adapteeImpl{}
//}
//
////AdapteeImpl 是被适配的目标类
//type adapteeImpl struct{}
//
////SpecificRequest 是目标类的一个方法
//func (*adapteeImpl) SpecificRequest() string {
//	return "adaptee method"
//}
//
////NewAdapter 是Adapter的工厂函数
//func NewAdapter(adaptee Adaptee) Target {
//	return &adapter{
//		Adaptee: adaptee,
//	}
//}
//
////Adapter 是转换Adaptee为Target接口的适配器
//type adapter struct {
//	Adaptee
//}
//
////Request 实现Target接口
//func (a *adapter) Request() string {
//	return a.SpecificRequest()
//}

type ITarget interface {
	F1()
	F2()
	Fc()
}

type Adaptee struct {
}

func (a *Adaptee) Fa() {

}

func (a *Adaptee) Fb() {

}

func (a *Adaptee) Fc() {
}

// 类适配器，基于继承
type Adaptor struct {
	*Adaptee
}

func (a *Adaptor) F1() {
	a.Adaptee.Fa()
}

func (a *Adaptor) F2() {
	//重新实现f2()...
}

// Fc()不需要实现，直接继承自Adaptee，这是跟对象适配器最大的不同

// 对象适配器，基于组合
type Adaptor2 struct {
	adaptee *Adaptee
}

func (a *Adaptor2) F1() {
	a.adaptee.Fa()
}

func (a *Adaptor2) F2() {
	//重新实现f2()...
}

func (a *Adaptor2) Fc() {
	a.adaptee.Fc()
}

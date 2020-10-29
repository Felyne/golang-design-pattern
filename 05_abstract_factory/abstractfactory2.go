package abstractfactory

import "fmt"

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParer() ISystemConfigParser
}

type IRuleConfigParser interface {
	ParseRule(rule string)
}

type ISystemConfigParser interface {
	ParseSystem(osType string)
}

// 一个工厂类创建两种不同类型的对象
type JsonConfigParserFactory struct {
}

func (JsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return JsonRuleConfigParser{}
}

func (JsonConfigParserFactory) CreateSystemParer() ISystemConfigParser {
	return JsonSystemConfigParser{}
}

type JsonRuleConfigParser struct {
}

func (JsonRuleConfigParser) ParseRule(rule string) {
	fmt.Println(rule)
}

type JsonSystemConfigParser struct {
}

func (JsonSystemConfigParser) ParseSystem(osType string) {
	fmt.Println(osType)
}

// 一个工厂类创建两种不同类型的对象
type XmlConfigParserFactory struct {
}

func (XmlConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return XmlRuleConfigParser{}
}

func (XmlConfigParserFactory) CreateSystemParer() ISystemConfigParser {
	return XmlSystemConfigParser{}
}

type XmlRuleConfigParser struct {
}

func (XmlRuleConfigParser) ParseRule(rule string) {
	fmt.Println(rule)
}

type XmlSystemConfigParser struct {
}

func (XmlSystemConfigParser) ParseSystem(osType string) {
	fmt.Println(osType)
}

package strategy

import (
	"time"
)

func ExamplePay() {
	// 运行时动态确定策略
	strategy := NewPaymentStrategy(getRandPaymentType())
	payment := NewPayment("Ada", "", 123, strategy)
	payment.Pay()
}

func getRandPaymentType() PaymentType {
	if time.Now().Unix()%2 == 0 {
		return PaymentTypeCash
	}
	return PaymentTypeBank
}

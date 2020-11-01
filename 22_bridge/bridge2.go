package bridge

import "fmt"

type NotificationEmergencyLevel string

const (
	SEVERE  NotificationEmergencyLevel = "SEVERE"
	URGENCY NotificationEmergencyLevel = "URGENCY"
	NORMAL  NotificationEmergencyLevel = "NORMAL"
	TRIVIAL NotificationEmergencyLevel = "TRIVIAL"
)

type Notification interface {
	Notify(msg string)
}

type MsgSender interface {
	Send(msg string)
}

type TelephoneMsgSender struct {
	telephone string
}

func (t TelephoneMsgSender) Send(msg string) {
	fmt.Printf("send msg to telephone %s: %s\n", t.telephone, msg)
}

type EmailMsgSender struct {
	email string
}

func (e EmailMsgSender) Send(msg string) {
	fmt.Printf("send msg to eamil %s: %s\n", e.email, msg)
}

type UrgencyNotification struct {
	msgSender MsgSender
}

func (u UrgencyNotification) Notify(msg string) {
	u.msgSender.Send(fmt.Sprintf("[Urgency] %s", msg))

}

type NormalNotification struct {
	msgSender MsgSender
}

func (n NormalNotification) Notify(msg string) {
	n.msgSender.Send(fmt.Sprintf("[Normal] %s", msg))
}

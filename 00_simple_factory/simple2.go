package simplefactory

import "fmt"

type Speaker interface {
	Speak(s string)
}

type chineseSpeaker struct {
}

func (chineseSpeaker) Speak(s string) {
	fmt.Println("你们好:", s)
}

type englishSpeaker struct {
}

func (englishSpeaker) Speak(s string) {
	fmt.Println("hello:", s)
}

var (
	chinese = chineseSpeaker{}
	english = englishSpeaker{}
)

type SpeakerType string

const (
	Chinese SpeakerType = "chinese"
	English SpeakerType = "english"
)

// 类似简单工厂+单例模式
func NewSpeaker(speakerType SpeakerType) Speaker {
	switch speakerType {
	case Chinese:
		return chinese
	case English:
		return english
	default:
		return english
	}
}

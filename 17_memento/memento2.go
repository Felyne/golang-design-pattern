package memento

import (
	"strings"
)

type InputText struct {
	text strings.Builder
}

func NewInputText() *InputText {
	return &InputText{}
}

func (i *InputText) GetText() string {
	return i.text.String()
}

func (i *InputText) Append(input string) {
	i.text.WriteString(input)
}

func (i *InputText) CreateSnapshot() *Snapshot {
	return NewSnapshot(i.GetText())
}

func (i *InputText) RestoreSnapshot(snapshot *Snapshot) {
	i.text.Reset()
	i.text.WriteString(snapshot.GetText())
}

type Snapshot struct {
	text string
}

func NewSnapshot(text string) *Snapshot {
	return &Snapshot{
		text: text,
	}
}

func (s *Snapshot) GetText() string {
	return s.text
}

type SnapshotHolder struct {
	stack []*Snapshot
}

func NewSnapshotHolder() *SnapshotHolder {
	return &SnapshotHolder{}
}

func (s *SnapshotHolder) PushSnapshot(item *Snapshot) {
	s.stack = append(s.stack, item)
}

func (s *SnapshotHolder) PopSnapshot() *Snapshot {
	if len(s.stack) > 0 {
		top := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}
	return nil
}

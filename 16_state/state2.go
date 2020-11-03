package state

type State string

const (
	MarioStateSmall State = "Small"
	MarioStateSuper State = "Super"
	MarioStateFire  State = "Fire"
	MarioStateCape  State = "Cape"
)

// 状态接口
type Imario interface {
	GetName() State
	ObtainMushRoom(stateMachine *MarioStateMachine)
	ObtainCape(stateMachine *MarioStateMachine)
	ObtainFireFlower(stateMachine *MarioStateMachine)
	MeetMonster(stateMachine *MarioStateMachine)
}

// 状态类不包含任何成员变量,可以设计成单例
var (
	smallMario Imario = newSmallMario()
	superMario Imario = newSuperMario()
	capeMario  Imario = newCapeMario()
	fireMario  Imario = newFireMario()
)

type SmallMario struct {
}

func newSmallMario() *SmallMario {
	return &SmallMario{}
}

func (s *SmallMario) GetName() State {
	return MarioStateSmall
}

func (s *SmallMario) ObtainMushRoom(stateMachine *MarioStateMachine) {
	stateMachine.SetCurrentState(superMario)
	stateMachine.SetScore(stateMachine.GetScore() + 100)
}

func (s *SmallMario) ObtainCape(stateMachine *MarioStateMachine) {
	stateMachine.SetCurrentState(capeMario)
	stateMachine.SetScore(stateMachine.GetScore() + 200)
}

func (s *SmallMario) ObtainFireFlower(stateMachine *MarioStateMachine) {
	stateMachine.SetCurrentState(fireMario)
	stateMachine.SetScore(stateMachine.GetScore() + 300)
}

func (s *SmallMario) MeetMonster(stateMachine *MarioStateMachine) {
	// do nothing...
}

type SuperMario struct {
}

func newSuperMario() *SuperMario {
	return &SuperMario{}
}

func (s *SuperMario) GetName() State {
	return MarioStateSuper
}

func (s *SuperMario) ObtainMushRoom(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (s *SuperMario) ObtainCape(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (s *SuperMario) ObtainFireFlower(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (s *SuperMario) MeetMonster(stateMachine *MarioStateMachine) {
	panic("implement me")
}

type CapeMario struct {
}

func newCapeMario() *CapeMario {
	return &CapeMario{}
}

func (c *CapeMario) GetName() State {
	return MarioStateCape
}

func (c *CapeMario) ObtainMushRoom(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (c *CapeMario) ObtainCape(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (c *CapeMario) ObtainFireFlower(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (c *CapeMario) MeetMonster(stateMachine *MarioStateMachine) {
	panic("implement me")
}

type FireMario struct {
}

func newFireMario() *FireMario {
	return &FireMario{}
}

func (f *FireMario) GetName() State {
	return MarioStateFire
}

func (f *FireMario) ObtainMushRoom(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (f *FireMario) ObtainCape(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (f *FireMario) ObtainFireFlower(stateMachine *MarioStateMachine) {
	panic("implement me")
}

func (f *FireMario) MeetMonster(stateMachine *MarioStateMachine) {
	panic("implement me")
}

// 状态机
type MarioStateMachine struct {
	score        int
	currentState Imario
}

func NewMarioStateMachine() *MarioStateMachine {
	return &MarioStateMachine{
		score:        0,
		currentState: smallMario,
	}
}

func (m *MarioStateMachine) ObtainMushRoom() {
	m.currentState.ObtainMushRoom(m)
}

func (m *MarioStateMachine) ObtainCape() {
	m.currentState.ObtainCape(m)
}

func (m *MarioStateMachine) ObtainFireFlower() {
	m.currentState.ObtainFireFlower(m)
}

func (m *MarioStateMachine) MeetMonster() {
	m.currentState.MeetMonster(m)
}

func (m *MarioStateMachine) GetScore() int {
	return m.score
}

func (m *MarioStateMachine) GetCurrentState() State {
	return m.currentState.GetName()
}

func (m *MarioStateMachine) SetScore(score int) {
	m.score = score
}

func (m *MarioStateMachine) SetCurrentState(currentState Imario) {
	m.currentState = currentState
}

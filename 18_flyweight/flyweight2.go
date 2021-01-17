package flyweight

type Color string

const (
	ColorRed   Color = "red"
	ColorBlack Color = "black"
)

type ChessPieceUnit struct {
	id    int
	text  string
	color Color
}

func NewChessPieceUnit(id int, text string, color Color) *ChessPieceUnit {
	return &ChessPieceUnit{
		id:    id,
		text:  text,
		color: color,
	}
}

func GetChessPieceUnit(id int) *ChessPieceUnit {
	return chessPieceUnitFactory.GetChessPieceUnit(id)
}

var chessPieceUnitFactory = newChessPieceUnitFactory()

type ChessPieceUnitFactory struct {
	pieces map[int]*ChessPieceUnit
}

func newChessPieceUnitFactory() *ChessPieceUnitFactory {
	factory := &ChessPieceUnitFactory{
		pieces: make(map[int]*ChessPieceUnit),
	}
	factory.init()
	return factory
}

func (c *ChessPieceUnitFactory) init() {
	c.pieces[1] = NewChessPieceUnit(1, "车", ColorRed)
	c.pieces[2] = NewChessPieceUnit(2, "马", ColorBlack)
	//...
}

func (c *ChessPieceUnitFactory) GetChessPieceUnit(id int) *ChessPieceUnit {
	return c.pieces[id]
}

// 棋子
type ChessPiece struct {
	*ChessPieceUnit
	positionX int
	positionY int
}

func NewChessPiece(c *ChessPieceUnit, x, y int) *ChessPiece {
	return &ChessPiece{
		ChessPieceUnit: c,
		positionX:      x,
		positionY:      y,
	}
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{}
	board.init()
	return board
}

func (c *ChessBoard) init() {
	c.chessPieces[1] = NewChessPiece(GetChessPieceUnit(1), 0, 0)
	c.chessPieces[2] = NewChessPiece(GetChessPieceUnit(2), 1, 0)
	//...
}

func (c *ChessBoard) Move(chessPieceId int, x, y int) {
	//...
}

package jungle

type Move struct {
	Turn  Turn
	Steps int
}

var ClockwiseTurn Move = Move{Turn: CW}
var CounterclockwiseTurn Move = Move{Turn: CCW}

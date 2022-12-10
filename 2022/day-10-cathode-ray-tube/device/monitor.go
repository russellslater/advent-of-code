package device

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

type CRTMonitor struct {
	spritePosition int
	drawPosition   int
	currCRTRow     []bool
	crtRows        [][]bool
}

func NewCRTMonitor() *CRTMonitor {
	return &CRTMonitor{
		spritePosition: 1,
		drawPosition:   0,
		currCRTRow:     make([]bool, 40),
		crtRows:        make([][]bool, 0),
	}
}

func (m *CRTMonitor) SetSpritePosition(value int) {
	m.spritePosition = value
}

func (m *CRTMonitor) DrawPixel() {
	m.currCRTRow[m.drawPosition] = (util.Abs(m.drawPosition-m.spritePosition) <= 1)
	m.drawPosition++

	if m.drawPosition%40 == 0 {
		m.drawRow()
	}
}

func (m *CRTMonitor) drawRow() {
	m.crtRows = append(m.crtRows, m.currCRTRow)
	m.currCRTRow = make([]bool, 40)
	m.drawPosition = 0
}

func (m *CRTMonitor) PrintCrtRows() {
	for _, row := range m.crtRows {
		for _, pixel := range row {
			if pixel {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

package cheeky

type Troop []*Monkey

func (t Troop) LowestCommonMultiple() int {
	lcm := 1
	for _, monkey := range t {
		lcm *= monkey.Divisor
	}
	return lcm
}

type Monkey struct {
	InspectionCount int
	Divisor         int
	Operation       func(int) int
	items           []int
	receivers       []int
}

func NewMonkey() *Monkey {
	return &Monkey{
		items:     []int{},
		receivers: []int{},
	}
}

func (m *Monkey) AddReceiver(ordinal int) {
	m.receivers = append(m.receivers, ordinal)
}

func (m *Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) InspectItems(troop Troop, worryLevel func(int) int) {
	for len(m.items) > 0 {
		m.SendItem(troop, worryLevel)
	}
}

func (m *Monkey) SendItem(troop Troop, worryLevel func(int) int) {
	if len(m.items) == 0 {
		return
	}
	item := m.items[0]
	m.items = m.items[1:]

	m.InspectionCount++

	value := worryLevel(m.Operation(item))

	if value%m.Divisor == 0 {
		troop[m.receivers[0]].AddItem(value)
	} else {
		troop[m.receivers[1]].AddItem(value)
	}
}

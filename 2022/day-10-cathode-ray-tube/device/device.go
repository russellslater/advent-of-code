package device

type Device struct {
	registerX int
	monitor   *CRTMonitor
}

func NewDevice() *Device {
	return &Device{
		registerX: 1,
		monitor:   NewCRTMonitor(),
	}
}

func (d *Device) add(value int) {
	d.registerX += value
	d.monitor.SetSpritePosition(d.registerX)
}

func (d *Device) signalStrength(cycle int) int {
	return d.registerX * cycle
}

func (d *Device) Run(instructions []*Instruction) int {
	sum := 0
	cycle := 0

	for len(instructions) > 0 {
		ins := instructions[0]
		instructions = instructions[1:]
		for i := 0; i < ins.Cycles; i++ {
			cycle++
			d.monitor.DrawPixel()

			if (cycle-20)%40 == 0 {
				sum += d.signalStrength(cycle)
			}
		}
		d.add(ins.AddValue)
	}

	return sum
}

func (d *Device) Display() {
	d.monitor.PrintCrtRows()
}

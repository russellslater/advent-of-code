package device

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Device struct {
	DetectionLimit int
}

func (d Device) Detect(sig string) int {
	last := make([]rune, 0, d.DetectionLimit)
	for i, r := range sig {
		last = append(last, r)
		fmt.Println(last, len(last), last[1:])
		if len(last) > d.DetectionLimit {
			last = last[1:]
			if util.Unique(last) {
				return i + 1
			}
		}
	}
	return -1
}

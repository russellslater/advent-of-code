package tetris

type Shape [][]rune

func (s Shape) Width() int {
	return len(s[0])
}

func (s Shape) Height() int {
	return len(s)
}

var (
	// 	####
	Pancake = Shape{
		{'@', '@', '@', '@'},
	}

	// .#.
	// ###
	// .#.
	HotCross = Shape{
		{'.', '@', '.'},
		{'@', '@', '@'},
		{'.', '@', '.'},
	}

	// ..#
	// ..#
	// ###
	Elbow = Shape{
		{'.', '.', '@'},
		{'.', '.', '@'},
		{'@', '@', '@'},
	}

	// #
	// #
	// #
	// #
	Ih = Shape{
		{'@'},
		{'@'},
		{'@'},
		{'@'},
	}

	// ##
	// ##
	Square = Shape{
		{'@', '@'},
		{'@', '@'},
	}

	ShapeInventory = []Shape{
		Pancake,
		HotCross,
		Elbow,
		Ih,
		Square,
	}
)

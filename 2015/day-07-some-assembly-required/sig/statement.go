package sig

import "fmt"

type Statement struct {
	Operator     Operator
	InputOne     any
	InputTwo     any
	Output       string
	OutputSignal uint16
}

func (s *Statement) Inflate(store map[string]uint16) {
	if !s.IsOperable() {
		if s.InputOne != nil {
			if inputOne, ok := s.InputOne.(string); ok {
				if _, ok := store[inputOne]; ok {
					s.InputOne = store[inputOne]
				}
			}
		}
		if s.InputTwo != nil {
			if inputTwo, ok := s.InputTwo.(string); ok {
				if _, ok := store[inputTwo]; ok {
					s.InputTwo = store[inputTwo]
				}
			}
		}
	}
}

func (s *Statement) Execute() (uint16, bool) {
	if s.IsOperable() {
		switch s.Operator {
		case OperatorPassThrough:
			return s.InputOne.(uint16), true
		case OperatorNot:
			return ^s.InputOne.(uint16), true
		case OperatorAnd:
			return s.InputOne.(uint16) & s.InputTwo.(uint16), true
		case OperatorOr:
			return s.InputOne.(uint16) | s.InputTwo.(uint16), true
		case OperatorRShift:
			return s.InputOne.(uint16) >> s.InputTwo.(uint16), true
		case OperatorLShift:
			return s.InputOne.(uint16) << s.InputTwo.(uint16), true
		}
	}
	return 0, false
}

func (s *Statement) IsOperable() bool {
	inputOneIsNum := false
	if _, ok := s.InputOne.(uint16); ok {
		inputOneIsNum = true
	}

	inputTwoIsNum := false
	if _, ok := s.InputTwo.(uint16); ok {
		inputTwoIsNum = true
	}

	switch s.Operator {
	case OperatorPassThrough:
		fallthrough
	case OperatorNot:
		return inputOneIsNum
	case OperatorAnd:
		fallthrough
	case OperatorOr:
		fallthrough
	case OperatorRShift:
		fallthrough
	case OperatorLShift:
		return inputOneIsNum && inputTwoIsNum
	}

	return false
}

func (s *Statement) String() string {
	return fmt.Sprintf("{%v %v %v -> %v}", s.InputOne, s.Operator, s.InputTwo, s.Output)
}

func NewPassThroughStatement(input any, output string) *Statement {
	return &Statement{
		Operator: OperatorPassThrough,
		InputOne: input,
		Output:   output,
	}
}

func NewNotStatement(input any, output string) *Statement {
	return &Statement{
		Operator: OperatorNot,
		InputOne: input,
		Output:   output,
	}
}

func NewAndStatement(inputOne any, inputTwo any, output string) *Statement {
	return &Statement{
		Operator: OperatorAnd,
		InputOne: inputOne,
		InputTwo: inputTwo,
		Output:   output,
	}
}

func NewOrStatement(inputOne any, inputTwo any, output string) *Statement {
	return &Statement{
		Operator: OperatorOr,
		InputOne: inputOne,
		InputTwo: inputTwo,
		Output:   output,
	}
}

func NewRShiftStatement(inputOne any, inputTwo any, output string) *Statement {
	return &Statement{
		Operator: OperatorRShift,
		InputOne: inputOne,
		InputTwo: inputTwo,
		Output:   output,
	}
}

func NewLShiftStatement(inputOne any, inputTwo any, output string) *Statement {
	return &Statement{
		Operator: OperatorLShift,
		InputOne: inputOne,
		InputTwo: inputTwo,
		Output:   output,
	}
}

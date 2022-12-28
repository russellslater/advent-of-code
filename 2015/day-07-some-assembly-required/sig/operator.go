package sig

type Operator int

func (o Operator) String() string {
	switch o {
	case OperatorPassThrough:
		return " -> "
	case OperatorNot:
		return "NOT"
	case OperatorAnd:
		return "AND"
	case OperatorOr:
		return "OR"
	case OperatorRShift:
		return "RSHIFT"
	case OperatorLShift:
		return "LSHIFT"
	}
	return "UNKNOWN"
}

const (
	OperatorPassThrough Operator = iota
	OperatorNot
	OperatorAnd
	OperatorOr
	OperatorRShift
	OperatorLShift
)

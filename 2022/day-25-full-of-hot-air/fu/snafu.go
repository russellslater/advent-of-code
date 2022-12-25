package fu

type SNAFU string

func (s SNAFU) Decimal() int {
	decimal := 0
	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i]
		placeValue := 1
		for j := i; j < len(s)-1; j++ {
			placeValue *= 5
		}
		switch digit {
		case '2':
			decimal += 2 * placeValue
		case '1':
			decimal += 1 * placeValue
		case '0':
			decimal += 0 * placeValue
		case '-':
			decimal += -1 * placeValue
		case '=':
			decimal += -2 * placeValue
		}
	}
	return decimal
}

func NewSNAFU(decimal int) SNAFU {
	snafu := ""
	for decimal > 0 {
		remainder := decimal % 5
		switch remainder {
		case 0:
			snafu = "0" + snafu
		case 1:
			snafu = "1" + snafu
		case 2:
			snafu = "2" + snafu
		case 3: // 3 = 5 - 2
			snafu = "=" + snafu
			decimal += 5
		case 4: // 4 = 5 - 1
			snafu = "-" + snafu
			decimal += 5
		}
		decimal /= 5
	}
	return SNAFU(snafu)
}

package internals

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (sc *StringCalculator) Add(s string) (result int) {
	if len(s) == 0 {
		return 0
	}
	return 1
}

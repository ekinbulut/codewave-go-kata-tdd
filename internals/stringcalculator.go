package internals

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (sc *StringCalculator) Add(s string) (result int) {
	return 1
}

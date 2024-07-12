package parser

type month struct {
	min int
	max int

	numToEng map[int]string
	engToNum map[string]int
}

func newMonth() *month {
	return &month{
		min: 1,
		max: 12,
	}
}

func (m *month) Validate() error {
	return nil
}

func (m *month) Expand(field string) ([]string, error) {
	return expand(field, m.min, m.max)
}

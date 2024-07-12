package parser

type minute struct {
	min int
	max int

	parsedField []string
}

func newMinute() *minute {
	return &minute{
		min: 0,
		max: 59,
	}
}

func (m *minute) Validate() error {
	return nil
}

func (m *minute) Expand(field string) ([]string, error) {
	return expand(field, m.min, m.max)
}

func (m *minute) Print() {

}

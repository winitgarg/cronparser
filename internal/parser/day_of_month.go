package parser

type dayOfMonth struct {
	min int
	max int
}

func newDayOfMonth() *dayOfMonth {
	return &dayOfMonth{
		min: 1,
		max: 31,
	}
}

func (d *dayOfMonth) Expand(field string) ([]string, error) {
	return expand(field, d.min, d.max)
}

package parser

type hour struct {
	min int
	max int
}

func newHour() *hour {
	return &hour{
		min: 0,
		max: 23,
	}
}

func (h *hour) Validate() error {
	return nil
}

func (h *hour) Expand(field string) ([]string, error) {
	return expand(field, h.min, h.max)
}

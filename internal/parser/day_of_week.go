package parser

type dayOfWeek struct {
	min int
	max int
}

func newDayOfWeek() *dayOfWeek {
	return &dayOfWeek{
		min: 0,
		max: 7,
	}
}

func (d *dayOfWeek) Expand(field string) ([]string, error) {
	result, err := expand(field, d.min, d.max)

	// if both "0" and "7" exists then remove "7"
	var newResult []string
	zeroExists := false

	for _, each := range result {
		if each == "0" {
			zeroExists = true
		}

		if each == "7" && zeroExists {
			continue
		} else {
			newResult = append(newResult, each)
		}
	}

	return newResult, err
}

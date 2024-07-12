package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func expand(field string, min, max int) ([]string, error) {
	if field == "*" {
		result := expandAllValues(min, max)

		return result, nil
	}

	var (
		result []string
		err    error
	)

	values := strings.Split(field, ",") // only handles the scenario 0-4,8-12, other than that len(values) should be 1
	for _, value := range values {
		if strings.Contains(value, "/") { // handling range/step scenario
			stepResult, err := expandSteps(value, min, max)
			if err != nil {
				return nil, fmt.Errorf("error in expanding steps for value: %s, err: %s", value, err)
			}

			result = append(result, stepResult...)
		} else if strings.Contains(value, "-") { // handling scenario 4-23
			rangeResult, err := expandRangeValues(value, min, max)
			if err != nil {
				return nil, fmt.Errorf("invalid range field: %s", value)
			}

			result = append(result, rangeResult...)

		} else { // handling just a single number
			var num int

			num, err = strconv.Atoi(value)
			if err != nil || num < min || num > max {
				return nil, fmt.Errorf("invalid value: %s", value)
			}

			result = append(result, value)
		}
	}

	return result, nil
}

func expandAllValues(min, max int) []string {
	var result []string

	for i := min; i <= max; i++ {
		result = append(result, strconv.Itoa(i))
	}

	return result
}

func expandSteps(value string, min, max int) ([]string, error) {
	if value == "" {
		return nil, fmt.Errorf("value cannot be empty")
	}

	if max < min {
		return nil, fmt.Errorf("upper bound cannot be less than lower bound")
	}

	parts := strings.Split(value, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid increment field: %s", value)
	}

	var step int
	step, err := strconv.Atoi(parts[1])
	if err != nil || step <= 0 {
		return nil, fmt.Errorf("invalid interval value: %s", value)
	}

	var lowerRange, upperRange int
	if parts[0] == "*" { // scenario */2, if minute then, 0,2,4,6,8,10....58
		lowerRange = min
		upperRange = max
	} else if strings.Contains(parts[0], "-") { // scenario 0-23/2, if minute then, 0,2,4,6,....22
		ranges := strings.Split(parts[0], "-")

		lowerRange, err = strconv.Atoi(ranges[0])
		if err != nil {
			return nil, fmt.Errorf("invalid range field: %s", value)
		}

		upperRange, err = strconv.Atoi(ranges[1])
		if err != nil {
			return nil, fmt.Errorf("invalid range field: %s", value)
		}

	} else { // scenario: 5/3 ; if minute, then 5,8,11,14,17,....
		lowerRange, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid base value: %s", value)
		}

		upperRange = max
	}

	result, err := generateSteps(step, lowerRange, upperRange)
	if err != nil {
		return nil, fmt.Errorf("error in generating steps, err: %w", err)
	}

	return result, nil
}

// expand range/step: 0-23/2 => 0,2,4,6,.....22
func generateSteps(step, min, max int) ([]string, error) {
	if step <= 0 {
		return nil, fmt.Errorf("step cannot be 0 or negative")
	}

	err := verifyMinMax(min, max)
	if err != nil {
		return nil, err
	}

	var result []string

	for i := min; i <= max; i += step {
		result = append(result, strconv.Itoa(i))
	}

	return result, nil
}

func expandRangeValues(value string, min, max int) ([]string, error) {
	if value == "" {
		return nil, fmt.Errorf("value cannot be empty")
	}

	err := verifyMinMax(min, max)
	if err != nil {
		return nil, err
	}

	var result []string

	rangeParts := strings.Split(value, "-")
	if len(rangeParts) != 2 {
		return nil, fmt.Errorf("invalid range field: %s", value)
	}

	start, err1 := strconv.Atoi(rangeParts[0])
	end, err2 := strconv.Atoi(rangeParts[1])
	if err1 != nil || err2 != nil || start > end || start < min || end > max {
		return nil, fmt.Errorf("invalid range values: %s", value)
	}

	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
	}

	return result, nil
}

func verifyMinMax(min, max int) error {
	if min < 0 || max < 0 {
		return fmt.Errorf("upper bound or lower bound cannot be less than 0")
	}

	if max < min {
		return fmt.Errorf("upper bound cannot be less than lower bound")
	}

	return nil
}

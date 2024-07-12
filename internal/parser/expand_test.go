package parser

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGenerateSteps(t *testing.T) {
	tests := []struct {
		msg        string
		step       int
		lowerBound int
		upperBound int
		expOutput  []string
		expErr     error
	}{
		// Valid scenarios
		{"Valid scenario", 2, 0, 10, []string{"0", "2", "4", "6", "8", "10"}, nil},
		{"Valid scenario", 3, 5, 20, []string{"5", "8", "11", "14", "17", "20"}, nil},
		{"Valid scenario", 1, 0, 3, []string{"0", "1", "2", "3"}, nil},
		{"Valid scenario", 10, 0, 100, []string{"0", "10", "20", "30", "40", "50", "60", "70", "80", "90", "100"}, nil},

		// Edge case when step equals range span
		{"Edge case", 15, 0, 15, []string{"0", "15"}, nil},
		{"Edge case", 59, 0, 59, []string{"0", "59"}, nil},
		{"Edge case: step larger than span", 60, 0, 59, []string{"0"}, nil}, // Step larger than range span

		// Invalid step scenarios
		{"Invalid step: 0", 0, 0, 10, nil, errors.New("step cannot be 0 or negative")},
		{"Invalid step: negative ", -1, 0, 10, nil, errors.New("step cannot be 0 or negative")},
		{"Invalid step: upper bound less than lower bound", 1, 10, 5, nil, errors.New("upper bound cannot be less than lower bound")},
		{"Invalid step: upper bound less than lower bound with negative value", 5, 0, -5, nil, errors.New("upper bound or lower bound cannot be less than 0")},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualOutput, actualErr := generateSteps(test.step, test.lowerBound, test.upperBound)

			if actualErr != nil {
				assert.Equal(t, test.expErr, actualErr)
			} else {
				assert.Nil(t, test.expErr)
			}

			assert.Equal(t, test.expOutput, actualOutput)
		})
	}

}

func TestExpandRangeValues(t *testing.T) {
	tests := []struct {
		msg        string
		value      string
		lowerBound int
		upperBound int
		expOutput  []string
		expErr     error
	}{
		// Valid range within min and max
		{"Valid range within lower and upper bound: 1-5", "1-5", 0, 10, []string{"1", "2", "3", "4", "5"}, nil},
		{"Valid range within lower and upper bound: 0-3", "0-3", 0, 5, []string{"0", "1", "2", "3"}, nil},
		{"Valid range within lower and upper bound: 5-8", "5-8", 1, 10, []string{"5", "6", "7", "8"}, nil},

		// Invalid range: start > end
		{"Invalid range: start > end", "5-3", 1, 10, nil, errors.New("invalid range values: 5-3")},

		// Out of bounds range
		{"Invalid range: out of bounds 0-10", "0-10", 1, 9, nil, errors.New("invalid range values: 0-10")},
		{"Invalid range: out of bounds 5-15", "5-15", 1, 10, nil, errors.New("invalid range values: 5-15")},

		// Invalid format
		{"Invalid format 5-", "5-", 1, 10, nil, errors.New("invalid range values: 5-")},
		{"Invalid format -3", "-3", 0, 5, nil, errors.New("invalid range values: -3")},
		{"Invalid format 5-10-15", "5-10-15", 0, 20, nil, errors.New("invalid range field: 5-10-15")},

		// Empty value
		{"Empty value passed", "", 0, 10, nil, errors.New("value cannot be empty")},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualOutput, actualErr := expandRangeValues(test.value, test.lowerBound, test.upperBound)

			assert.Equal(t, test.expOutput, actualOutput)
			assert.Equal(t, test.expErr, actualErr)
		})
	}
}

func TestExpandSteps(t *testing.T) {
	tests := []struct {
		msg        string
		value      string
		lowerBound int
		upperBound int
		expOutput  []string
		expErr     error
	}{
		// Valid scenarios
		{"Valid scenario: */2", "*/2", 0, 59, []string{"0", "2", "4", "6", "8", "10", "12", "14", "16", "18", "20", "22", "24", "26", "28", "30", "32", "34", "36", "38", "40", "42", "44", "46", "48", "50", "52", "54", "56", "58"}, nil},
		{"Valid scenario: 0-10/3", "0-10/3", 0, 59, []string{"0", "3", "6", "9"}, nil},
		{"Valid scenario: 5/3", "5/3", 0, 59, []string{"5", "8", "11", "14", "17", "20", "23", "26", "29", "32", "35", "38", "41", "44", "47", "50", "53", "56", "59"}, nil},
		{"Valid scenario: 10-30/5", "10-30/5", 0, 59, []string{"10", "15", "20", "25", "30"}, nil},

		// Invalid scenarios
		{"Invalid step: 0", "*/0", 0, 59, nil, errors.New("invalid interval value: */0")},
		{"Invalid step: negative", "*/-1", 0, 59, nil, errors.New("invalid interval value: */-1")},
		{"Invalid step: non-integer", "5/abc", 0, 59, nil, errors.New("invalid interval value: 5/abc")},
		{"Invalid format: 5/3/2", "5/3/2", 0, 59, nil, errors.New("invalid increment field: 5/3/2")},
		{"Invalid range format: 5-/2", "5-/2", 0, 59, nil, errors.New("invalid range field: 5-/2")},
		{"Invalid range values: lowerBound > upperBound", "5-3/2", 0, 59, nil, errors.New("error in generating steps, err: upper bound cannot be less than lower bound")},
		{"Invalid, lowerBound > upperBound", "5/3", 10, 5, nil, errors.New("upper bound cannot be less than lower bound")},
		{"Invalid range: step as *)", "*/*", 0, 59, nil, errors.New("invalid interval value: */*")},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualOutput, actualErr := expandSteps(test.value, test.lowerBound, test.upperBound)

			assert.Equal(t, test.expOutput, actualOutput)

			if actualErr != nil {
				assert.Equal(t, test.expErr.Error(), actualErr.Error())
			} else {
				assert.Nil(t, test.expErr)
			}

		})
	}
}

func TestExpandAllValues(t *testing.T) {
	tests := []struct {
		msg        string
		lowerBound int
		upperBound int
		expOutput  []string
	}{
		// Simple Range
		{"Range 0-5", 0, 5, []string{"0", "1", "2", "3", "4", "5"}},
		{"Range 1-3", 1, 3, []string{"1", "2", "3"}},

		// Single Value Range
		{"Range with same lower and bound: 4", 4, 4, []string{"4"}},
		{"Range with same lower and bound: 0", 0, 0, []string{"0"}},

		// Full Range
		{"Range covering 0-59", 0, 59, func() []string {
			var result []string
			for i := 0; i <= 59; i++ {
				result = append(result, strconv.Itoa(i))
			}
			return result
		}()},

		// Edge Cases
		{"Range when upper bound < lower bound", 5, 1, nil},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualOutput := expandAllValues(test.lowerBound, test.upperBound)

			assert.Equal(t, test.expOutput, actualOutput)
		})
	}
}

func TestExpand(t *testing.T) {
	tests := []struct {
		msg        string
		field      string
		lowerBound int
		upperBound int
		expOutput  []string
		expErr     error
	}{
		// Wildcard
		{"Handling Wildcard * : with 0 to 5", "*", 0, 5, []string{"0", "1", "2", "3", "4", "5"}, nil},
		{"Handling Wildcard * : with 1 to 3", "*", 1, 3, []string{"1", "2", "3"}, nil},

		// Single Values
		{"Handling Single value: 3", "3", 0, 5, []string{"3"}, nil},
		{"Handling Single value: 0", "0", 0, 10, []string{"0"}, nil},
		{"Handling Single value: 10", "10", 0, 10, []string{"10"}, nil},

		// Ranges
		{"Handling Ranges: 1-3", "1-3", 0, 5, []string{"1", "2", "3"}, nil},
		{"Handling Ranges: 0-4", "0-4", 0, 10, []string{"0", "1", "2", "3", "4"}, nil},
		{"Handling Ranges with same lower and upper bound: 5-5", "5-5", 0, 10, []string{"5"}, nil},

		// Steps
		{"Handling Steps: */2", "*/2", 0, 5, []string{"0", "2", "4"}, nil},
		{"Handling Steps: 1-4/2", "1-4/2", 0, 5, []string{"1", "3"}, nil},
		{"Handling Steps: 5/2", "5/5", 0, 10, []string{"5", "10"}, nil},

		// List
		{"Handling List: 1,3,5", "1,3,5", 0, 5, []string{"1", "3", "5"}, nil},
		{"Handling List: 10,20,30", "10,20,30", 0, 30, []string{"10", "20", "30"}, nil},

		// Invalid Cases
		{"Invalid Cases: a", "a", 0, 5, nil, errors.New("invalid value")},                     // Invalid string
		{"Invalid Cases: */a", "*/a", 0, 5, nil, errors.New("invalid interval")},              // Invalid step
		{"Invalid Cases: 3-1", "3-1", 0, 5, nil, errors.New("invalid range")},                 // Invalid range (start > end)
		{"Invalid Cases: 4-2/2", "4-2/2", 0, 5, nil, errors.New("error in generating steps")}, // Invalid range (start > end)
		{"Invalid Cases: 5/0", "5/0", 0, 10, nil, errors.New("invalid interval")},             // Invalid step (0)
		{"Invalid Cases: 6", "6", 0, 5, nil, errors.New("invalid value")},                     // Out of bounds single value
		{"Invalid Cases: 0-10", "0-10", 0, 5, nil, errors.New("invalid range")},               // Out of bounds range
		{"Invalid Cases: */0", "*/0", 0, 5, nil, errors.New("invalid interval")},              // Invalid step (0)
		{"Invalid Cases: -1/5", "-1/5", 0, 5, nil, errors.New("invalid range")},               // Invalid range (negative start)
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualOutput, actualErr := expand(test.field, test.lowerBound, test.upperBound)

			assert.Equal(t, test.expOutput, actualOutput)

			if actualErr != nil {
				assert.ErrorContains(t, actualErr, test.expErr.Error())
				//assert.Equal(t, test.expErr, actualErr)
			} else {
				assert.Nil(t, test.expErr)
			}
		})
	}
}

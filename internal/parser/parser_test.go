package parser

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		msg      string
		input    string
		expError error
	}{
		{"the input has 5 space separated cron fields", "*/15 0 1,15 * 1-5", nil},
		{"the input has 5 space separated any fields", "a b c d e", nil},
		{"the input has 0 fields", "", errors.New("invalid cron expression: expected 5 fields, got 0")},
		{"the input has 3 fields", "a b c", errors.New("invalid cron expression: expected 5 fields, got 3")},
		{"the input has 6 fields", "a b c d e f", errors.New("invalid cron expression: expected 5 fields, got 6")},
	}

	cron := New()
	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualErr := cron.validate(test.input)

			assert.Equal(t, test.expError, actualErr)
		})
	}
}

func TestExpandCron(t *testing.T) {
	tests := []struct {
		msg   string
		input string

		expOutMinute     []string
		expOutHour       []string
		expOutDayOfMonth []string
		expOutMonth      []string
		expOutDayOfWeek  []string

		expError error
	}{
		{
			msg:              "Test case for simple cron expression",
			input:            "0-59 0-23 1-31 1-12 0-6 /command",
			expOutMinute:     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"},
			expOutHour:       []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"},
			expOutDayOfMonth: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
			expOutMonth:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
			expOutDayOfWeek:  []string{"0", "1", "2", "3", "4", "5", "6"},
			expError:         nil,
		},
		{
			msg:              "Test case with step values",
			input:            "*/15 */6 */10 */2 */2 /command",
			expOutMinute:     []string{"0", "15", "30", "45"},
			expOutHour:       []string{"0", "6", "12", "18"},
			expOutDayOfMonth: []string{"1", "11", "21", "31"},
			expOutMonth:      []string{"1", "3", "5", "7", "9", "11"},
			expOutDayOfWeek:  []string{"0", "2", "4", "6"},
			expError:         nil,
		},
		{
			msg:              "Test case with specific values",
			input:            "0,30 9,21 1,15 1,7 1,3,5 /command",
			expOutMinute:     []string{"0", "30"},
			expOutHour:       []string{"9", "21"},
			expOutDayOfMonth: []string{"1", "15"},
			expOutMonth:      []string{"1", "7"},
			expOutDayOfWeek:  []string{"1", "3", "5"},
			expError:         nil,
		},
		{
			msg:              "Test case for asterisk values",
			input:            "* * * * * /command",
			expOutMinute:     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"},
			expOutHour:       []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"},
			expOutDayOfMonth: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
			expOutMonth:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
			expOutDayOfWeek:  []string{"0", "1", "2", "3", "4", "5", "6"},
			expError:         nil,
		},
		{
			msg:              "Test case with invalid minute value",
			input:            "60 12 10 5 3 /command",
			expOutMinute:     nil,
			expOutHour:       nil,
			expOutDayOfMonth: nil,
			expOutMonth:      nil,
			expOutDayOfWeek:  nil,
			expError:         errors.New("error in expanding cron expression"),
		},
		{
			msg:              "Test case with invalid day of month value",
			input:            "30 14 -1 6 4 /command",
			expOutMinute:     nil,
			expOutHour:       nil,
			expOutDayOfMonth: nil,
			expOutMonth:      nil,
			expOutDayOfWeek:  nil,
			expError:         errors.New("error in expanding cron expression"),
		},
		{
			msg:              "Test case with mixed list and range",
			input:            "0,15-30,45 5,9-11,14 1-10,15,20-25 1,6-8,11 0,3-5,7 /command",
			expOutMinute:     []string{"0", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "45"},
			expOutHour:       []string{"5", "9", "10", "11", "14"},
			expOutDayOfMonth: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "15", "20", "21", "22", "23", "24", "25"},
			expOutMonth:      []string{"1", "6", "7", "8", "11"},
			expOutDayOfWeek:  []string{"0", "3", "4", "5"},
			expError:         nil,
		},
		{
			msg:              "Test case with combined wildcards and list",
			input:            "0,30 */12 */5 */3 1,3,5 /command",
			expOutMinute:     []string{"0", "30"},
			expOutHour:       []string{"0", "12"},
			expOutDayOfMonth: []string{"1", "6", "11", "16", "21", "26", "31"},
			expOutMonth:      []string{"1", "4", "7", "10"},
			expOutDayOfWeek:  []string{"1", "3", "5"},
			expError:         nil,
		},
		{
			msg:              "Test case with complex range and step",
			input:            "0-59/20 0-23/8 1-30/10 1-12/4 0-6/2 /command",
			expOutMinute:     []string{"0", "20", "40"},
			expOutHour:       []string{"0", "8", "16"},
			expOutDayOfMonth: []string{"1", "11", "21"},
			expOutMonth:      []string{"1", "5", "9"},
			expOutDayOfWeek:  []string{"0", "2", "4", "6"},
			expError:         nil,
		},
		{
			msg:              "Test case with non-numeric values",
			input:            "a b c d e /command",
			expOutMinute:     nil,
			expOutHour:       nil,
			expOutDayOfMonth: nil,
			expOutMonth:      nil,
			expOutDayOfWeek:  nil,
			expError:         errors.New("error in expanding cron expression"),
		},
		{
			msg:              "Test case with missing fields",
			input:            "30  10  3 /command",
			expOutMinute:     nil,
			expOutHour:       nil,
			expOutDayOfMonth: nil,
			expOutMonth:      nil,
			expOutDayOfWeek:  nil,
			expError:         errors.New("invalid cron expression"),
		},
	}

	cron := New()

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			actualErr := cron.Parse(test.input)

			if actualErr != nil {
				assert.ErrorContains(t, actualErr, test.expError.Error())
			} else {
				assert.Nil(t, test.expError)
				assert.Equal(t, test.expOutMinute, cron.minute.minuteParsed)
				assert.Equal(t, test.expOutHour, cron.hour.hourParsed)
				assert.Equal(t, test.expOutDayOfMonth, cron.dayOfMonth.dayOfMonthParsed)
				assert.Equal(t, test.expOutMonth, cron.month.monthParsed)
				assert.Equal(t, test.expOutDayOfWeek, cron.dayOfWeek.dayOfWeekParsed)
			}
		})
	}
}

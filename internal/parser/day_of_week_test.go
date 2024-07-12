package parser

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayOfWeekExpand(t *testing.T) {
	tests := []struct {
		msg    string
		field  string
		min    int
		max    int
		expOut []string
		expErr error
	}{
		{
			msg:    "Test with full range wildcard",
			field:  "*",
			min:    0,
			max:    6,
			expOut: []string{"0", "1", "2", "3", "4", "5", "6"},
			expErr: nil,
		},
		{
			msg:    "Test with specific values",
			field:  "0,7",
			min:    0,
			max:    6,
			expOut: []string{"0"},
			expErr: nil,
		},
		{
			msg:    "Test with mixed values",
			field:  "0,2,7",
			min:    0,
			max:    6,
			expOut: []string{"0", "2"},
			expErr: nil,
		},
		{
			msg:    "Test with range values",
			field:  "0-7",
			min:    0,
			max:    6,
			expOut: []string{"0", "1", "2", "3", "4", "5", "6"},
			expErr: nil,
		},
		{
			msg:    "Test with invalid value",
			field:  "8",
			min:    0,
			max:    6,
			expOut: nil,
			expErr: errors.New("invalid value: 8"),
		},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			dow := newDayOfWeek()
			actualOut, actualErr := dow.Expand(test.field)

			if actualErr != nil {
				assert.ErrorContains(t, actualErr, test.expErr.Error())
			} else {
				assert.Nil(t, test.expErr)
				assert.Equal(t, test.expOut, actualOut)
			}
		})
	}
}

package parser

import (
	"fmt"
	"strings"
)

type CronField interface {
	Expand(field string) ([]string, error)
}

type Cron struct {
	expression string
	command    string

	minute     *minuteHandler
	hour       *hourHandler
	dayOfMonth *dayOfMonthHandler
	month      *monthHandler
	dayOfWeek  *dayOfWeekHandler
}

func New() *Cron {
	newMinuteHandler := &minuteHandler{
		minuteParser: newMinute(),
	}

	newHourHandler := &hourHandler{
		hourParser: newHour(),
	}

	newDayOfMonthHandler := &dayOfMonthHandler{
		dayOfMonthParser: newDayOfMonth(),
	}

	newMonthHandler := &monthHandler{
		monthParser: newMonth(),
	}

	newDayOfWeekHandler := &dayOfWeekHandler{
		dayOfWeekParser: newDayOfWeek(),
	}

	return &Cron{
		minute:     newMinuteHandler,
		hour:       newHourHandler,
		dayOfMonth: newDayOfMonthHandler,
		month:      newMonthHandler,
		dayOfWeek:  newDayOfWeekHandler,
	}
}

func (c *Cron) Parse(input string) error {
	parts := strings.SplitN(input, " ", 6)
	if len(parts) != 6 {
		return fmt.Errorf("incorrect input format")
	}

	cronExpr := strings.Join(parts[:5], " ")
	c.command = parts[5]

	fields := strings.Fields(cronExpr)
	if len(fields) != 5 {
		return fmt.Errorf("invalid cron expression: expected 5 fields, got %d", len(fields))
	}

	c.minute.minuteField = fields[0]
	c.hour.hourField = fields[1]
	c.dayOfMonth.dayOfMonthField = fields[2]
	c.month.monthField = fields[3]
	c.dayOfWeek.dayOfWeekField = fields[4]

	err := c.expand()
	if err != nil {
		return fmt.Errorf("error in expanding cron expression: %s, err: %w", cronExpr, err)
	}

	c.print()

	return nil
}

func (c *Cron) validate(exp string) error {
	fields := strings.Fields(exp)

	if len(fields) != 5 {
		return fmt.Errorf("invalid cron expression: expected 5 fields, got %d", len(fields))
	}

	c.minute.minuteField = fields[0]
	c.hour.hourField = fields[1]
	c.dayOfMonth.dayOfMonthField = fields[2]
	c.month.monthField = fields[3]
	c.dayOfWeek.dayOfWeekField = fields[4]

	return nil
}

func (c *Cron) expand() error {
	var err error

	c.minute.minuteParsed, err = c.minute.minuteParser.Expand(c.minute.minuteField)
	if err != nil {
		return fmt.Errorf("error in parsing minute. err: %w", err)
	}

	c.hour.hourParsed, err = c.hour.hourParser.Expand(c.hour.hourField)
	if err != nil {
		return fmt.Errorf("error in parsing hour. err: %w", err)
	}

	c.dayOfMonth.dayOfMonthParsed, err = c.dayOfMonth.dayOfMonthParser.Expand(c.dayOfMonth.dayOfMonthField)
	if err != nil {
		return fmt.Errorf("error in parsing day of month. err: %w", err)
	}

	c.month.monthParsed, err = c.month.monthParser.Expand(c.month.monthField)
	if err != nil {
		return fmt.Errorf("error in parsing month. err: %w", err)
	}

	c.dayOfWeek.dayOfWeekParsed, err = c.dayOfWeek.dayOfWeekParser.Expand(c.dayOfWeek.dayOfWeekField)
	if err != nil {
		return fmt.Errorf("error in parsing day of week. err: %w", err)
	}

	return nil
}

func (c *Cron) print() {
	fmt.Printf("%-14s%s\n", "minute", strings.Join(c.minute.minuteParsed, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(c.hour.hourParsed, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(c.dayOfMonth.dayOfMonthParsed, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(c.month.monthParsed, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(c.dayOfWeek.dayOfWeekParsed, " "))
	fmt.Printf("%-14s%s\n", "command", c.command)
}

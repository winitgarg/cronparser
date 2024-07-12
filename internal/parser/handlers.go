package parser

type minuteHandler struct {
	minuteField  string
	minuteParsed []string
	minuteParser CronField
}

type hourHandler struct {
	hourField  string
	hourParsed []string
	hourParser CronField
}

type dayOfMonthHandler struct {
	dayOfMonthField  string
	dayOfMonthParsed []string
	dayOfMonthParser CronField
}

type monthHandler struct {
	monthField  string
	monthParsed []string
	monthParser CronField
}

type dayOfWeekHandler struct {
	dayOfWeekField  string
	dayOfWeekParsed []string
	dayOfWeekParser CronField
}

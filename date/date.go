package date

import (
	"strconv"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d Date) IsEmpty() bool {
	return d.Year == 0 && d.Month == 0 && d.Day == 0
}

func (d Date) ToString() string {
	return strconv.Itoa(d.Year) + "-" + strconv.Itoa(d.Month) + "-" + strconv.Itoa(d.Day)
}

func (d Date) Equal(target Date) bool {
	return d.Year == target.Year && d.Month == target.Month && d.Day == target.Day
}

func NewDate(y int, m int, d int) Date {
	return Date{
		Year:  y,
		Month: m,
		Day:   d,
	}
}

func YesterdayDate() Date {
	t := time.Now().AddDate(0, 0, -1)
	return NewDate(t.Year(), int(t.Month()), t.Day())
}

func Parse(v string) (Date, error) {
	t, err := time.Parse("2006-1-2", v)
	if err != nil {
		return Date{}, err
	}
	return NewDate(t.Year(), int(t.Month()), t.Day()), nil
}

func ParseDate(v string) (Date, error) {
	t, err := time.Parse("2006年1月2日", v)
	if err != nil {
		return Date{}, err
	}
	return NewDate(t.Year(), int(t.Month()), t.Day()), nil
}

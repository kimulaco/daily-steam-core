package date

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	t.Run("Date.ToString", func(t *testing.T) {
		d := Date{Year: 2011, Month: 11, Day: 11}
		if d.ToString() != "2011-11-11" {
			t.Error(d.ToString())
		}
	})

	t.Run("Date.IsEmpty", func(t *testing.T) {
		d := Date{Year: 0, Month: 0, Day: 0}
		if !d.IsEmpty() {
			t.Error("IsEmpty should return true for 0-0-0")
		}
	})

	t.Run("Date.Equal", func(t *testing.T) {
		d1 := Date{Year: 2011, Month: 11, Day: 11}
		d2 := Date{Year: 2011, Month: 11, Day: 11}
		if !d1.Equal(d2) {
			t.Error("d1 and d2 are equal")
		}
	})
}

func TestNewDate(t *testing.T) {
	d := NewDate(2011, 11, 11)
	if d.Year != 2011 || d.Month != 11 || d.Day != 11 {
		t.Error(d.ToString())
	}
}

func TestYesterdayDate(t *testing.T) {
	d := YesterdayDate()
	_t := time.Now()
	if d.Year != _t.Year() || d.Month != int(_t.Month()) || d.Day != _t.Day()-1 {
		t.Error(d.ToString())
	}
}

func TestParse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		d, err := Parse("2022-8-1")
		if err != nil {
			t.Error(err)
		}
		if d.Year != 2022 || d.Month != 8 || d.Day != 1 {
			t.Error(d.ToString())
		}
	})

	t.Run("error", func(t *testing.T) {
		d, err := ParseDate("20220801")
		if err == nil {
			t.Error("return an error because 20220801 cannot be parsed")
		}
		if !d.IsEmpty() {
			t.Error("IsEmpty should be true on error")
		}
	})
}

func TestParseDate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		d, err := ParseDate("2022年8月1日")
		if err != nil {
			t.Error(err)
		}
		if d.Year != 2022 || d.Month != 8 || d.Day != 1 {
			t.Error(d.ToString())
		}
	})

	t.Run("error", func(t *testing.T) {
		d, err := ParseDate("20220801")
		if err == nil {
			t.Error("return an error because 20220801 cannot be parsed")
		}
		if !d.IsEmpty() {
			t.Error("IsEmpty should be true on error")
		}
	})
}

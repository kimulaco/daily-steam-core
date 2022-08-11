package apiutil

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	e := Error{
		StatusCode: 400,
		ErrorCode:  "REQUIRED_ID",
		Message:    "required id paramater",
	}

	t.Run("create Error", func(t *testing.T) {
		if e.StatusCode != 400 ||
			e.ErrorCode != "REQUIRED_ID" ||
			e.Message != "required id paramater" {
			t.Error("invalid Error struct")
		}
	})

	t.Run("Error.ToBytes()", func(t *testing.T) {
		_, err := e.ToBytes()
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestNewError(t *testing.T) {
	t.Run("create Error", func(t *testing.T) {
		e := NewError(400, "REQUIRED_ID", errors.New("required id paramater"))
		if e.StatusCode != 400 ||
			e.ErrorCode != "REQUIRED_ID" ||
			e.Message != "required id paramater" {
			t.Error("invalid Error struct")
		}
	})
}

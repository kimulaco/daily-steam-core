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

	t.Run("Error.IsEmpty()", func(t *testing.T) {
		if e.IsEmpty() {
			t.Error("e is not empty")
		}
		empty := Error{}
		if !empty.IsEmpty() {
			t.Error("empty is not empty")
		}
	})
}

func TestNewError(t *testing.T) {
	e := NewError(400, "REQUIRED_ID", errors.New("required id paramater"))
	if e.StatusCode != 400 ||
		e.ErrorCode != "REQUIRED_ID" ||
		e.Message != "required id paramater" {
		t.Error("invalid Error struct")
	}
}

func TestNew400(t *testing.T) {
	e := NewError400("REQUIRED_ID", "required id paramater")
	if e.StatusCode != 400 ||
		e.ErrorCode != "REQUIRED_ID" ||
		e.Message != "required id paramater" {
		t.Error("invalid Error struct")
	}
}

func TestNew500(t *testing.T) {
	t.Run("struct value", func(t *testing.T) {
		e := NewError500("INTERNAL_ERROR", "error text")
		if e.StatusCode != 500 ||
			e.ErrorCode != "INTERNAL_ERROR" ||
			e.Message != "error text" {
			t.Error("invalid Error struct")
		}
	})

	t.Run("default msg", func(t *testing.T) {
		e := NewError500("INTERNAL_ERROR", "")
		if e.StatusCode != 500 ||
			e.ErrorCode != "INTERNAL_ERROR" ||
			e.Message != MSG_500 {
			t.Error("default msg is '" + MSG_500 + "'")
		}
	})
}

package apiutil

import (
	"encoding/json"
	"errors"
)

type Error struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func (e Error) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}

func (e Error) IsEmpty() bool {
	return e.StatusCode == 0 && e.ErrorCode == "" && e.Message == ""
}

func NewError(
	statucCode int,
	errorCode string,
	err error,
) Error {
	return Error{
		StatusCode: statucCode,
		ErrorCode:  errorCode,
		Message:    err.Error(),
	}
}

func NewError400(errorCode string, msg string) Error {
	return NewError(400, errorCode, errors.New(msg))
}

func NewError500(errorCode string, msg string) Error {
	if msg == "" {
		msg = MSG_500
	}
	return NewError(500, errorCode, errors.New(msg))
}

const MSG_500 = "internal server error"

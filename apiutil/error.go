package apiutil

import (
	"encoding/json"
)

type Error struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func (e Error) ToBytes() ([]byte, error) {
	return json.Marshal(e)
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

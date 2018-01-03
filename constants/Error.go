package constants

import (
	"errors"
	"fmt"
	"strings"
)

const (
	ErrMsgInternalServerError = "tInternalServerError"
)

var (
	ErrParameters = errors.New("wrong parameters")

	ErrIdDuplicated = errors.New("data is duplicated")

	IsDuplicatedErr = func(err error) bool {
		return strings.ContainsAny(err.Error(), "Duplicate")
	}

	ErrNoAuth = errors.New("unauthorized")

	ErrClientIp = errors.New("client ip is different")

	ErrNoTransaction = errors.New("this transaction is empty")

	Err404NotFound = errors.New("page not found")

	ErrNotNewData = errors.New("this is not a new data")

	ErrUnusedFunc = errors.New("unused function")

	ErrNoId = errors.New("ID not specified")
)

func ErrDataNotExist(dataName string) error {
	return errors.New(fmt.Sprintf("This %s does not exist.", dataName))
}

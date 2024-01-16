package rest

import (
// errPkg "board/utils/errors"
)

var statusCode = map[string]int{
	"ok":       200,
	"Created":  201,
	"Accepted": 202,
}

func ErrorMessageGenerator(err error) (string, int) {

	return "", 200
}

func SuccessfulMessageGenerator(status string, info string) (string, int) {

	code := statusCode[status]

	return "", code
}

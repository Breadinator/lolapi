package errors

import "errors"

type ErrDataLimitExceeded error

func RaiseDataLimitExceeded(message string) ErrDataLimitExceeded {
	return ErrDataLimitExceeded(errors.New(message))
}

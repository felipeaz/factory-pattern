package errors

import "fmt"

func WithStack(message string, err error) error {
	return fmt.Errorf(message, err)
}

package errors

import "fmt"

func New(text string, args ...any) error {
	return fmt.Errorf("%s %v", text, args)
}

func WithStack(text string, err error) error {
	return fmt.Errorf(text, err)
}

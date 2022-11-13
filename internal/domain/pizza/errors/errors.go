package errors

import (
	"fmt"
)

var (
	errPizzaToppingNotFoundMessage = "pizza topping not found %s"
)

func NotFound(args ...string) error {
	return fmt.Errorf(errPizzaToppingNotFoundMessage, args)
}

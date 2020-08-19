package otpgo

import (
	"fmt"
)

type ErrorInvalidKey struct {
	msg string
}

func (eik ErrorInvalidKey) Error() string {
	return fmt.Sprintf("invalid key: %s", eik.msg)
}

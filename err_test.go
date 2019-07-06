package mplus

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateError_Error(t *testing.T) {
	err := ValidateErrorWrap(errors.New("body parse error"), ErrBodyParse)

	if assert.NotPanics(t, func() { fmt.Println(err.Error()) }) {
		return
	}
}

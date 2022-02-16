package apperrors

import (
	"github.com/wolftotem4/shaolin-ben-don/internal/parser"
)

type LoginRequiredError struct {
	Msg  string
	Form *parser.LoginForm
}

func NewLoginRequiredError(msg string, form *parser.LoginForm) *LoginRequiredError {
	return &LoginRequiredError{
		Msg:  msg,
		Form: form,
	}
}

func (err LoginRequiredError) Error() string {
	return err.Msg
}

package globals

import (
	"errors"
	"fmt"
)

type FreeformgenError struct {
	Context string
	Err     error
}

func (e FreeformgenError) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

func newFreeformgenError() (e FreeformgenError) {
	e.Context = "freeformgen"
	return
}

func NoCommandError() error {
	e := newFreeformgenError()
	e.Err = errors.New("no command provided")
	return e
}

func UnknownCommandError(cmd string) error {
	e := newFreeformgenError()
	e.Err = fmt.Errorf("unknown command %s", cmd)
	return e
}

type DirectivesError struct {
	FreeformgenError
}

func newDirectivesError() (e DirectivesError) {
	e.Context = "directives"
	return
}

func NoSourceError() error {
	e := newDirectivesError()
	e.Err = errors.New("no source provided")
	w := newFreeformgenError()
	w.Err = e
	return w
}

func IncorrectNumberOfArgumentsError(got int) error {
	e := newDirectivesError()
	e.Err = fmt.Errorf("wrong number of arguments: %d", got)
	w := newFreeformgenError()
	w.Err = e
	return w
}

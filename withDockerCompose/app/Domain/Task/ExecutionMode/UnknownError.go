package ExecutionMode

import "fmt"

type UnknownError struct {
	modeString string
}

func (e *UnknownError) Error() string {
	return fmt.Sprintf("mode: %s is invalid", e.modeString)
}

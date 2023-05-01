package usecase

import "fmt"

var (
	SessionExistsErr    = fmt.Errorf("there is active session")
	SessionNotExistsErr = fmt.Errorf("no active session")
)

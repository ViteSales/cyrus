package exceptions

import "fmt"

// UserError is an error that must rollback the current transaction and
// be displayed as a warning to the user.
type UserError struct {
	Message string
	Debug   string
}

// Error method for the UserError type.
// Returns the message.
func (u UserError) Error() string {
	return fmt.Sprintf("%s\n----------------------------------\n%s", u.Message, u.Debug)
}

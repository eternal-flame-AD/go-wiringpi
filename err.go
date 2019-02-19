package wiringpi

import "fmt"

// RetCode is returned when an operation returned non-zero value
type RetCode struct {
	Code int
}

func (c *RetCode) Error() string {
	return fmt.Sprintf("operation failed to complete. retcode: %d", c.Code)
}

func checkRetCode(code int) error {
	if code != 0 {
		return RetCode{code}
	}
	return nil
}

package api

import "fmt"

type MartError struct {
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

func (m MartError) Error() string {
	return fmt.Sprintf("msg: %s", m.Description)
}

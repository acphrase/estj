package exception

import (
	"net/http"
)

type ProfileErrors struct {
	BaseError
}

func (e *ProfileErrors) Error() string {
	return e.setMessage()
}

func CreateProfileErrors(instanceName string, message string) *ProfileErrors {
	return &ProfileErrors{
		BaseError{
			StatusCode:     http.StatusInternalServerError,
			DefaultMessage: "Failed to load profile.",
			InstanceName:   instanceName,
			Message:        message,
		},
	}
}

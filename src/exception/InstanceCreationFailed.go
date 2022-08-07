package exception

import (
	"net/http"
)

type InstanceCreationFailed struct {
	BaseError
}

func (e *InstanceCreationFailed) Error() string {
	return e.setMessage()
}

func CreateInstanceCreationFailed(instanceName string, message string) *InstanceCreationFailed {
	return &InstanceCreationFailed{
		BaseError{
			StatusCode:     http.StatusInternalServerError,
			DefaultMessage: "Failed to create instance.",
			InstanceName:   instanceName,
			Message:        message,
		},
	}
}

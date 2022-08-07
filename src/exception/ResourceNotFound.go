package exception

import (
	"net/http"
)

type ResourceNotFound struct {
	BaseError
}

func (e *ResourceNotFound) Error() string {
	return e.setMessage()
}

func CreateResourceNotFound(instanceName string, message string) *ResourceNotFound {
	return &ResourceNotFound{
		BaseError{
			StatusCode:     http.StatusNotFound,
			DefaultMessage: "Resource not found.",
			InstanceName:   instanceName,
			Message:        message,
		},
	}
}

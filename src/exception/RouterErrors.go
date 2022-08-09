package exception

import (
	"net/http"
)

type RouterErrors struct {
	BaseError
}

func (e *RouterErrors) Error() string {
	return e.setMessage()
}

func CreateRouterErrors(instanceName string, message string) *RouterErrors {
	return &RouterErrors{
		BaseError{
			StatusCode:     http.StatusInternalServerError,
			DefaultMessage: "Failed to configure router.",
			InstanceName:   instanceName,
			Message:        message,
		},
	}
}

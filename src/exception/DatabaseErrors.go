package exception

import (
	"net/http"
)

type DatabaseErrors struct {
	BaseError
}

func (e *DatabaseErrors) Error() string {
	return e.setMessage()
}

func CreateDatabaseErrors(instanceName string, message string) *DatabaseErrors {
	return &DatabaseErrors{
		BaseError{
			StatusCode:     http.StatusInternalServerError,
			DefaultMessage: "DataBase creation failed.",
			InstanceName:   instanceName,
			Message:        message,
		},
	}
}

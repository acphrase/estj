package exception

import (
	"strconv"
	"strings"
)

type BaseError struct {
	StatusCode     int    `json:"status"`
	DefaultMessage string `json:"defaultMessage,omitempty"`
	InstanceName   string `json:"instance,omitempty"`
	Message        string `json:"message"`
}

func (e *BaseError) GetMessage() string {
	return e.setMessage()
}

func (e *BaseError) setMessage() string {
	originMessage := e.Message
	e.Message = "[" + e.InstanceName + "]" + "[" + strconv.Itoa(e.StatusCode) + "] "
	if strings.Contains(originMessage, e.Message) {
		e.Message = originMessage
		return e.Message
	}
	if originMessage == "" {
		e.Message = e.Message + e.DefaultMessage
		e.DefaultMessage = ""
	} else {
		e.Message = e.Message + originMessage
	}
	return e.Message
}

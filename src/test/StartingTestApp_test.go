package test

import (
	"estj/src/core"
	"testing"
)

func TestStarting(t *testing.T) {
	// Just start app
	core.GetApp().RunApp(false)
}

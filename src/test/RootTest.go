package test

import (
	"estj/src/config"
	"estj/src/core"
)

func StartingTest() {
	// Start app : It should be used early in testing.
	core.GetApp().RunApp(false)
}

func EndTest() {
	// End app : It should be used at the end of the test.
	_ = config.GetDB().Close()
}

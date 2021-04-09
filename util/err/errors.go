package err

import "errors"

var (
	ConfigFilePathIsInvalid      = errors.New("config file path is invalid")
	MainProcessContextTerminated = errors.New("main process context done received")
)

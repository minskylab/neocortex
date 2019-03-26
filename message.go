package neocortex

// LogLevelType is level of the logs
type LogLevelType string

// Info is a level of log
var Info LogLevelType = "info"

// Error is a level of log
var Error LogLevelType = "error"

// Warn is a level of log
var Warn LogLevelType = "warn"

// LogMessage represents a snapshot of the messages around the dialog
type LogMessage struct {
	Level   LogLevelType
	Message string
}

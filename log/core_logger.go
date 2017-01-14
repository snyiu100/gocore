package log

import (
	"time"

	"github.com/go-kit/kit/log/levels"
)

type CoreLogger struct {
	levels.Levels
	hideTimestamp bool
}

func NewCoreLogger(l levels.Levels) *CoreLogger {
	return &CoreLogger{Levels: l}
}

func (cl *CoreLogger) LogInfoMessage(message string, keyvalues ...interface{}) {
	cl.LogInfo(append(keyvalues, "msg", message)...)
}

func (cl *CoreLogger) LogErrorMessage(message string, keyvalues ...interface{}) {
	cl.LogError(append(keyvalues, "msg", message)...)
}

func (cl *CoreLogger) LogInfo(keyvals ...interface{}) {
	if len(keyvals) == 1 {
		keyvals = []interface{}{"msg", keyvals[0]}
	}
	cl.Levels.Info().Log(encodeCompoundValues(cl.logTimestamp(keyvals)...)...)
}

func (cl *CoreLogger) LogError(keyvals ...interface{}) {
	if len(keyvals) == 1 {
		keyvals = []interface{}{"msg", keyvals[0]}
	}
	cl.Levels.Error().Log(encodeCompoundValues(cl.logTimestamp(keyvals)...)...)
}

func (cl *CoreLogger) SetStandardFields(keyvals ...interface{}) Logger {
	encoded := encodeCompoundValues(keyvals...)
	newLogger := NewCoreLogger(cl.Levels.With(encoded...))
	newLogger.hideTimestamp = cl.hideTimestamp
	return newLogger
}

func (cl *CoreLogger) With(keyvals ...interface{}) Logger {
	return cl.SetStandardFields(keyvals...)
}

func (cl *CoreLogger) logTimestamp(keyvals []interface{}) []interface{} {
	if !cl.hideTimestamp {
		return append(keyvals, "timestamp", defaultTimeUTC())
	}
	return keyvals
}

func defaultTimeUTC() string {
	return time.Now().UTC().Format(time.RFC3339)
}

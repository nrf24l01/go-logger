package gologger

type Level int

const (
	LevelDebug Level = iota
	LevelTrace
	LevelInfo
	LevelSuccess
	LevelOk
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelTrace:
		return "TRACE"
	case LevelInfo:
		return "INFO"
	case LevelSuccess:
		return "SUCCESS"
	case LevelOk:
		return "OK"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

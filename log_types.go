package gologger

type LogType string

const (
	TypeError LogType = "ERROR"
	TypeInfo  LogType = "INFO"
	TypeDebug LogType = "DEBUG"
)

var defaultLogTypes = []LogType{
	TypeError,
	TypeInfo,
	TypeDebug,
}

func (t LogType) String() string {
	if t == "" {
		return "OTHER"
	}
	return string(t)
}

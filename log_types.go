package gologger

type LogType string

const (
	TypeAuth  LogType = "AUTH"
	TypeHTTP  LogType = "HTTP"
	TypeDB    LogType = "DB"
	TypeCache LogType = "CACHE"
	TypeError LogType = "ERROR"
	TypeInfo  LogType = "INFO"
	TypeDebug LogType = "DEBUG"
	TypeOther LogType = "OTHER"
)

var defaultLogTypes = []LogType{
	TypeAuth,
	TypeHTTP,
	TypeDB,
	TypeCache,
	TypeError,
	TypeInfo,
	TypeDebug,
	TypeOther,
}

func (t LogType) String() string {
	if t == "" {
		return TypeOther.String()
	}
	return string(t)
}

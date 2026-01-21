package gologger

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Logger struct {
	out             io.Writer
	serviceName     string
	typeBgColors    map[LogType]string
	rawLevelIcons   map[Level]string
	levelIconColors map[Level]string
	maxTypeWidth    int
	customTypes     map[LogType]struct{}
}

func NewLogger(out io.Writer, serviceName string, opts ...LoggerOption) *Logger {
	lg := &Logger{
		out:         out,
		serviceName: serviceName,
		typeBgColors: map[LogType]string{
			TypeAuth:  bgGreen,
			TypeHTTP:  bgCyan,
			TypeDB:    bgYellow,
			TypeCache: bgMagenta,
			TypeError: bgRed,
			TypeInfo:  bgWhite,
			TypeDebug: dim,
			TypeOther: bgBlue,
		},
		rawLevelIcons: map[Level]string{
			LevelInfo:    "[i]",
			LevelSuccess: "[v]",
			LevelOk:      "[v]",
			LevelError:   "[e]",
			LevelWarn:    "[w]",
			LevelDebug:   "[d]",
			LevelTrace:   "[t]",
			LevelFatal:   "[!]",
		},
		levelIconColors: map[Level]string{
			LevelInfo:    brightBlue,
			LevelSuccess: brightGreen,
			LevelOk:      brightGreen,
			LevelError:   brightRed,
			LevelWarn:    brightYellow,
			LevelDebug:   dim,
			LevelTrace:   cyan,
			LevelFatal:   bgRed + brightRed,
		},
		customTypes: map[LogType]struct{}{},
	}

	for _, opt := range opts {
		if opt != nil {
			opt(lg)
		}
	}
	lg.updateMaxTypeWidth()
	return lg
}

func (lg *Logger) updateMaxTypeWidth() {
	max := 0
	for _, lt := range defaultLogTypes {
		length := len(lt.String())
		if length > max {
			max = length
		}
	}
	for lt := range lg.customTypes {
		length := len(lt.String())
		if length > max {
			max = length
		}
	}
	lg.maxTypeWidth = max
}

func (lg *Logger) colorizeType(lt LogType) string {
	name := lt.String()

	bg, ok := lg.typeBgColors[lt]
	if !ok {
		bg = bgWhite
	}
	padded := fmt.Sprintf("%-*s", lg.maxTypeWidth, name)
	return fmt.Sprintf("%s %s %s", bg, padded, reset)
}

func (lg *Logger) colorizeIcon(level Level) string {
	raw, ok := lg.rawLevelIcons[level]
	if !ok {
		fallback := "(" + strings.ToLower(level.String()) + ")"
		padded := fmt.Sprintf("%-4s", fallback)
		return white + padded + reset
	}
	color := lg.levelIconColors[level]
	visiblePadded := fmt.Sprintf("%-4s", raw)
	return color + visiblePadded + reset
}

func (lg *Logger) Log(level Level, lt LogType, message, requestID string) {
	timestamp := time.Now().Format("15:04:05.000")
	typeColored := lg.colorizeType(lt)
	iconColored := lg.colorizeIcon(level)

	parts := []string{
		timestamp,
		lg.serviceName,
		typeColored,
		iconColored,
		message,
	}

	if requestID != "" {
		parts = append(parts, fmt.Sprintf("[%s]", requestID))
	}

	fmt.Fprintln(lg.out, strings.Join(parts, " "))
}

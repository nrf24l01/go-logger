package gologger

type LoggerOption func(*Logger)

func WithCustomTypes(types ...LogType) LoggerOption {
	return func(lg *Logger) {
		if lg.customTypes == nil {
			lg.customTypes = map[LogType]struct{}{}
		}
		for _, t := range types {
			if t == "" {
				continue
			}
			lg.customTypes[t] = struct{}{}
		}
	}
}

// WithTypeColors sets background/color strings for given LogType keys.
// It also registers any provided LogType values as custom types so they
// are accounted for in width calculations.
func WithTypeColors(colors map[LogType]string) LoggerOption {
	return func(lg *Logger) {
		if lg.typeBgColors == nil {
			lg.typeBgColors = map[LogType]string{}
		}
		if lg.customTypes == nil {
			lg.customTypes = map[LogType]struct{}{}
		}
		for lt, col := range colors {
			if lt == "" {
				continue
			}
			lg.typeBgColors[lt] = col
			lg.customTypes[lt] = struct{}{}
		}
	}
}

// WithTypeColor is a convenience helper to set a single type color.
func WithTypeColor(lt LogType, color string) LoggerOption {
	return WithTypeColors(map[LogType]string{lt: color})
}

package gologger

const (
	reset     = "\033[0m"
	bold      = "\033[1m"
	dim       = "\033[2m"
	red       = "\033[31m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	white     = "\033[37m"
	bgRed     = "\033[41m"
	bgGreen   = "\033[42m"
	bgYellow  = "\033[43m"
	bgBlue    = "\033[44m"
	bgMagenta = "\033[45m"
	bgCyan    = "\033[46m"
	bgWhite   = "\033[47m"

	brightRed    = "\033[91m"
	brightGreen  = "\033[92m"
	brightYellow = "\033[93m"
	brightBlue   = "\033[94m"
)

// Exported aliases for package users
const (
	Reset   = reset
	Bold    = bold
	Dim     = dim
	Red     = red
	Green   = green
	Yellow  = yellow
	Blue    = blue
	Magenta = magenta
	Cyan    = cyan
	White   = white

	BgRed     = bgRed
	BgGreen   = bgGreen
	BgYellow  = bgYellow
	BgBlue    = bgBlue
	BgMagenta = bgMagenta
	BgCyan    = bgCyan
	BgWhite   = bgWhite

	BrightRed    = brightRed
	BrightGreen  = brightGreen
	BrightYellow = brightYellow
	BrightBlue   = brightBlue
)

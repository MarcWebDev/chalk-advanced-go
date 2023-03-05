package ChalkAdvanced

import "fmt"

var (
	ColorReset = "\033[0m"

	// Special //
	SpecialBold      = "\033[1m"
	SpecialUnderline = "\033[4m"

	// Text colors //
	ColorBlack  = "\033[30m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[37m"
	ColorWhite  = "\033[97m"

	// Background colors //
	BackgroundColorBlack  = "\033[40m"
	BackgroundColorRed    = "\033[41m"
	BackgroundColorGreen  = "\033[42m"
	BackgroundColorYellow = "\033[43m"
	BackgroundColorBlue   = "\033[44m"
	BackgroundColorPurple = "\033[45m"
	BackgroundColorCyan   = "\033[46m"
	BackgroundColorGray   = "\033[47m"
	BackgroundColorWhite  = "\033[107m"
)

func color(color string, s any) string {
	switch s.(type) {
	case string:
		return color + s.(string) + ColorReset
	default:
		return color + fmt.Sprint(s) + ColorReset
	}
}

func Black(message string) string {
	return color(ColorBlack, message)
}

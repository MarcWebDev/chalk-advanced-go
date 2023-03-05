package ChalkAdvanced

import "fmt"

var (
	Reset = "\033[0m"

	SpecialBold      = "\033[1m"
	SpecialUnderline = "\033[4m"

	ColorBlack  = "\033[30m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[37m"
	ColorWhite  = "\033[97m"

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
		return color + s.(string) + Reset
	default:
		return color + fmt.Sprint(s) + Reset
	}
}

func Black(message string) string {
	return color(ColorBlack, message)
}

func Red(message string) string {
	return color(ColorRed, message)
}

func Green(message string) string {
	return color(ColorGreen, message)
}

func Yellow(message string) string {
	return color(ColorYellow, message)
}

func Blue(message string) string {
	return color(ColorBlue, message)
}

func Purple(message string) string {
	return color(ColorPurple, message)
}

func Cyan(message string) string {
	return color(ColorCyan, message)
}

func Gray(message string) string {
	return color(ColorGray, message)
}

func White(message string) string {
	return color(ColorWhite, message)
}
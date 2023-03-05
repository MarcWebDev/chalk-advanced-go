package ChalkAdvanced

import "fmt"

var (
	Reset = "\033[0m"

	SpecialBold          = "\033[1m"
	SpecialUnderline     = "\033[4m"
	SpecialDim           = "\033[2m"
	SpecialItalic        = "\033[3m"
	SpecialInverse       = "\033[7m"
	SpecialHide          = "\033[8m"
	SpecialStrikethrough = "\033[9m"

	ColorBlack   = "\033[30m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorGray    = "\033[37m"
	ColorWhite   = "\033[97m"

	ColorRedBright     = "\033[91m"
	ColorGreenBright   = "\033[92m"
	ColorYellowBright  = "\033[93m"
	ColorBlueBright    = "\033[94m"
	ColorMagentaBright = "\033[95m"
	ColorCyanBright    = "\033[96m"
	ColorWhiteBright   = "\033[97m"

	BackgroundColorBlack   = "\033[40m"
	BackgroundColorRed     = "\033[41m"
	BackgroundColorGreen   = "\033[42m"
	BackgroundColorYellow  = "\033[43m"
	BackgroundColorBlue    = "\033[44m"
	BackgroundColorMagenta = "\033[45m"
	BackgroundColorCyan    = "\033[46m"
	BackgroundColorGray    = "\033[47m"
	BackgroundColorWhite   = "\033[107m"

	BackgroundColorRedBright     = "\033[101m"
	BackgroundColorGreenBright   = "\033[102m"
	BackgroundColorYellowBright  = "\033[103m"
	BackgroundColorBlueBright    = "\033[104m"
	BackgroundColorMagentaBright = "\033[105m"
	BackgroundColorCyanBright    = "\033[106m"
	BackgroundColorWhiteBright   = "\033[107m"
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

func Magenta(message string) string {
	return color(ColorMagenta, message)
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

func BgBlack(message string) string {
	return color(BackgroundColorBlack, message)
}

func BgRed(message string) string {
	return color(BackgroundColorRed, message)
}

func BgGreen(message string) string {
	return color(BackgroundColorGreen, message)
}

func BgYellow(message string) string {
	return color(BackgroundColorYellow, message)
}

func BgBlue(message string) string {
	return color(BackgroundColorBlue, message)
}

func BgMagenta(message string) string {
	return color(BackgroundColorMagenta, message)
}

func BgCyan(message string) string {
	return color(BackgroundColorCyan, message)
}

func BgGray(message string) string {
	return color(BackgroundColorGray, message)
}

func BgWhite(message string) string {
	return color(BackgroundColorWhite, message)
}

func Bold(message string) string {
	return color(SpecialBold, message)
}

func Underline(message string) string {
	return color(SpecialUnderline, message)
}

func Dim(message string) string {
	return color(SpecialDim, message)
}

func Italic(message string) string {
	return color(SpecialItalic, message)
}

func Inverse(message string) string {
	return color(SpecialInverse, message)
}

func Hide(message string) string {
	return color(SpecialHide, message)
}

func Strikethrough(message string) string {
	return color(SpecialStrikethrough, message)
}

func RedBright(message string) string {
	return color(ColorRedBright, message)
}

func GreenBright(message string) string {
	return color(ColorGreenBright, message)
}

func YellowBright(message string) string {
	return color(ColorYellowBright, message)
}

func BlueBright(message string) string {
	return color(ColorBlueBright, message)
}

func MagentaBright(message string) string {
	return color(ColorMagentaBright, message)
}

func CyanBright(message string) string {
	return color(ColorCyanBright, message)
}

func WhiteBright(message string) string {
	return color(ColorWhiteBright, message)
}

func bgRedBright(message string) string {
	return color(BackgroundColorRedBright, message)
}

func bgGreenBright(message string) string {
	return color(BackgroundColorGreenBright, message)
}

func bgYellowBright(message string) string {
	return color(BackgroundColorYellowBright, message)
}

func bgBlueBright(message string) string {
	return color(BackgroundColorBlueBright, message)
}

func bgMagentaBright(message string) string {
	return color(BackgroundColorMagentaBright, message)
}

func bgCyanBright(message string) string {
	return color(BackgroundColorCyanBright, message)
}

func bgWhiteBright(message string) string {
	return color(BackgroundColorWhiteBright, message)
}

package ChalkAdvanced

import "log"

func Black(message string) {
	log.Printf("\\x1b[30m%v\\x1b[0m", message)
}

package ChalkAdvanced

import (
	"fmt"
)

func Black(message string) {
	fmt.Printf("\\x1b[30m%v\\x1b[0m", message)
}

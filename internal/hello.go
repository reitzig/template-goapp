package internal

import (
	"fmt"
)

func Hello(config HelloConfig, hailee string) string {
	if config.Dear.Value {
		return fmt.Sprintf("Hello, dear %s!", hailee)
	} else {
		return fmt.Sprintf("Hello, %s!", hailee)
	}
}

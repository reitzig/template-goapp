package internal

import (
	"fmt"
)

func Hello(dear bool, hailee string) string {
	if dear {
		return fmt.Sprintf("Hello, dear %s!", hailee)
	} else {
		return fmt.Sprintf("Hello, %s!", hailee)
	}
}

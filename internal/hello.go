package internal

import (
	"fmt"
	"log"
)

func Hello(dear bool, hailee string) {
	if dear {
		log.Println(fmt.Sprintf("Hello, dear %s!", hailee))
	} else {
		log.Println(fmt.Sprintf("Hello, %s!", hailee))
	}
}

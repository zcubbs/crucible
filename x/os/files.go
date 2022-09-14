package os

import (
	"fmt"
	"log"
	"os"
)

func Chmod(path string) {
	stats, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File permissions before: %s\n", stats.Mode())
	err = os.Chmod(path, 0700)
	if err != nil {
		log.Fatal(err)
	}

	stats, err = os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File permissions after: %s\n", stats.Mode())
}

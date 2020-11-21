package main

import (
	"os"

	"github.com/knightso/trimfaketime"
)

func main() {
	trimfaketime.TrimFakeTime(os.Stdout, os.Stdin)
}

package main

import (
	"os"
	"quadchecker"
)

func main() {

	quadchecker.CheckForExecutables()

	lines := quadchecker.GetShape(os.Stdin)

	if len(lines) > 0 {
		quadMatches := quadchecker.Check(lines)

		quadchecker.PrintResult(quadMatches)
	}

}

package version

import (
	"fmt"
	"os"
)

func ShowVersionDetect(version string) {
	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "-V" || arg == "--version" || arg == "version" {
			fmt.Println(version)
			os.Exit(0)
		}
	}
}

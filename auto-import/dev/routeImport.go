package main

import (
	"fmt"
	"os"
)

func routeImport() {
	dirs := map[string]struct{}{}
	collectPackages("./route", "", dirs)
	filePath := "./routeImport.go"
	content := "package main\nimport (\n"
	for dir := range dirs {
		content += fmt.Sprintf("\t_ \"%s\"\n", "auto-import/route"+dir)
	}
	content += ")\n"
	os.WriteFile(filePath, []byte(content), 0644)
}

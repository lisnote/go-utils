package main

import (
	"libs/errorx"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func collectPackages(path string, prefix string, dirs map[string]struct{}) {
	entries, err := os.ReadDir(path)
	errorx.Fatal(err, "读取文件目录失败")

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	for _, entry := range entries {
		name := entry.Name()
		switch {
		case strings.HasPrefix(name, "_"):
			continue
		case entry.IsDir():
			collectPackages(filepath.Join(path, name), prefix+"/"+name, dirs)
		case strings.HasSuffix(name, ".go"):
			dirs[prefix] = struct{}{}
		}
	}
}

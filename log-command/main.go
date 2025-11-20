package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	fullname := filepath.Base(exe)
	extname := filepath.Ext(fullname)
	name := strings.TrimSuffix(fullname, extname)
	logFile, _ := os.OpenFile(filepath.Join(dir, name+".log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(logFile)

	log.Println(fullname, strings.Join(os.Args[1:], " "))

	cmd := exec.Command(name+"-proxy"+extname, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Run()
}

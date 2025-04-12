package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	} else if err != nil {
		return err
	}
	return nil
}

func Cp(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func touch(path, filename string) (*os.File, error) { return os.Create(filepath.Join(path, filename)) }

func CreateEntryPoint(path, filename string) {
	f, err := touch(path, filename)
	if err != nil {
		Debugf("Unable to create file: %v\n", err)
	}
	defer f.Close()
	f.WriteString(`package main

func main() {

}
`)
}

func LogFatalAndExit(err error, exitCode int) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "FATAL ERROR: %v\n", err)
	os.Exit(exitCode)
}

func DefineSeparator(s string) string {
	commaWithSpace := ", "
	comma := ","
	space := " "

	if strings.Contains(s, commaWithSpace) {
		return commaWithSpace
	} else if strings.Contains(s, comma) {
		return comma
	} else if strings.Contains(s, space) {
		return space
	}

	return ""
}

func IsContain(s string, slice []string) bool {
	for _, elem := range slice {
		if s == elem {
			return true
		}
	}
	return false
}

func GoModInit(moduleName string) *exec.Cmd {
	args := []string{
		"mod",
		"init",
		moduleName,
	}
	cmd := exec.Command("go", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func IsGoInstalled() (string, bool) {
	path, err := exec.LookPath("go")
	if err == nil {
		return path, true
	}

	if runtime.GOOS == "windows" {
		defaultPaths := []string{
			`C:\Program Files\Go\bin\go.exe`,
			`C:\Program Files (x86)\Go\bin\go.exe`,
		}

		for _, p := range defaultPaths {
			if _, err := os.Stat(p); err == nil {
				return p, true
			}
		}
	}

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if path, err := exec.LookPath("go"); err == nil {
			return path, true
		}
	}

	return "", false
}

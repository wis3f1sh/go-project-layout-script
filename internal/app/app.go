package app

import (
	"flag"
	"fmt"
	"go-project-layout-script/internal/globals"
	"go-project-layout-script/internal/utils"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Run() {
	pwd, err := os.Getwd()
	utils.LogFatalAndExit(err, 1)

	var excludedDirsFlag string
	var onlyDirsFlag string
	var pathFlag string

	flag.StringVar(&excludedDirsFlag, "exclude-dirs", "", "directory names to be excluded")
	flag.StringVar(&onlyDirsFlag, "only-dirs", "", "directory names to create")
	flag.StringVar(&pathFlag, "path", pwd, "defines path where createad project layout")

	flag.Parse()

	if pathFlag != "" {
		if !filepath.IsAbs(pathFlag) {
			err := utils.Mkdir(pathFlag)
			pathFlag = filepath.Join(pwd, pathFlag)
			utils.LogFatalAndExit(err, 1)
		} else {
			err := utils.Mkdir(pathFlag)
			utils.LogFatalAndExit(err, 1)
		}
	}

	if excludedDirsFlag == "" && onlyDirsFlag == "" {
		for _, dir := range globals.Dirs {
			if dir == "cmd" {
				moduleName := filepath.Base(pathFlag)
				path := filepath.Join(pathFlag, "cmd", moduleName)
				utils.ToReportf("path: %s\n", path)
				err := utils.Mkdir(path)
				utils.LogFatalAndExit(err, 1)

				utils.CreateEntryPoint(path, "main.go")
				utils.LogFatalAndExit(err, 1)
			} else {
				path := filepath.Join(pathFlag, dir)
				err := utils.Mkdir(path)
				utils.LogFatalAndExit(err, 1)
			}
		}
		utils.ToReport(true, "Project layout was successfully created...")
	}

	if excludedDirsFlag != "" && onlyDirsFlag != "" {
		fmt.Fprintln(os.Stderr, "Can't use both flags. Just only one...")
	}

	if excludedDirsFlag != "" {
		sep := utils.DefineSeparator(excludedDirsFlag)
		excludedDirs := strings.Split(excludedDirsFlag, sep)

		for _, dir := range globals.Dirs {
			if utils.IsContain(dir, excludedDirs) {
				continue
			} else {
				if dir == "cmd" {
					moduleName := filepath.Base(pathFlag)
					path := filepath.Join(pathFlag, "cmd", moduleName)
					err := utils.Mkdir(path)
					utils.LogFatalAndExit(err, 1)

					utils.CreateEntryPoint(path, "main.go")
					utils.LogFatalAndExit(err, 1)
				} else {
					path := filepath.Join(pathFlag, dir)
					err := utils.Mkdir(path)
					utils.LogFatalAndExit(err, 1)
				}
			}
		}
		utils.ToReport(true, "Project layout was successfully created...")
	}

	if onlyDirsFlag != "" {
		sep := utils.DefineSeparator(onlyDirsFlag)
		onlyDirs := strings.Split(onlyDirsFlag, sep)

		for _, dir := range globals.Dirs {
			if utils.IsContain(dir, onlyDirs) {
				if dir == "cmd" {
					moduleName := filepath.Base(pathFlag)
					path := filepath.Join(pathFlag, "cmd", moduleName)
					utils.ToReportf("path: %s\n", path)
					err := utils.Mkdir(path)
					utils.LogFatalAndExit(err, 1)

					utils.CreateEntryPoint(path, "main.go")
					utils.LogFatalAndExit(err, 1)
				} else {
					path := filepath.Join(pathFlag, dir)
					err := utils.Mkdir(path)
					utils.LogFatalAndExit(err, 1)
				}
			} else {
				continue
			}
		}
		utils.ToReport(true, "Project layout was successfully created...")
	}

	utils.ToReport(true, "Checking if go is installed...")
	path, isInstalled := utils.IsGoInstalled()
	if isInstalled {
		utils.ToReportf("[✓] go is installed: %s\n", path)

		moduleName := filepath.Base(pathFlag)
		cmd := utils.GoModInit(moduleName)

		if err := cmd.Run(); err != nil {
			utils.ToReportf("Error executing command 'go mod init %s': %v\n", moduleName, err)
		}

		dst := filepath.Join(pathFlag, "go.mod")
		if err := utils.Cp("go.mod", dst); err != nil {
			utils.ToReportf("Error copying go.mod: %v\n", err)
		}

		if err := os.Remove("go.mod"); err != nil {
			utils.ToReportf("Error removing go.mod: %v\n", err)
		}
	} else {
		utils.ToReport(true, "[✗] go not found in system")

		switch runtime.GOOS {
		case "windows":
			utils.ToReport(true, `
Download MSI file of Go from official site and install it:
https://go.dev/dl/
Install in default directory`)
		case "linux":
			utils.ToReport(true, `
Download tar archive from official site:
https://go.dev/dl/

Remove any previous Go installation
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz

Add /usr/local/go/bin to the PATH environment variable
$ export PATH=$PATH:/usr/local/go/bin`)
		case "darwin":
			utils.ToReport(true, `
Download pkg file from official site and install it:
https://go.dev/dl/`)
		}
	}
}

// Copyright 2015 ThoughtWorks, Inc.
// Copyright (C) 2012 vova616 <vova616@gmail.com>

// This file is part of Gauge-Screenshot.

// Gauge-Screenshot is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge-Screenshot is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge-Screenshot.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/getgauge/common"
	"github.com/getgauge/gauge_screenshot/version"
)

const (
	CGO_ENABLED     = "CGO_ENABLED"
	GOARCH          = "GOARCH"
	GOOS            = "GOOS"
	X86             = "386"
	X86_64          = "amd64"
	darwin          = "darwin"
	linux           = "linux"
	windows         = "windows"
	bin             = "bin"
	deploy          = "deploy"
	gaugeScreenshot = "gauge_screenshot"
	CC              = "CC"
)

var platformEnvs = []map[string]string{
	map[string]string{GOARCH: X86, GOOS: darwin, CGO_ENABLED: "0"},
	map[string]string{GOARCH: X86_64, GOOS: darwin, CGO_ENABLED: "0"},
	map[string]string{GOARCH: X86, GOOS: linux, CGO_ENABLED: "0"},
	map[string]string{GOARCH: X86_64, GOOS: linux, CGO_ENABLED: "0"},
	map[string]string{GOARCH: X86, GOOS: windows, CC: "i586-mingw32-gcc", CGO_ENABLED: "1"},
	map[string]string{GOARCH: X86_64, GOOS: windows, CC: "x86_64-w64-mingw32-gcc", CGO_ENABLED: "1"},
}
var cCompile = flag.Bool("all", false, "Cross compile")

func main() {
	flag.Parse()
	if *cCompile {
		crossCompile()
	} else {
		compile()
	}
}

func crossCompile() {
	for _, platformEnv := range platformEnvs {
		setEnv(platformEnv)
		log.Printf("Compiling for platform => OS:%s ARCH:%s \n", platformEnv[GOOS], platformEnv[GOARCH])
		compile()
		createDistro()
	}
}

func setEnv(envVariables map[string]string) {
	for k, v := range envVariables {
		os.Setenv(k, v)
	}
}

func compile() {
	runProcess("go", "get", "./...")
	runProcess("go", "build", "-o", getExecutablePath(gaugeScreenshot))
}

func createDistro() {
	packageName := fmt.Sprintf("%s-%s-%s.%s", gaugeScreenshot, version.Version, getGOOS(), getArch())
	distroDir := filepath.Join(deploy, packageName)
	copyPluginFiles(distroDir)
	createZipFromUtil(deploy, packageName)
	os.RemoveAll(distroDir)
}

func createZipFromUtil(dir, name string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Chdir(filepath.Join(dir, name))
	runProcess("zip", "-r", filepath.Join("..", name+".zip"), ".")
	os.Chdir(wd)
}

func copyPluginFiles(destDir string) {
	files := make(map[string]string)
	files[getExecutablePath(gaugeScreenshot)] = ""
	copyFiles(files, destDir)
}

// key will be the source file and value will be the target
func copyFiles(files map[string]string, installDir string) {
	for src, dst := range files {
		base := filepath.Base(src)
		installDst := filepath.Join(installDir, dst)
		log.Printf("Copying %s -> %s\n", src, installDst)
		stat, err := os.Stat(src)
		if err != nil {
			panic(err)
		}
		if stat.IsDir() {
			_, err = common.MirrorDir(src, installDst)
		} else {
			err = common.MirrorFile(src, filepath.Join(installDst, base))
		}
		if err != nil {
			panic(err)
		}
	}
}

func runProcess(command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Execute %v\n", cmd.Args)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func getExecutablePath(file string) string {
	return filepath.Join(getBinDir(), getExecutableName(file))
}

func getBinDir() string {
	return filepath.Join(bin, fmt.Sprintf("%s_%s", getGOOS(), getGOARCH()))
}

func getGOARCH() string {
	goArch := os.Getenv(GOARCH)
	if goArch == "" {
		goArch = runtime.GOARCH
	}
	return goArch
}

func getExecutableName(file string) string {
	if getGOOS() == windows {
		return file + ".exe"
	}
	return file
}
func getGOOS() string {
	goOS := os.Getenv(GOOS)
	if goOS == "" {
		goOS = runtime.GOOS
	}
	return goOS
}
func getArch() string {
	arch := getGOARCH()
	if arch == X86 {
		return "x86"
	}
	return "x86_64"
}

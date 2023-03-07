package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func setSystemHosts(content string) {
	hosts := getHostPath()
	stringBytes := []byte(content)
	if err := ioutil.WriteFile(hosts, stringBytes, os.ModePerm); err != nil {
		println(err)
	}
}

func getHostPath() string {
	hostsPath := "/etc/hosts"
	if runtime.GOOS == "windows" {
		hostsPath = getWinSystemDir()
		hostsPath = filepath.Join(hostsPath, "system32", "drivers", "etc", "hosts")
	}
	return hostsPath
}

func getUserHome() string {
	home := ""
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}

	return home
}

func getWinSystemDir() string {
	dir := ""
	if runtime.GOOS == "windows" {
		dir = os.Getenv("windir")
	}

	return dir
}

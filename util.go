package main

import (
	"bytes"
	"os"
	"os/exec"
)

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func FileExists(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}

func DirExists(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}

func Run(dir string, exe string, arg ...string) (stdout *bytes.Buffer, stderr *bytes.Buffer, err error) {
	cmd := exec.Command(exe, arg...)
	var errbuf bytes.Buffer
	var outbuf bytes.Buffer
	cmd.Dir = dir
	cmd.Stderr = &errbuf
	cmd.Stdout = &outbuf
	err = cmd.Run()
	return &outbuf, &errbuf, err
}

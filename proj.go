package main

import (
	"fmt"
	"os"
)

type Project struct {
	Path string
}

func NewProject(path string) *Project {
	path = TrimRightPathSep(path)
	return &Project{Path: path}
}

func TrimRightPathSep(path string) string {
	n := len(path)
	if n > 0 {
		if path[n-1] == os.PathSeparator {
			path = path[:n-1]
		}
	}
	return path
}

func (p *Project) GetWriteDir() (string, error) {
	wa := fmt.Sprintf(".%cdot.geist%c%s%c", os.PathSeparator, os.PathSeparator, p.Path, os.PathSeparator)
	err := os.MkdirAll(wa, 0777)
	return wa, err
}

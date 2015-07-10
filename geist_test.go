package main

import (
	"fmt"
	"strings"
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

// A gest script is #! runnable, no manual compilation step,
// not even "go run", is needed. The gest executable
// will read the script, add the boilerplate for you:
// the package main and the main() method that wraps
// the code in the script will be written for you.
// Then gest will compile and run it for you.

// Caching: the compilation actually only happens if the script
// has been changed (we cached the prior version; so
// we offer super fast execution of gest scripts as
// a result.

// Compilation happens in an isolated GOROOT/GOPATH
// environment under $HOME/.gest by default (settable
// with GEST_HOME env variable). This means you can
// run gest starting in any directory, and it will
// ignore any other golang code, GOROOT/GOPATH settings.

// A given gest script should have an implicit main
// and several implicit imports (fmt, math, io),
// possibly other standard packages.

// Additional imports can be added anywhere, no problem.
// gest will collect them all to the top of the file.

// structs and functions can be defined in a gest
// file, and the implicit main stops at the point
// of the first (if any) struct or fuction definition.

// The gest executable will check for a cached
// (already compiled) version of the file first, and
// run that if available. Otherwise it should
// build it, save it, run it.

// gest will add the main() function declaraction
// to the script at hand, wrapping the given script
// instructions.

var built bool

func Build() {
	if !built {
		outbuf, errbuf, err := Run(".", "go", "build")
		if err != nil {
			panic(err)
		}
		fmt.Printf("go build succeeded, stdout='%s', stderr='%s', err='%v'\n", outbuf.String(), errbuf.String(), err)
		built = true
	}
}

func TestGeistRuns(t *testing.T) {
	Build()
	cv.Convey("Given the geist executable, geist should be able to receive the text body of the file it is run from, when using '#!/usr/bin/env geist' (or local test equivalent) at the top of the script", t, func() {

		outbuf, errbuf, err := Run("testdata", "./example.g")
		fmt.Printf("\n err on cmd.Run() was '%v'\n", err)
		fmt.Printf("\n Stderr was '%s'\n", errbuf.String())
		fmt.Printf("\n Stdout was '%s'\n", outbuf.String())
		cv.So(err, cv.ShouldEqual, nil)
		cv.So(strings.Contains(outbuf.String(), "Welcome to Geist"), cv.ShouldEqual, true)
	})
}

func TestGeistRequiresArg1(t *testing.T) {
	Build()
	cv.Convey("Given the geist executable, when geist is invoked, it must require one argument (the path to the script it should execute) or fail complaining a path is required.", t, func() {
		_, errbuf, err := Run(".", "./geist")
		cv.So(err, cv.ShouldNotEqual, nil)
		cv.So(errbuf.String(), cv.ShouldEqual, "geist error: missing path to script to run as argument to geist.")
	})
}

func TestEmptyGeistSciptRuns(t *testing.T) {
	Build()
	cv.Convey("Given an empty geist scipt, geist should run it without issue.", t, func() {

		outbuf, errbuf, err := Run("testdata", "./empty.g")
		fmt.Printf("\n err on cmd.Run() was '%v'\n", err)
		fmt.Printf("\n Stderr was '%s'\n", errbuf.String())
		fmt.Printf("\n Stdout was '%s'\n", outbuf.String())
		cv.So(err, cv.ShouldEqual, nil)
		cv.So(strings.Contains(outbuf.String(), ""), cv.ShouldEqual, true)
	})
}

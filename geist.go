package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Printf("geist main() sees os.Args = '%#v'\n", os.Args)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "geist error: missing path to script to run as argument to geist.")
		os.Exit(1)
	}
	if !FileExists(os.Args[1]) {
		fmt.Fprintf(os.Stderr, "geist error: file not found: '%s'.")
		os.Exit(1)
	}

	origPath := os.Args[1]
	newPath, err := GeistFileToGoFile(origPath)
	panicOn(err)

	outb, errb, err := Run(".", "go", "run", newPath)
	if err != nil {
		panic(fmt.Errorf("geist error: could not run script '%s': '%v'. Stderr: '%s'. Stdout: '%s'.", origPath, err, errb.String(), outb.String()))
	}
	os.Stdout.Write(outb.Bytes())
	os.Stderr.Write(errb.Bytes())
}

func GeistFileToGoFile(geistPath string) (string, error) {
	f, err := os.Open(geistPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// Can't work in the cwd directory.
	// make a new project dir (or run a cached one)
	p := NewProject(geistPath)
	writeDir, err := p.GetWriteDir()
	panicOn(err)

	writePath := writeDir + "main.go"

	of, err := os.OpenFile(writePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return "", fmt.Errorf("GeistFileToGoFile(): error durring OpenFile on '%s': '%v'",
			writePath, err)
	}
	defer of.Close()

	// preamble
	fmt.Fprintf(of, `package main`)
	fmt.Fprintf(of, stdlibImports)
	fmt.Fprintf(of, `func main() {
`)

	// we've got to remove the first line from f
	_, err = r.ReadString('\n')
	if err == io.EOF {
		return "", fmt.Errorf("not enough input in '%s': hit EOF before finding end of first line", geistPath)
	}
	panicOn(err)

	// copy the rest of the script.g -> main.go file
	_, err = io.Copy(of, r)
	if err != nil {
		return "", fmt.Errorf("GeistFileToGoFile(): error during creation of .go file in '%s': '%v'", writePath, err)
	}

	// postamble
	fmt.Fprintf(of, `
}
`)

	return writePath, nil
}

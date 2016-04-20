package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kisom/goutils/lib"
)

var generators = map[string]func(*opts) error{
	"c":   writeCSource,
	"h":   writeCSource,
	"fth": writeForthSource,
	"cc":  writeCCSource,
	"hh":  writeCCSource,
}

var fileTypes = map[string]string{
	"c":   "C source file",
	"h":   "C header file",
	"fth": "Forth definitions",
	"cc":  "C++ source file",
	"hh":  "C++ header file",
}

func init() {
	flag.Usage = func() { usage(os.Stdout) }
}

func usage(w io.Writer) {
	fmt.Fprintf(w, `Usage: %s [-f] [-l license] language file

	Flags:
		-f		Force writing the file, even if it exists.
		-l license	Specify the license to use.

	Creates a new source file with some basic templating. Valid
	languages are
`, lib.ProgName())

	for lang, descr := range fileTypes {
		fmt.Fprintf(w, "\t\t+ %s: %s\n", lang, descr)
	}

	fmt.Fprintf(w, "\n\tValid licenses are\n")
	for license := range licenses {
		fmt.Fprintf(w, "\t\t+ %s\n", license)
	}
}

func main() {
	var opts opts
	var help bool

	flag.BoolVar(&opts.force, "f", false, "Force write the file, even if it exists.")
	flag.BoolVar(&help, "h", false, "Display short usage message.")
	flag.StringVar(&opts.license, "l", "mit", "`License` for this file.")
	flag.Parse()

	if help {
		usage(os.Stderr)
		os.Exit(0)
	}

	if flag.NArg() != 2 {
		usage(os.Stderr)
		os.Exit(1)
	}

	lang := flag.Arg(0)
	gen, ok := generators[lang]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown license %s.\n", lang)
		usage(os.Stderr)
		os.Exit(1)
	}

	opts.body, ok = sources[lang]
	if !ok {
		panic(`oops: forgot to define template for language ` + lang)
	}

	opts.path = flag.Arg(1) + "." + lang
	err := gen(&opts)
	if err != nil {
		lib.Err(2, err, "failed to create source.\n")
	}
}

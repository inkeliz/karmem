package main

import (
	"errors"
	"flag"
	"fmt"
	"karmem.org/cmd/karmem/kmgen"
	"karmem.org/cmd/karmem/kmparser"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "build":
		f := Build{}
		f.Parse()
		if err := f.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

type Language string

func (l Language) Generator() kmgen.Generator {
	switch l {
	case LanguageGolang:
		return kmgen.GolangGenerator()
	case LanguageAssemblyScript:
		return kmgen.AssemblyScriptGenerator()
	case LanguageZig:
		return kmgen.ZigGenerator()
	case LanguageSwift:
		return kmgen.SwiftGenerator()
	case LanguageC:
		return kmgen.CGenerator()
	default:
		return nil
	}
}

const (
	LanguageGolang         Language = "golang"
	LanguageZig            Language = "zig"
	LanguageAssemblyScript Language = "assemblyscript"
	LanguageSwift          Language = "swift"
	LanguageC              Language = "c"
)

var (
	Languages       = [...]Language{LanguageGolang, LanguageZig, LanguageAssemblyScript, LanguageSwift, LanguageC}
	LanguagesFormat = [...]string{".go", ".zig", ".ts", ".swift", ".h"}
)

type Build struct {
	language [len(Languages)]bool
	output   string
	input    string
}

func (b *Build) Parse() {
	flags := flag.NewFlagSet("build", flag.ExitOnError)
	for i, l := range Languages {
		flags.BoolVar(&b.language[i], string(l), false, fmt.Sprintf("Enable geneartion for %s language.", l))
	}
	flags.StringVar(&b.output, "o", ".", "Output directory path.")
	flags.Parse(flag.Args()[1:])

	b.input = flags.Arg(0)
}

func (b *Build) Generator() (v kmgen.Generator, err error) {
	for i, l := range Languages {
		if b.language[i] {
			if v != nil {
				return nil, errors.New("multiple languages is not supported")
			}
			v = l.Generator()
		}
	}
	if v == nil {
		return nil, errors.New("missing language. Please, specify one output language (such as --golang)")
	}
	return v, nil
}

func (b *Build) FileFormat() string {
	for i := range Languages {
		if b.language[i] {
			return LanguagesFormat[i]
		}
	}
	return ""
}

func (b *Build) Execute() error {
	gen, err := b.Generator()
	if err != nil {
		return err
	}

	inputFile, err := os.Open(b.input)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	parsed, err := kmparser.NewReader(b.input, inputFile).Parser()
	if err != nil {
		return err
	}

	if err := gen.Start(parsed); err != nil {
		return err
	}

	if b.output != "." {
		if err := os.MkdirAll(b.output, os.ModePerm); err != nil {
			return err
		}
	}

	outputFile, err := os.Create(filepath.Join(b.output, parsed.Header.Name+"_generated"+b.FileFormat()))
	if err != nil {
		return err
	}

	if err := gen.Save(outputFile); err != nil {
		return err
	}

	return outputFile.Close()
}

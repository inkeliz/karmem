package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"karmem.org/cmd/karmem/kmgen"
	"karmem.org/cmd/karmem/kmparser"
)

// errHelpOnly is an error that only prints the help message and returns no error.
var (
	errHelpOnly = errors.New("ErrHelpOnly")
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: karmem <command> [<args>]")
		fmt.Println("Commands:")
		fmt.Println("  build")
		fmt.Println("    Builds the schema file.")
		fmt.Println("    Use \"build help\" to see the list of available options.")
		os.Exit(1)
	}
	flag.Parse()

	var fn func() error
	switch flag.Arg(0) {
	case "build":
		fn = func() error {
			b, err := NewBuild()
			if err != nil {
				return err
			}
			return b.Execute()
		}
	case "fmt":
		fallthrough
	case "format":
		fn = func() error {
			f, err := NewFormat()
			if err != nil {
				return err
			}
			return f.Execute()
		}
	}

	if fn == nil {
		flag.Usage()
		os.Exit(1)
	}

	if err := fn(); err != nil && err != errHelpOnly {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

type SchemaFile struct {
	input string
}

func (s *SchemaFile) ParseIDL() (*kmparser.Content, error) {
	inputFile, err := os.Open(s.input)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer inputFile.Close()

	info, err := inputFile.Stat()
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, fmt.Errorf("input file is a directory: %s", s.input)
	}

	parsed, err := kmparser.NewReader(s.input, inputFile).Parser()
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

type Build struct {
	SchemaFile
	language  []bool
	output    string
	generator []kmgen.Generator
}

func NewBuild() (b *Build, _ error) {
	b = new(Build)
	flags := flag.NewFlagSet("build", flag.ExitOnError)

	b.language = make([]bool, len(kmgen.Generators))
	for i, g := range kmgen.Generators {
		flags.BoolVar(&b.language[i], g.Language(), false, fmt.Sprintf("Enable geneartion for %s language.", g.Language()))
	}
	flags.StringVar(&b.output, "o", ".", "Output directory path.")
	if err := flags.Parse(flag.Args()[1:]); err != nil {
		return nil, err
	}

	b.input = flags.Arg(0)

	if b.input == "help" {
		flags.Usage()
		return nil, errHelpOnly
	}

	for i, g := range kmgen.Generators {
		if b.language[i] {
			b.generator = append(b.generator, g)
		}
	}

	if len(b.generator) == 0 {
		return nil, errors.New("missing language. Please, specify one output language (such as --golang)")
	}

	return b, nil
}

func (b *Build) Execute() error {
	parsed, err := b.ParseIDL()
	if err != nil {
		return err
	}

	for _, gen := range b.generator {
		compiler, err := gen.Start(parsed)
		if err != nil {
			return err
		}

		if b.output != "." {
			if err := os.MkdirAll(b.output, os.ModePerm); err != nil {
				return err
			}
		}

		for i, c := range compiler.Template {
			var buffer bytes.Buffer

			for _, n := range compiler.Modules {
				if err := c.ExecuteTemplate(&buffer, n, kmgen.TemplateData{Content: parsed}); err != nil {
					return err
				}
			}

			outputFile, err := os.Create(filepath.Join(b.output, parsed.Module+"_generated"+gen.Extensions()[i]))
			if err != nil {
				return err
			}

			if err := gen.Finish(outputFile, &buffer); err != nil {
				return err
			}

			if err := outputFile.Close(); err != nil {
				return err
			}
		}
	}

	return nil
}

type Format struct {
	Build
	save bool
}

func NewFormat() (f *Format, _ error) {
	f = new(Format)
	flags := flag.NewFlagSet("build", flag.ExitOnError)

	f.language = make([]bool, len(kmgen.Generators))
	flags.BoolVar(&f.save, "s", false, "Save and override the original schema file.")
	if err := flags.Parse(flag.Args()[1:]); err != nil {
		return nil, err
	}

	f.input = flags.Arg(0)

	if f.input == "help" {
		flags.Usage()
		return nil, errHelpOnly
	}

	return f, nil
}

func (b *Format) Execute() error {
	parsed, err := b.ParseIDL()
	if err != nil {
		return err
	}

	gen := kmgen.KarmemSchemaGenerator()
	compiler, err := gen.Start(parsed)
	if err != nil {
		return err
	}

	for _, c := range compiler.Template {
		var buffer bytes.Buffer

		for _, n := range compiler.Modules {
			if err := c.ExecuteTemplate(&buffer, n, kmgen.TemplateData{Content: parsed}); err != nil {
				return err
			}
		}

		outputFile := os.Stdout
		if b.save {
			outputFile, err = os.Create(b.input)
			if err != nil {
				return err
			}
		}

		if err := gen.Finish(outputFile, &buffer); err != nil {
			return err
		}

		if err := outputFile.Close(); err != nil {
			return err
		}
	}
	return nil
}

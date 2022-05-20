package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"karmem.org/cmd/karmem/kmgen"
	"karmem.org/cmd/karmem/kmparser"
	karmem "karmem.org/golang"
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

type Build struct {
	language []bool
	output   string
	input    string
}

func (b *Build) Parse() {
	flags := flag.NewFlagSet("build", flag.ExitOnError)
	b.language = make([]bool, len(kmgen.Generators))
	for i, g := range kmgen.Generators {
		flags.BoolVar(&b.language[i], g.Language(), false, fmt.Sprintf("Enable geneartion for %s language.", g.Language()))
	}
	flags.StringVar(&b.output, "o", ".", "Output directory path.")
	flags.Parse(flag.Args()[1:])

	b.input = flags.Arg(0)
}

func (b *Build) Generator() (v kmgen.Generator, err error) {
	for i, g := range kmgen.Generators {
		if b.language[i] {
			if v != nil {
				return nil, errors.New("multiple languages is not supported")
			}
			v = g
		}
	}
	if v == nil {
		return nil, errors.New("missing language. Please, specify one output language (such as --golang)")
	}
	return v, nil
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

	writer := karmem.NewWriter(8_000_000)
	if _, err := parsed.WriteAsRoot(writer); err != nil {
		return err
	}

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

	return nil
}

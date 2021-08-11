package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasttemplate"
)

var debug bool

type RequiredField struct {
	InputFilename  string `envconfig:"INPUT" required:"true"`
	OutputFilename string `envconfig:"OUTPUT" required:"true"`
	Debug          bool   `envconfig:"DEBUG" default:"false"`
	StartTag       string `envconfig:"START_TAG"`
	EndTag         string `envconfig:"END_TAG"`
}

func getInputEnvironment(name string, defaultValue string) string {
	if v := os.Getenv("INPUT_" + strings.ToUpper(name)); v != "" {
		return v
	}
	return defaultValue
}

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

var (
	ErrEnvRequired = errors.New("required environment variales is INPUT or OUTPUT")
)

type ErrNoFileExists struct {
	err error
}

var _ error = ErrNoFileExists{}

func (e ErrNoFileExists) Error() string {
	return e.err.Error()
}

func (e ErrNoFileExists) Unwrap() error {
	return e.err
}

func debugPrint(format string, v ...interface{}) {
	if debug {
		log.Printf(format, v...)
	}
}

func Process() error {
	flag.Parse()
	if inputVar != "" {
		os.Setenv("INPUT_INPUT", inputVar)
	}
	if outputVar != "" {
		os.Setenv("INPUT_OUTPUT", outputVar)
	}
	requiredField := RequiredField{}
	if err := envconfig.Process("INPUT", &requiredField); err != nil {
		return ErrEnvRequired
	}
	if requiredField.StartTag == "" {
		requiredField.StartTag = "${{"
	}
	if requiredField.EndTag == "" {
		requiredField.EndTag = "}}"
	}

	debug = requiredField.Debug
	if !isExists(requiredField.InputFilename) {
		return ErrNoFileExists{err: fmt.Errorf("No such file or directory: %s", requiredField.InputFilename)}
	}
	inputFileRead, err := os.ReadFile(requiredField.InputFilename)
	if err != nil {
		return err
	}
	bufferWrite := &bytes.Buffer{}
	inputBuffer := bytes.NewBuffer(inputFileRead)
	tpl := fasttemplate.New(inputBuffer.String(), requiredField.StartTag, requiredField.EndTag)
	_, err = tpl.ExecuteFunc(bufferWrite, func(w io.Writer, tag string) (int, error) {
		debugPrint(tag)
		v := getInputEnvironment(tag, os.Getenv(tag))
		return w.Write([]byte(v))
	})
	if err != nil {
		return err
	}
	if debug {
		debugPrint(bufferWrite.String())
	}
	outputFileWriter, err := os.Create(requiredField.OutputFilename)
	if err != nil {
		return err
	}
	defer outputFileWriter.Close()
	_, err = io.Copy(outputFileWriter, bufferWrite)
	return err
}

func main() {
	if err := Process(); err != nil {
		log.Fatalln(err)
	}
}

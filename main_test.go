package main

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedFile = "tests/test.txt"
)

func TestRequiredInputErr(t *testing.T) {
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "test")
	err := Process()
	assert.Error(t, err, ErrEnvRequired.Error())
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
}

func TestRequiredOutputErr(t *testing.T) {
	os.Setenv("INPUT_INPUT", "test")
	os.Setenv("INPUT_OUTPUT", "")
	err := Process()
	assert.Error(t, err, ErrEnvRequired.Error())
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
}

func TestInputFileNotFound(t *testing.T) {
	const (
		tplFile = "tests/notfound"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", "test")
	err := Process()
	m := &ErrNoFileExists{}
	assert.ErrorAs(t, err, m)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
}

func TestHelloWorld1(t *testing.T) {
	const (
		tplFile    = "tests/text-tpl1.txt"
		actualFile = "tests/TestHelloWorld1.txt"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("INPUT_NAME", "world")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile(expectedFile)
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("INPUT_NAME", "")
}

func TestHelloWorld2(t *testing.T) {
	const (
		tplFile    = "tests/text-tpl1.txt"
		actualFile = "tests/TestHelloWorld2.txt"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("NAME", "world")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile(expectedFile)
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("NAME", "")
}

func TestHelloWorldOriginalTag(t *testing.T) {
	const (
		tplFile    = "tests/text-tpl2.txt"
		actualFile = "tests/TestHelloWorldOriginalTag.txt"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("INPUT_START_TAG", "[")
	os.Setenv("INPUT_END_TAG", "]")
	os.Setenv("INPUT_NAME", "world")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile(expectedFile)
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("INPUT_START_TAG", "")
	os.Setenv("INPUT_END_TAG", "")
	os.Setenv("INPUT_NAME", "")
}

func TestHelloWorldOriginalTag2(t *testing.T) {
	const (
		tplFile    = "tests/text-tpl2.txt"
		actualFile = "tests/TestHelloWorldOriginalTag2.txt"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("INPUT_START_TAG", "[")
	os.Setenv("INPUT_END_TAG", "]")
	os.Setenv("NAME", "world")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile(expectedFile)
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("INPUT_START_TAG", "")
	os.Setenv("INPUT_END_TAG", "")
	os.Setenv("NAME", "")
}

func TestHelloWorldOriginalTag3(t *testing.T) {
	const (
		tplFile    = "tests/text-tpl2.txt"
		actualFile = "tests/TestHelloWorldOriginalTag3.txt"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("INPUT_START_TAG", "[")
	os.Setenv("INPUT_END_TAG", "]")
	os.Setenv("NAME", "world")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile(expectedFile)
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("INPUT_START_TAG", "")
	os.Setenv("INPUT_END_TAG", "")
	os.Setenv("NAME", "")
}

func TestDeploymentAndConfigMap(t *testing.T) {
	const (
		tplFile    = "tests/deployment-tpl.yml"
		actualFile = "tests/TestDeploymentAndConfigMap.yml"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("version", "alpine")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile("tests/deployment.yml")
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("version", "")
}

func TestDeploymentAndConfigMap2(t *testing.T) {
	const (
		tplFile    = "tests/deployment-tpl.yml"
		actualFile = "tests/TestDeploymentAndConfigMap2.txt"
	)
	os.Setenv("INPUT_INPUT", tplFile)
	os.Setenv("INPUT_OUTPUT", actualFile)
	os.Setenv("VERSION", "alpine")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile("tests/deployment.yml")
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.NotEqual(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("VERSION", "")
}

func TestDeploymentAndConfigMapFlag(t *testing.T) {
	const (
		tplFile    = "tests/deployment-tpl.yml"
		actualFile = "tests/TestDeploymentAndConfigMap2.txt"
	)
	_ = flag.CommandLine.Set("i", tplFile)
	_ = flag.CommandLine.Set("o", actualFile)
	os.Setenv("VERSION", "alpine")
	err := Process()
	if assert.NoError(t, err) {
		expected, err := os.ReadFile("tests/deployment.yml")
		assert.NoError(t, err)
		actual, err := os.ReadFile(actualFile)
		assert.NoError(t, err)
		assert.NotEqual(t, string(expected), string(actual))
	}
	os.Remove(actualFile)
	os.Setenv("INPUT_INPUT", "")
	os.Setenv("INPUT_OUTPUT", "")
	os.Setenv("VERSION", "")
}

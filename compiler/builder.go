package compiler

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	// "github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/ONSBR/bddTest.Go/util"

	"bytes"
)

var logBuilder = util.GetLogger("compiler.builder")

type (
	BuilderFileError struct {
		Filename string
		Error    string
	}
	BuilderError struct {
		Errors []BuilderFileError
	}

	Builder struct{}

	BuilderInterface interface {
		BuildFile(string) (ExecutionTestTree, error)
		BuildFiles(string) ([]ExecutionTestTree, error)
		GenerateYamlPageObject(ExecutionTestTree) (string, error)
		GenerateYamlPageObjects([]ExecutionTestTree) ([]string, error)
	}

	BuilderErrorInterface interface {
		Error() string
	}
)

func (this *Builder) BuildFile(filename string) (tree ExecutionTestTree) {
	codeCompiler := &CodeCompiler{}
	fileHandler := &util.FileHandler{}

	content, errRead := fileHandler.ReadFile(filename)
	logBuilder.Infof("%s", content)
	if errRead != nil {
		tree = ExecutionTestTree{Error: lexer.ParserError{Message: errRead.Error()}, HasError: true}
		return
	}
	tree = codeCompiler.BuildExecutionTestTree(content)
	logBuilder.Infof("%s", tree.Error)
	return
}

func (this *Builder) BuildFiles(folderPattern string) (trees []ExecutionTestTree, err error) {
	var builderError = &BuilderError{}
	fileHandler := &util.FileHandler{}
	filenames, errFind := fileHandler.FindFiles(folderPattern)
	if errFind != nil {
		err = BuilderError{Errors: []BuilderFileError{BuilderFileError{Error: errFind.Error()}}}
		return
	}
	for _, filename := range filenames {
		tree := this.BuildFile(filename)
		if tree.HasError {
			message := concatenateErrorMessage(filename, tree.Error)
			builderFileError := BuilderFileError{Filename: filename, Error: message}
			builderError.Errors = append(builderError.Errors, builderFileError)
			err = builderError
		}
		trees = append(trees, tree)
	}
	return
}

func (this *Builder) GenerateYamlPageObject(tree ExecutionTestTree) (yaml string, err error) {
	pageObjectDefParser := &PageObjectDefParser{}
	yaml, err = pageObjectDefParser.GetDefinitionFromTree(tree)
	return
}

func (this *Builder) GenerateYamlPageObjects(trees []ExecutionTestTree) (yamls []string, err error) {
	for _, tree := range trees {
		yaml, errYaml := this.GenerateYamlPageObject(tree)
		if errYaml != nil {
			err = errYaml
			return
		}
		yamls = append(yamls, yaml)
	}
	return
}

func concatenateErrorMessage(filename string, err lexer.ParserError) string {
	var buffer bytes.Buffer
	buffer.WriteString(filename)
	buffer.WriteString("> ")
	buffer.WriteString(err.Message)
	buffer.WriteString(": ")
	buffer.WriteString(err.Token)
	return buffer.String()
}

func (this BuilderError) Error() string {
	var buffer bytes.Buffer

	for _, builderFileError := range this.Errors {
		buffer.WriteString(builderFileError.Filename)
		buffer.WriteString(": ")
		buffer.WriteString(builderFileError.Error)
		buffer.WriteString("\n")
	}
	return buffer.String()
}

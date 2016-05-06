package compiler

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/ONSBR/bddTest.Go/util"

	"bytes"
	"fmt"
	"os"
	"strings"
)

const sep = string(os.PathSeparator)

var logBuilder = util.GetLogger("compiler.builder")

type (
	BuilderFileError struct {
		Filename string
		Error    string
	}
	BuilderError struct {
		Errors []BuilderFileError
	}

	ExecutionTestTreeResult struct {
		Filename      string
		ExecutionTree ExecutionTestTree
	}

	PageObjectResult struct {
		Filename   string
		PageObject pageObject.PageObject
	}

	Execution struct {
		Filename      string
		PageObject    pageObject.PageObject
		ExecutionTree ExecutionTestTree
		HasError      bool
		Error         string
	}

	Builder struct{}

	BuilderInterface interface {
		BuildFile(string) (ExecutionTestTree, error)
		BuildFiles(string) ([]ExecutionTestTreeResult, error)
		GenerateYamlPageObject(ExecutionTestTree) (string, error)
		GenerateYamlPageObjects([]ExecutionTestTree) ([]string, error)
		GeneratePageObject(string, string) (pageObject.PageObject, error)
		BuildExecution(string) Execution
		BuildExecutions(string) ([]Execution, error)
		BuildYamlPageObjectFile(string) error
		BuildYamlPageObjectFiles(string) error
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

func (this *Builder) BuildFiles(folderPattern string) (trees []ExecutionTestTreeResult, err error) {
	var builderError = &BuilderError{}
	var path = folderPattern
	fileHandler := &util.FileHandler{}

	if !strings.Contains(path, sep+"*.specs") {
		if path[len(path)-1:] != sep {
			path += sep
		}
		path += "*.spec"
	}

	filenames, errFind := fileHandler.FindFiles(path)
	if errFind != nil {
		err = BuilderError{Errors: []BuilderFileError{BuilderFileError{Error: errFind.Error()}}}
		return
	}
	for _, filename := range filenames {
		logBuilder.Infof("PATH %s", filename)
		tree := this.BuildFile(filename)
		if tree.HasError {
			message := concatenateErrorMessage(filename, tree.Error)
			builderFileError := BuilderFileError{Filename: filename, Error: message}
			builderError.Errors = append(builderError.Errors, builderFileError)
			err = builderError
		}
		result := ExecutionTestTreeResult{Filename: filename, ExecutionTree: tree}
		trees = append(trees, result)
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

func (this *Builder) GeneratePageObject(filename string, baseUri string) (page pageObject.PageObject, err error) {
	pageObjectDefParser := &PageObjectDefParser{}
	fileHandler := &util.FileHandler{}

	content, errRead := fileHandler.ReadFile(filename)
	logBuilder.Infof("%s", content)
	if errRead != nil {
		err = BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errRead.Error()}}}
		return
	}
	pageObj, errParser := pageObjectDefParser.GetPageObject(content, baseUri)
	if errParser != nil {
		err = BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errParser.Error()}}}
		return
	}
	page = *pageObj
	return
}

func (this *Builder) BuildExecution(filename string, baseUri string) Execution {
	fileHandler := &util.FileHandler{}
	var exec Execution

	executionTestTreeResult := this.BuildFile(filename)

	if executionTestTreeResult.HasError {
		exec = Execution{Filename: filename, HasError: true, Error: concatenateErrorMessage(filename, executionTestTreeResult.Error)}
	} else {
		pageFilename := fmt.Sprintf("%s.page", filename)
		if !fileHandler.DoesFileExists(pageFilename) {
			exec = Execution{Filename: filename, HasError: true, Error: fmt.Sprintf("Page Object file missing: %s", pageFilename)}
		} else {
			if page, errPage := this.GeneratePageObject(pageFilename, baseUri); errPage == nil {
				exec = Execution{Filename: filename, ExecutionTree: executionTestTreeResult, PageObject: page}
			} else {
				exec = Execution{Filename: filename, HasError: true, Error: errPage.Error()}
			}
		}

	}

	return exec
}

func (this *Builder) BuildExecutions(folderPattern string, baseUri string) (exec []Execution, err error) {
	fileHandler := &util.FileHandler{}

	executionTestTreeResults, errBuild := this.BuildFiles(folderPattern)

	logBuilder.Errorf("COUNT %d", len(executionTestTreeResults))

	if errBuild != nil {
		err = errBuild
		return
	}
	for _, executionTestTreeResult := range executionTestTreeResults {
		pageFilename := fmt.Sprintf("%s.page", executionTestTreeResult.Filename)
		if !fileHandler.DoesFileExists(pageFilename) {
			exec = append(exec, Execution{Filename: executionTestTreeResult.Filename, HasError: true, Error: fmt.Sprintf("Page Object file missing: %s", pageFilename)})
		} else {
			if page, errPage := this.GeneratePageObject(pageFilename, baseUri); errPage == nil {
				exec = append(exec, Execution{Filename: executionTestTreeResult.Filename, ExecutionTree: executionTestTreeResult.ExecutionTree, PageObject: page})
			} else {
				exec = append(exec, Execution{Filename: executionTestTreeResult.Filename, HasError: true, Error: errPage.Error()})
			}
		}
	}
	return
}

func (this *Builder) BuildYamlPageObjectFile(filename string) error {
	fileHandler := &util.FileHandler{}

	executionTestTree := this.BuildFile(filename)
	if executionTestTree.HasError {
		errorMessage := concatenateErrorMessage(filename, executionTestTree.Error)
		return BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errorMessage}}}
	}

	yaml, err := this.GenerateYamlPageObject(executionTestTree)
	if err != nil {
		errorMessage := err.Error()
		return BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errorMessage}}}
	}
	if fileHandler.DoesFileExists(filename + ".page") {
		errBkp := fileHandler.BackupFile(filename + ".page")
		if errBkp != nil {
			return BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename + ".page", Error: errBkp.Error()}}}
		}
	}
	errWrite := fileHandler.WriteFile(filename+".page", yaml)
	if errWrite != nil {
		return BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename + ".page", Error: errWrite.Error()}}}
	}
	return nil
}

func (this *Builder) BuildYamlPageObjectFiles(folderPattern string) error {
	builderError := &BuilderError{}
	fileHandler := &util.FileHandler{}

	executionTestTreeResults, errBuild := this.BuildFiles(folderPattern)

	logBuilder.Errorf("COUNT %d", len(executionTestTreeResults))

	if errBuild != nil {
		return errBuild
	}
	for _, executionTestTreeResult := range executionTestTreeResults {
		executionTestTree := executionTestTreeResult.ExecutionTree
		filename := executionTestTreeResult.Filename

		if executionTestTree.HasError {
			errorMessage := concatenateErrorMessage(filename, executionTestTree.Error)
			builderError.Errors = append(builderError.Errors, BuilderFileError{Filename: filename, Error: errorMessage})
		}

		yaml, err := this.GenerateYamlPageObject(executionTestTree)
		if err != nil {
			errorMessage := err.Error()
			builderError.Errors = append(builderError.Errors, BuilderFileError{Filename: filename, Error: errorMessage})
		}
		if fileHandler.DoesFileExists(filename + ".page") {
			errBkp := fileHandler.BackupFile(filename + ".page")
			if errBkp != nil {
				builderError.Errors = append(builderError.Errors, BuilderFileError{Filename: filename + ".page", Error: errBkp.Error()})
			}
		}
		errWrite := fileHandler.WriteFile(filename+".page", yaml)
		if errWrite != nil {
			builderError.Errors = append(builderError.Errors, BuilderFileError{Filename: filename + ".page", Error: errWrite.Error()})
		}
	}
	return builderError
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

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
	
	//BuilderFileError support type for error handler
	BuilderFileError struct {
		Filename string
		Error    string
	}
	
	//BuilderError support type for error handler
	BuilderError struct {
		Errors []BuilderFileError
	}

	//ExecutionTestTreeResult is the parsed execution test tree based on .spec file
	ExecutionTestTreeResult struct {
		Filename      string
		ExecutionTree ExecutionTestTree
	}

	//PageObjectResult is the parsed page object instance based on .spec.page file
	PageObjectResult struct {
		Filename   string
		PageObject pageObject.PageObject
	}

	//Execution defines a single execution element based on a spec file with execution tree and page object
	Execution struct {
		Filename      string
		PageObject    pageObject.PageObject
		ExecutionTree ExecutionTestTree
		HasError      bool
		Error         string
	}

	//Builder is the base structure of bddTest.Go builder
	Builder struct{}

	iBuilder interface {
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

	iBuilderError interface {
		Error() string
	}
)

//BuildFile get a single spec file and try to validate 
func (builder *Builder) BuildFile(filename string) (tree ExecutionTestTree) {
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

//BuildFiles find all spec files based on folder pattern and try to validate 
func (builder *Builder) BuildFiles(folderPattern string) (trees []ExecutionTestTreeResult, err error) {
	// var builderError = &BuilderError{}
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
		tree := builder.BuildFile(filename)
		// if tree.HasError {
		// 	message := concatenateErrorMessage(filename, tree.Error)
		// 	builderFileError := BuilderFileError{Filename: filename, Error: message}
		// 	builderError.Errors = append(builderError.Errors, builderFileError)
		// 	err = builderError
		// }
		result := ExecutionTestTreeResult{Filename: filename, ExecutionTree: tree}
		trees = append(trees, result)
	}
	return
}

//GenerateYamlPageObject generate yaml stub definition of page object based on parsed execution test tree
func (builder *Builder) GenerateYamlPageObject(tree ExecutionTestTree) (yaml string, err error) {
	pageObjectDefParser := &PageObjectDefParser{}
	yaml, err = pageObjectDefParser.GetDefinitionFromTree(tree)
	return
}

//GenerateYamlPageObjects generate yaml stub definitions of page objects based on parsed execution test trees
func (builder *Builder) GenerateYamlPageObjects(trees []ExecutionTestTree) (yamls []string, err error) {
	for _, tree := range trees {
		yaml, errYaml := builder.GenerateYamlPageObject(tree)
		if errYaml != nil {
			err = errYaml
			return
		}
		yamls = append(yamls, yaml)
	}
	return
}

//GeneratePageObject instantiate a page object based on page Object Definition file
func (builder *Builder) GeneratePageObject(filename string, baseURI string) (page pageObject.PageObject, err error) {
	pageObjectDefParser := &PageObjectDefParser{}
	fileHandler := &util.FileHandler{}

	content, errRead := fileHandler.ReadFile(filename)
	logBuilder.Infof("%s", content)
	if errRead != nil {
		err = BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errRead.Error()}}}
		return
	}
	pageObj, errParser := pageObjectDefParser.GetPageObject(content, baseURI)
	if errParser != nil {
		err = BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errParser.Error()}}}
		return
	}
	page = *pageObj
	return
}

//BuildExecution get single spec and page object definition and create execution tree of test
func (builder *Builder) BuildExecution(filename string, baseURI string) Execution {
	fileHandler := &util.FileHandler{}
	var exec Execution

	executionTestTreeResult := builder.BuildFile(filename)

	if executionTestTreeResult.HasError {
		exec = Execution{Filename: filename, HasError: true, Error: concatenateErrorMessage(filename, executionTestTreeResult.Error)}
	} else {
		pageFilename := fmt.Sprintf("%s.page", filename)
		if !fileHandler.DoesFileExists(pageFilename) {
			exec = Execution{Filename: filename, HasError: true, Error: fmt.Sprintf("Page Object file missing: %s", pageFilename)}
		} else {
			if page, errPage := builder.GeneratePageObject(pageFilename, baseURI); errPage == nil {
				exec = Execution{Filename: filename, ExecutionTree: executionTestTreeResult, PageObject: page}
			} else {
				exec = Execution{Filename: filename, HasError: true, Error: errPage.Error()}
			}
		}

	}

	return exec
}

//BuildExecutions find specs and page object definitions on folder pattern and create execution tree of tests
func (builder *Builder) BuildExecutions(folderPattern string, baseURI string) (exec []Execution, err error) {
	fileHandler := &util.FileHandler{}

	executionTestTreeResults, errBuild := builder.BuildFiles(folderPattern)

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
			if page, errPage := builder.GeneratePageObject(pageFilename, baseURI); errPage == nil {
				exec = append(exec, Execution{Filename: executionTestTreeResult.Filename, ExecutionTree: executionTestTreeResult.ExecutionTree, PageObject: page})
			} else {
				exec = append(exec, Execution{Filename: executionTestTreeResult.Filename, HasError: true, Error: errPage.Error()})
			}
		}
	}
	return
}

//BuildYamlPageObjectFile get a single spec file and create stub of Yaml Page Object Definition file
func (builder *Builder) BuildYamlPageObjectFile(filename string) error {
	fileHandler := &util.FileHandler{}

	executionTestTree := builder.BuildFile(filename)
	if executionTestTree.HasError {
		errorMessage := concatenateErrorMessage(filename, executionTestTree.Error)
		return BuilderError{Errors: []BuilderFileError{BuilderFileError{Filename: filename, Error: errorMessage}}}
	}

	yaml, err := builder.GenerateYamlPageObject(executionTestTree)
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

//BuildYamlPageObjectFiles find spec files on folder pattern and create stub of Yaml Page Object Definition files
func (builder *Builder) BuildYamlPageObjectFiles(folderPattern string) error {
	builderError := &BuilderError{}
	fileHandler := &util.FileHandler{}

	executionTestTreeResults, errBuild := builder.BuildFiles(folderPattern)

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

		yaml, err := builder.GenerateYamlPageObject(executionTestTree)
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

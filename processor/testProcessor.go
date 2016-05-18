package processor

import (
	"fmt"
	"strings"
	"time"

	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/fatih/color"
)

/*
TestProcessor processes full tests.
*/
type TestProcessor struct {
	Execution *compiler.Execution
}

/*
NewTestProcessor creates a new instance of TestProcessor
*/
func NewTestProcessor(execution *compiler.Execution) *TestProcessor {
	return &TestProcessor{
		Execution: execution,
	}
}

/*
Process processes an Execution Tree.
*/
func (t *TestProcessor) Process() error {
	if err := t.Execution.PageObject.Open(); err != nil {
		color.Red("===> FAIL: %s", t.Execution.Feature.FullText)
		color.Red("===> FAIL: Could not open page %s (%s) due to %s.", t.Execution.PageObject.Page, t.Execution.PageObject.Uri, err)
		return err
	}

	return t.ProcessFeature(&t.Execution.Feature)
}

/*

 */
func printInfo(header, body string) {
	yellow := color.New(color.FgCyan).SprintfFunc()
	white := color.New(color.FgWhite).SprintfFunc()
	fmt.Printf("%s:\t%s\n", yellow("===> %s", header), white(body))
}

/*
ProcessFeature processes a feature and all it's scenarios.
*/
func (t *TestProcessor) ProcessFeature(a *lexer.Feature) error {
	printInfo("FEATURE", a.Name)
	printInfo("SPEC FILE", t.Execution.Filename)
	printInfo("PAGE", a.PageName)
	printInfo("URI", t.Execution.PageObject.Uri)

	for _, scenario := range a.Scenarios {
		if err := t.ProcessScenario(&scenario); err != nil {
			filename := fmt.Sprintf("%s.png", strings.Replace(scenario.Name, " ", "", -1))
			t.Execution.PageObject.SaveScreenShot(filename)
			color.Yellow("===> INFO: Screenshot available at: %s", filename)
			return err
		}
	}

	return nil
}

/*
ProcessScenario processes a scenario and all it's action tests.
*/
func (t *TestProcessor) ProcessScenario(a *lexer.Scenario) error {
	printInfo("SCENARIO", a.Name)

	for _, action := range a.Actions {
		if err := t.ProcessAction(&action); err != nil {
			return err
		}
	}

	return nil
}

/*
ProcessAction processes an action test.
*/
func (t *TestProcessor) ProcessAction(a *lexer.Expect_action) error {
	var err error
	defer printTestResult(time.Now(), a.FullText, &err)

	p, err := t.Execution.PageObject.FindPageElement(a.ObjectId)

	if err != nil {
		return err
	}

	err = pageObject.Execute(p, a)
	return err
}

/*
printTestResult prints test result.
*/
func printTestResult(start time.Time, label string, err *error) {
	elapsed := time.Since(start)

	if *err != nil {
		color.Red("===> FAIL: %s (%s)", label, elapsed)
		color.Red("===> FAIL: %s", *err)
	} else {
		color.Green("===> PASS: %s (%s)", label, elapsed)
	}
}

package processor

import (
	"fmt"
	"strings"
	"time"

	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject/actions"
	"github.com/ONSBR/bddTest.Go/pageObject/assertions"
	"github.com/fatih/color"
)

// TestProcessor processes full tests.
type TestProcessor struct {
	Execution *compiler.Execution
}

// NewTestProcessor creates a new instance of TestProcessor
func NewTestProcessor(execution *compiler.Execution) *TestProcessor {
	return &TestProcessor{
		Execution: execution,
	}
}

// Process processes an Execution Tree.
func (t *TestProcessor) Process() error {
	if err := t.Execution.PageObject.Open(); err != nil {
		color.Red("===> FAIL: %s", t.Execution.Feature.FullText)
		color.Red("===> FAIL: Could not open page %s (%s) due to %s.", t.Execution.PageObject.Page, t.Execution.PageObject.URI, err)
		return err
	}

	return t.ProcessFeature(&t.Execution.Feature)
}

// ProcessFeature processes a feature and all it's scenarios.
func (t *TestProcessor) ProcessFeature(a *lexer.Feature) error {
	printInfo("FEATURE", a.Name)
	printInfo("SPEC FILE", t.Execution.Filename)
	printInfo("PAGE", a.PageName)
	printInfo("URI", t.Execution.PageObject.URI)

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

// ProcessScenario processes a scenario and all it's action tests.
func (t *TestProcessor) ProcessScenario(a *lexer.Scenario) error {
	fmt.Println("")
	printInfo("SCENARIO", a.Name)

	for _, action := range a.Actions {
		if err := t.ProcessAction(&action); err != nil {
			return err
		}
	}

	for _, expectation := range a.Expectations {
		if err := t.ProcessExpectation(&expectation); err != nil {
			return err
		}
	}

	return nil
}

// ProcessAction processes an action test.
func (t *TestProcessor) ProcessAction(action *lexer.Expect_action) error {
	var err error
	defer printTestResult(time.Now(), action.FullText, &err)

	element, err := t.Execution.PageObject.FindPageElement(action.ObjectId)

	if err != nil {
		return err
	}

	err = actions.Execute(element, action)
	return err
}

// ProcessExpectation processes an action test.
func (t *TestProcessor) ProcessExpectation(expectation *lexer.Expect_expression) error {
	var err error
	defer printTestResult(time.Now(), expectation.FullText, &err)

	element, err := t.Execution.PageObject.FindPageElement(expectation.ObjectId)

	if err != nil {
		return err
	}

	err = assertions.Assert(element, expectation)
	return err
}

func printTestResult(start time.Time, label string, err *error) {
	elapsed := time.Since(start).Seconds() * 1000

	if *err != nil {
		color.Red("===> FAIL: %s (%.2f ms)", label, elapsed)
		color.Red("===> FAIL: %s", *err)
	} else {
		color.Green("===> PASS: %s (%.2f ms)", label, elapsed)
	}
}

func printInfo(header, body string) {
	yellow := color.New(color.FgCyan).SprintfFunc()
	white := color.New(color.FgWhite).SprintfFunc()
	fmt.Printf("%s:\t%s\n", yellow("===> %s", header), white(body))
}

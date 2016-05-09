// Package cli provides functions to parsing and validade bddTest.Go command line
package cli

import (
    "github.com/voxelbrain/goptions"
    "github.com/ONSBR/bddTest.Go/util"
    "bytes"
    "os"
    "io"
)

var (
    logFlag = util.GetLogger("cli.flagparser")
)

const (
    ErrCommand = -1
    ValidadeCommand = iota
    YamlCommand
    RunCommand
)

type (
    OptionsDef struct {
        SpecFile        string `goptions:"-f, --specfile, description='Single spec file to be used'"`
        Multi           string `goptions:"-m, --multi, description='Spec file path pattern'"`
        BaseURI         string `goptions:"-u, --uri, description='Base URI concatenated to page URI on test execution'"`
        SeleniumServer  string `goptions:"-s, --seleniumserver, description='URI of selenium server'"`
        Help            goptions.Help `goptions:"-h, --help, description='Show this help'"`

        Verb goptions.Verbs
        Validate struct {
        } `goptions:"validate"`
        Yaml struct {
            Backup bool   `goptions:"-b, --backup, description='If set and page object file exists, it will be renamed to .bkp sufix'"`
        } `goptions:"yaml"`
        Run struct {
        } `goptions:"run"`
    }
    
    FlagParser struct {
        Options OptionsDef
    }
    
    iFlagParser interface {
        Usage() string
        Parse() int
    }
)

//Usage prepare help string for bddTest.Go command line
func (flagParser *FlagParser) Usage() string  {
    old := os.Stderr // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stderr = w

    goptions.PrintHelp()

    errC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        errC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stderr = old // restoring the real stdout
    out := <-errC

    return out
}

//Parse try to validade all flags and verbs from command line and return parsed code
func (flagParser *FlagParser) Parse() int {
    retCode := ErrCommand
    errParse := goptions.Parse(&flagParser.Options)
    if errParse != nil {
        retCode = ErrCommand
    } else {
        switch flagParser.Options.Verb {
        case "validate":
            // flagParser.Validade.Parse(os.Args[2:])
            retCode = ValidadeCommand
            break
        case "yaml":
            // flagParser.Yaml.Parse(os.Args[2:])
            retCode = YamlCommand
            break
        case "run":   
            // flagParser.Run.Parse(os.Args[2:])
            retCode = RunCommand
            break
        }
        if retCode != ErrCommand {
            
        }
    }
  
    return retCode
}

//NewFlagParser Constructor and flag register
func NewFlagParser()  *FlagParser{
    flagParser := &FlagParser{}
    flagParser.Options = OptionsDef{SeleniumServer:"http://127.0.0.1:4444/wd/hub", BaseURI:"http://127.0.0.1"}
    logFlag.Infof("%s",os.Args)
    _ = goptions.Parse(&flagParser.Options)
    return flagParser
}
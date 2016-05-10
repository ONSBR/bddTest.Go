bddTest.Go
=

BDD Test Framework with DSL for Human Test Specifications

This is a framework for End2End (Acceptance) tests using [Selenium Webdriver](http://www.seleniumhq.org/projects/webdriver/) to run tests on different browsers.

**bddTest.Go** also defines a domain language, based on [Gherkin](https://github.com/cucumber/cucumber/wiki/Gherkin) language and [BDD](https://en.wikipedia.org/wiki/Behavior-driven_development) style.

Features
-------------

- DSL defined for Behavior Driven Tests
- Run tests with Remote Selenium Webdriver Server
- Library of page elements abstracting DOM
- Auto generation of [page objects](http://martinfowler.com/bliki/PageObject.html) based on a simple YAML definition file

Command Line
---------------------

```
Usage: bddTest.Go [global options] <verb> [verb options]
Global options:
    -f, --specfile       Single spec file to be used
    -m, --multi          Spec file path pattern
    -u, --uri            Base URI concatenated to page URI on test execution (default: http://127.0.0.1)
    -s, --seleniumserver URI of selenium server (default: http://127.0.0.1:4444/wd/hub)
    -c, --configfile     Configuration file
    -h, --help           Show this help

Verbs:
    run:
    validate:
    yaml:
        -b, --backup         If set and page object file exists, it will be renamed to .bkp sufix
```

Build and Install
-----------------------
For build and install bddTest.Go, choose precompiled from:

- Nuget
- Pip
- Apt

Or you can build from source:

 1. Install Golang <https://golang.org/doc/install>
 2. Install Godep `go get github.com/tools/godep`
 3. Clone repository 
 4. Get dependencies `godep restore`
 5. Build and install bddTest.Go package `go build;go install`
 
Usage
--------
bddTest.Go is based on a DSL for write tests and a YAML file format to define Page Objects.
First, write some tests using [bddTest.Go DSL](https://github.com/ONSBR/bddTest.Go/blob/master/bddTest.Go.DSL.md) and then use bddTest.Go command line to generate Page Objects definitions:
 
1. Create a folder for your spec files
2. Write a .spec file for each page you want to test. This .spec file can have several scenarios
3. Run `bddTest.Go -m [your folder pattern: ./tests*/**/] validate` to build spec files
4. Run `bddTest.Go -m [your folder pattern: ./tests*/**/] yaml` to generate .spec.page files
5. Complete each .spec.page generated file with page relative URI and element locators
6. Run `bddTest.Go -m [your folder pattern: ./tests*/**/] -u [your base URI: http://localhost/] -s [your selenium server: http://localhost:4444/wd/hub] -c [your log configuration file] run`


Contribution
------------------


License
----------

[GNU GPL V3](https://raw.githubusercontent.com/ONSBR/bddTest.Go/master/LICENSE)
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
> Usage: bddTest.Go [global options] <verb> [verb options]</b>
> Global options:</br>
>&nbsp;&nbsp;&nbsp;  -f, --specfile       Single spec file to be used</b>
>&nbsp;&nbsp;&nbsp; -m, --multi          Spec file path pattern</b>
>&nbsp;&nbsp;&nbsp;-u, --uri            Base URI concatenated to page URI on test execution (default: http://127.0.0.1)</b>
>&nbsp;&nbsp;&nbsp;-s, --seleniumserver URI of selenium server (default: http://127.0.0.1:4444/wd/hub)</b>
> &nbsp;&nbsp;&nbsp;-c, --configfile     Configuration file</b>
>&nbsp;&nbsp;&nbsp; -h, --help           Show this help</b>
></b>
>Verbs:</b>
>&nbsp;&nbsp;&nbsp;run:</b>
> &nbsp;&nbsp;&nbsp;validate:</b>
>&nbsp;&nbsp;&nbsp;yaml:</b>
>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;-b, --backup         If set and page object file exists, it will be renamed to .bkp sufix</b>

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


Contribution
------------------


License
----------

[GNU GPL V3](https://raw.githubusercontent.com/ONSBR/bddTest.Go/master/LICENSE)
package main

import (
    "fmt"
    do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
//    do.Env = `GOPATH=.vendor::$GOPATH`

    p.Task("default", do.S{"hello"}, nil)

    p.Task("hello", nil, func(c *do.Context) {
        name := c.Args.AsString("name", "n")
        if name == "" {
            c.Bash("echo Hello $USER!")
        } else {
            fmt.Println("Hello", name)
        }
    })
    
    p.Task("SimpleTest",nil, func(c *do.Context) {
		c.Run("go run test/compiler/SimpleTest.go")
    })
//
//    p.Task("assets?", nil,  func(c *do.Context) {
//        // The "?" tells Godo to run this task ONLY ONCE regardless of
//        // how many tasks depend on it. In this case watchify watches
//        // on its own.
//        c.Run("watchify public/js/index.js d -o dist/js/app.bundle.js")
//    }).Src("public/**/*.{css,js,html}")
//
//    p.Task("build", nil, func(c *do.Context) {
//        c.Run("GOOS=linux GOARCH=amd64 go build", do.M{"$in": "compiler"})
//    }).Src("**/*.go")
//
//    p.Task("server", do.S{"views", "assets"}, func(c *do.Context) {
//        // rebuilds and restarts when a watched file changes
//        c.Start("main.go", do.M{"$in": "cmd/server"})
//    }).Src("server/**/*.go", "cmd/server/*.{go,json}").
//       Debounce(3000)
//
//    p.Task("views", nil, func(c *do.Context) {
//        c.Run("razor templates")
//    }).Src("templates/**/*.go.html")
}

func main() {
    do.Godo(tasks)
}
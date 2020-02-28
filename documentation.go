// (godoc-to-GitHub-MarkDown) generates package documentation in
// GitHub flavoured Markdown.
//
// Usage:
//
//    godoc2ghmd [options] [full package name]
//
// For full and up-to-date Usage information run:
//
//    $ godoc2ghmd -help
//
// Notes
//
// While the output of godoc2ghmd can simply be piped into a file, e.g.:
//
//    $ godoc2ghmd github.com/gglachant/godoc2ghmd > README.md
//
// it can also be used with the -file option:
//
//    $ godoc2ghmd -file=README.md github.com/gglachant/godoc2ghmd
//
// This invocation is particularly useful when using `go generate` to automate
// the creation of package documentation before sumbmitting code. For
// example, this directive:
//
//go:generate godoc2ghmd -file=README.md github.com/gglachant/godoc2ghmd
//
// in the  `documentation.go` file within this repositoy created this very
// `README.md` file by running:
//
//    $ go generate
//
// The same command also (re)generates all the documentation in the examples
// folder via other similar directives.
//
// History
//
// This is a fork of https://github.com/davecheney/godoc2md with (most of) the
// changes from https://github.com/wdamron/godoc2gh rolled in. I have then
// made some additional fixes and modifications. (I have changed the repo name
// only to make it easier to work on these changes while keeping the previous
// versions intact for comparison).
package main

// This comment block (re)generates the documentation in the examples folder.
//go:generate godoc2ghmd -ex -file=examples/doc/README.md go/doc
//go:generate godoc2ghmd -ex -file=examples/build/README.md go/build
//go:generate godoc2ghmd -ex -file=examples/ioutil/README.md io/ioutil
//go:generate godoc2ghmd -ex -file=examples/http/README.md net/http
//

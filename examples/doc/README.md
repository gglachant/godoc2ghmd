# doc
`import "go/doc"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package doc extracts source code documentation from a Go AST.

## <a name="pkg-imports">Imported Packages</a>

- internal/lazyregexp

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func IsPredeclared(s string) bool](#IsPredeclared)
* [func Synopsis(s string) string](#Synopsis)
* [func ToHTML(w io.Writer, text string, words map[string]string)](#ToHTML)
* [func ToText(w io.Writer, text string, indent, preIndent string, width int)](#ToText)
* [type Example](#Example)
  * [func Examples(testFiles ...\*ast.File) []\*Example](#Examples)
* [type Filter](#Filter)
* [type Func](#Func)
* [type Mode](#Mode)
* [type Note](#Note)
* [type Package](#Package)
  * [func New(pkg \*ast.Package, importPath string, mode Mode) \*Package](#New)
  * [func NewFromFiles(fset \*token.FileSet, files []\*ast.File, importPath string, opts ...interface{}) (\*Package, error)](#NewFromFiles)
  * [func (p \*Package) Filter(f Filter)](#Package.Filter)
* [type Type](#Type)
* [type Value](#Value)

#### <a name="pkg-examples">Examples</a>
* [NewFromFiles](#example_NewFromFiles)

#### <a name="pkg-files">Package files</a>
[comment.go](./comment.go) [doc.go](./doc.go) [example.go](./example.go) [exports.go](./exports.go) [filter.go](./filter.go) [reader.go](./reader.go) [synopsis.go](./synopsis.go) 

## <a name="pkg-variables">Variables</a>
``` go
var IllegalPrefixes = []string{
    "copyright",
    "all rights",
    "author",
}
```

## <a name="IsPredeclared">func</a> [IsPredeclared](./reader.go#L867)
``` go
func IsPredeclared(s string) bool
```
IsPredeclared reports whether s is a predeclared identifier.

## <a name="Synopsis">func</a> [Synopsis](./synopsis.go#L68)
``` go
func Synopsis(s string) string
```
Synopsis returns a cleaned version of the first sentence in s.
That sentence ends after the first period followed by space and
not preceded by exactly one uppercase letter. The result string
has no \n, \r, or \t characters and uses only single spaces between
words. If s starts with any of the IllegalPrefixes, the result
is the empty string.

## <a name="ToHTML">func</a> [ToHTML](./comment.go#L309)
``` go
func ToHTML(w io.Writer, text string, words map[string]string)
```
ToHTML converts comment text to formatted HTML.
The comment was prepared by DocReader,
so it is known not to have leading, trailing blank lines
nor to have trailing spaces at the end of lines.
The comment markers have already been removed.

Each span of unindented non-blank lines is converted into
a single paragraph. There is one exception to the rule: a span that
consists of a single line, is followed by another paragraph span,
begins with a capital letter, and contains no punctuation
other than parentheses and commas is formatted as a heading.

A span of indented lines is converted into a <pre> block,
with the common indent prefix removed.

URLs in the comment text are converted into links; if the URL also appears
in the words map, the link is taken from the map (if the corresponding map
value is the empty string, the URL is not converted into a link).

A pair of (consecutive) backticks (`) is converted to a unicode left quote (“), and a pair of (consecutive)
single quotes (') is converted to a unicode right quote (”).

Go identifiers that appear in the words map are italicized; if the corresponding
map value is not the empty string, it is considered a URL and the word is converted
into a link.

## <a name="ToText">func</a> [ToText](./comment.go#L426)
``` go
func ToText(w io.Writer, text string, indent, preIndent string, width int)
```
ToText prepares comment text for presentation in textual output.
It wraps paragraphs of text to width or fewer Unicode code points
and then prefixes each line with the indent. In preformatted sections
(such as program text), it prefixes each non-blank line with preIndent.

A pair of (consecutive) backticks (`) is converted to a unicode left quote (“), and a pair of (consecutive)
single quotes (') is converted to a unicode right quote (”).

## <a name="Example">type</a> [Example](./example.go#L22-L33)
``` go
type Example struct {
    Name        string // name of the item being exemplified (including optional suffix)
    Suffix      string // example suffix, without leading '_' (only populated by NewFromFiles)
    Doc         string // example function doc string
    Code        ast.Node
    Play        *ast.File // a whole program version of the example
    Comments    []*ast.CommentGroup
    Output      string // expected output
    Unordered   bool
    EmptyOutput bool // expect empty output
    Order       int  // original source code order
}

```
An Example represents an example function found in a test source file.

### <a name="Examples">func</a> [Examples](./example.go#L50)
``` go
func Examples(testFiles ...*ast.File) []*Example
```
Examples returns the examples found in testFiles, sorted by Name field.
The Order fields record the order in which the examples were encountered.
The Suffix field is not populated when Examples is called directly, it is
only populated by NewFromFiles for examples it finds in _test.go files.

Playable Examples must be in a package whose name ends in "_test".
An Example is "playable" (the Play field is non-nil) in either of these
circumstances:

	- The example function is self-contained: the function references only
	  identifiers from other packages (or predeclared identifiers, such as
	  "int") and the test file does not include a dot import.
	- The entire test file is the example: the file contains exactly one
	  example function, zero test or benchmark functions, and at least one
	  top-level function, type, variable, or constant declaration other
	  than the example function.

## <a name="Filter">type</a> [Filter](./filter.go#L9)
``` go
type Filter func(string) bool
```

## <a name="Func">type</a> [Func](./doc.go#L68-L83)
``` go
type Func struct {
    Doc  string
    Name string
    Decl *ast.FuncDecl

    // methods
    // (for functions, these fields have the respective zero value)
    Recv  string // actual   receiver "T" or "*T"
    Orig  string // original receiver "T" or "*T"
    Level int    // embedding level; 0 means not embedded

    // Examples is a sorted list of examples associated with this
    // function or method. Examples are extracted from _test.go files
    // provided to NewFromFiles.
    Examples []*Example
}

```
Func is the documentation for a func declaration.

## <a name="Mode">type</a> [Mode](./doc.go#L96)
``` go
type Mode int
```
Mode values control the operation of New and NewFromFiles.

``` go
const (
    // AllDecls says to extract documentation for all package-level
    // declarations, not just exported ones.
    AllDecls Mode = 1 << iota

    // AllMethods says to show all embedded methods, not just the ones of
    // invisible (unexported) anonymous fields.
    AllMethods

    // PreserveAST says to leave the AST unmodified. Originally, pieces of
    // the AST such as function bodies were nil-ed out to save memory in
    // godoc, but not all programs want that behavior.
    PreserveAST
)
```

## <a name="Note">type</a> [Note](./doc.go#L89-L93)
``` go
type Note struct {
    Pos, End token.Pos // position range of the comment containing the marker
    UID      string    // uid found with the marker
    Body     string    // note body text
}

```
A Note represents a marked comment starting with "MARKER(uid): note body".
Any note with a marker of 2 or more upper case [A-Z] letters and a uid of
at least one character is recognized. The ":" following the uid is optional.
Notes are collected in the Package.Notes map indexed by the notes marker.

## <a name="Package">type</a> [Package](./doc.go#L16-L38)
``` go
type Package struct {
    Doc        string
    Name       string
    ImportPath string
    Imports    []string
    Filenames  []string
    Notes      map[string][]*Note

    // Deprecated: For backward compatibility Bugs is still populated,
    // but all new code should use Notes instead.
    Bugs []string

    // declarations
    Consts []*Value
    Types  []*Type
    Vars   []*Value
    Funcs  []*Func

    // Examples is a sorted list of examples associated with
    // the package. Examples are extracted from _test.go files
    // provided to NewFromFiles.
    Examples []*Example
}

```
Package is the documentation for an entire package.

### <a name="New">func</a> [New](./doc.go#L118)
``` go
func New(pkg *ast.Package, importPath string, mode Mode) *Package
```
New computes the package documentation for the given package AST.
New takes ownership of the AST pkg and may edit or overwrite it.
To have the Examples fields populated, use NewFromFiles and include
the package's _test.go files.

### <a name="NewFromFiles">func</a> [NewFromFiles](./doc.go#L160)
``` go
func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...interface{}) (*Package, error)
```
NewFromFiles computes documentation for a package.

The package is specified by a list of *ast.Files and corresponding
file set, which must not be nil.
NewFromFiles uses all provided files when computing documentation,
so it is the caller's responsibility to provide only the files that
match the desired build context. "go/build".Context.MatchFile can
be used for determining whether a file matches a build context with
the desired GOOS and GOARCH values, and other build constraints.
The import path of the package is specified by importPath.

Examples found in _test.go files are associated with the corresponding
type, function, method, or the package, based on their name.
If the example has a suffix in its name, it is set in the
Example.Suffix field. Examples with malformed names are skipped.

Optionally, a single extra argument of type Mode can be provided to
control low-level aspects of the documentation extraction behavior.

NewFromFiles takes ownership of the AST files and may edit them,
unless the PreserveAST Mode bit is on.

#### Example:

<details>
<summary>Click to expand code.</summary>

```go
package doc_test
	
	import (
	    "bytes"
	    "fmt"
	    "go/ast"
	    "go/doc"
	    "go/format"
	    "go/parser"
	    "go/token"
	    "reflect"
	    "strings"
	    "testing"
	)
	
	const exampleTestFile = `
	package foo_test
	
	import (
	    "flag"
	    "fmt"
	    "log"
	    "sort"
	    "os/exec"
	)
	
	func ExampleHello() {
	    fmt.Println("Hello, world!")
	    // Output: Hello, world!
	}
	
	func ExampleImport() {
	    out, err := exec.Command("date").Output()
	    if err != nil {
	        log.Fatal(err)
	    }
	    fmt.Printf("The date is %s\n", out)
	}
	
	func ExampleKeyValue() {
	    v := struct {
	        a string
	        b int
	    }{
	        a: "A",
	        b: 1,
	    }
	    fmt.Print(v)
	    // Output: a: "A", b: 1
	}
	
	func ExampleKeyValueImport() {
	    f := flag.Flag{
	        Name: "play",
	    }
	    fmt.Print(f)
	    // Output: Name: "play"
	}
	
	var keyValueTopDecl = struct {
	    a string
	    b int
	}{
	    a: "B",
	    b: 2,
	}
	
	func ExampleKeyValueTopDecl() {
	    fmt.Print(keyValueTopDecl)
	    // Output: a: "B", b: 2
	}
	
	// Person represents a person by name and age.
	type Person struct {
	    Name string
	    Age  int
	}
	
	// String returns a string representation of the Person.
	func (p Person) String() string {
	    return fmt.Sprintf("%s: %d", p.Name, p.Age)
	}
	
	// ByAge implements sort.Interface for []Person based on
	// the Age field.
	type ByAge []Person
	
	// Len returns the number of elements in ByAge.
	func (a (ByAge)) Len() int { return len(a) }
	
	// Swap swaps the elements in ByAge.
	func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
	
	// people is the array of Person
	var people = []Person{
	    {"Bob", 31},
	    {"John", 42},
	    {"Michael", 17},
	    {"Jenny", 26},
	}
	
	func ExampleSort() {
	    fmt.Println(people)
	    sort.Sort(ByAge(people))
	    fmt.Println(people)
	    // Output:
	    // [Bob: 31 John: 42 Michael: 17 Jenny: 26]
	    // [Michael: 17 Jenny: 26 Bob: 31 John: 42]
	}
	`
	
	var exampleTestCases = []struct {
	    Name, Play, Output string
	}{
	    {
	        Name:   "Hello",
	        Play:   exampleHelloPlay,
	        Output: "Hello, world!\n",
	    },
	    {
	        Name: "Import",
	        Play: exampleImportPlay,
	    },
	    {
	        Name:   "KeyValue",
	        Play:   exampleKeyValuePlay,
	        Output: "a: \"A\", b: 1\n",
	    },
	    {
	        Name:   "KeyValueImport",
	        Play:   exampleKeyValueImportPlay,
	        Output: "Name: \"play\"\n",
	    },
	    {
	        Name:   "KeyValueTopDecl",
	        Play:   exampleKeyValueTopDeclPlay,
	        Output: "a: \"B\", b: 2\n",
	    },
	    {
	        Name:   "Sort",
	        Play:   exampleSortPlay,
	        Output: "[Bob: 31 John: 42 Michael: 17 Jenny: 26]\n[Michael: 17 Jenny: 26 Bob: 31 John: 42]\n",
	    },
	}
	
	const exampleHelloPlay = `package main
	
	import (
	    "fmt"
	)
	
	func main() {
	    fmt.Println("Hello, world!")
	}
	`
	const exampleImportPlay = `package main
	
	import (
	    "fmt"
	    "log"
	    "os/exec"
	)
	
	func main() {
	    out, err := exec.Command("date").Output()
	    if err != nil {
	        log.Fatal(err)
	    }
	    fmt.Printf("The date is %s\n", out)
	}
	`
	
	const exampleKeyValuePlay = `package main
	
	import (
	    "fmt"
	)
	
	func main() {
	    v := struct {
	        a string
	        b int
	    }{
	        a: "A",
	        b: 1,
	    }
	    fmt.Print(v)
	}
	`
	
	const exampleKeyValueImportPlay = `package main
	
	import (
	    "flag"
	    "fmt"
	)
	
	func main() {
	    f := flag.Flag{
	        Name: "play",
	    }
	    fmt.Print(f)
	}
	`
	
	const exampleKeyValueTopDeclPlay = `package main
	
	import (
	    "fmt"
	)
	
	var keyValueTopDecl = struct {
	    a string
	    b int
	}{
	    a: "B",
	    b: 2,
	}
	
	func main() {
	    fmt.Print(keyValueTopDecl)
	}
	`
	
	const exampleSortPlay = `package main
	
	import (
	    "fmt"
	    "sort"
	)
	
	// Person represents a person by name and age.
	type Person struct {
	    Name string
	    Age  int
	}
	
	// String returns a string representation of the Person.
	func (p Person) String() string {
	    return fmt.Sprintf("%s: %d", p.Name, p.Age)
	}
	
	// ByAge implements sort.Interface for []Person based on
	// the Age field.
	type ByAge []Person
	
	// Len returns the number of elements in ByAge.
	func (a ByAge) Len() int { return len(a) }
	
	// Swap swaps the elements in ByAge.
	func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
	
	// people is the array of Person
	var people = []Person{
	    {"Bob", 31},
	    {"John", 42},
	    {"Michael", 17},
	    {"Jenny", 26},
	}
	
	func main() {
	    fmt.Println(people)
	    sort.Sort(ByAge(people))
	    fmt.Println(people)
	}
	`
	
	func TestExamples(t *testing.T) {
	    fset := token.NewFileSet()
	    file, err := parser.ParseFile(fset, "test.go", strings.NewReader(exampleTestFile), parser.ParseComments)
	    if err != nil {
	        t.Fatal(err)
	    }
	    for i, e := range doc.Examples(file) {
	        c := exampleTestCases[i]
	        if e.Name != c.Name {
	            t.Errorf("got Name == %q, want %q", e.Name, c.Name)
	        }
	        if w := c.Play; w != "" {
	            g := formatFile(t, fset, e.Play)
	            if g != w {
	                t.Errorf("%s: got Play == %q, want %q", c.Name, g, w)
	            }
	        }
	        if g, w := e.Output, c.Output; g != w {
	            t.Errorf("%s: got Output == %q, want %q", c.Name, g, w)
	        }
	    }
	}
	
	const exampleWholeFile = `package foo_test
	
	type X int
	
	func (X) Foo() {
	}
	
	func (X) TestBlah() {
	}
	
	func (X) BenchmarkFoo() {
	}
	
	func Example() {
	    fmt.Println("Hello, world!")
	    // Output: Hello, world!
	}
	`
	
	const exampleWholeFileOutput = `package main
	
	type X int
	
	func (X) Foo() {
	}
	
	func (X) TestBlah() {
	}
	
	func (X) BenchmarkFoo() {
	}
	
	func main() {
	    fmt.Println("Hello, world!")
	}
	`
	
	func TestExamplesWholeFile(t *testing.T) {
	    fset := token.NewFileSet()
	    file, err := parser.ParseFile(fset, "test.go", strings.NewReader(exampleWholeFile), parser.ParseComments)
	    if err != nil {
	        t.Fatal(err)
	    }
	    es := doc.Examples(file)
	    if len(es) != 1 {
	        t.Fatalf("wrong number of examples; got %d want 1", len(es))
	    }
	    e := es[0]
	    if e.Name != "" {
	        t.Errorf("got Name == %q, want %q", e.Name, "")
	    }
	    if g, w := formatFile(t, fset, e.Play), exampleWholeFileOutput; g != w {
	        t.Errorf("got Play == %q, want %q", g, w)
	    }
	    if g, w := e.Output, "Hello, world!\n"; g != w {
	        t.Errorf("got Output == %q, want %q", g, w)
	    }
	}
	
	const exampleInspectSignature = `package foo_test
	
	import (
	    "bytes"
	    "io"
	)
	
	func getReader() io.Reader { return nil }
	
	func do(b bytes.Reader) {}
	
	func Example() {
	    getReader()
	    do()
	    // Output:
	}
	
	func ExampleIgnored() {
	}
	`
	
	const exampleInspectSignatureOutput = `package main
	
	import (
	    "bytes"
	    "io"
	)
	
	func getReader() io.Reader { return nil }
	
	func do(b bytes.Reader) {}
	
	func main() {
	    getReader()
	    do()
	}
	`
	
	func TestExampleInspectSignature(t *testing.T) {
	    // Verify that "bytes" and "io" are imported. See issue #28492.
	    fset := token.NewFileSet()
	    file, err := parser.ParseFile(fset, "test.go", strings.NewReader(exampleInspectSignature), parser.ParseComments)
	    if err != nil {
	        t.Fatal(err)
	    }
	    es := doc.Examples(file)
	    if len(es) != 2 {
	        t.Fatalf("wrong number of examples; got %d want 2", len(es))
	    }
	    // We are interested in the first example only.
	    e := es[0]
	    if e.Name != "" {
	        t.Errorf("got Name == %q, want %q", e.Name, "")
	    }
	    if g, w := formatFile(t, fset, e.Play), exampleInspectSignatureOutput; g != w {
	        t.Errorf("got Play == %q, want %q", g, w)
	    }
	    if g, w := e.Output, ""; g != w {
	        t.Errorf("got Output == %q, want %q", g, w)
	    }
	}
	
	const exampleEmpty = `
	package p
	func Example() {}
	func Example_a()
	`
	
	const exampleEmptyOutput = `package main
	
	func main() {}
	func main()
	`
	
	func TestExampleEmpty(t *testing.T) {
	    fset := token.NewFileSet()
	    file, err := parser.ParseFile(fset, "test.go", strings.NewReader(exampleEmpty), parser.ParseComments)
	    if err != nil {
	        t.Fatal(err)
	    }
	
	    es := doc.Examples(file)
	    if len(es) != 1 {
	        t.Fatalf("wrong number of examples; got %d want 1", len(es))
	    }
	    e := es[0]
	    if e.Name != "" {
	        t.Errorf("got Name == %q, want %q", e.Name, "")
	    }
	    if g, w := formatFile(t, fset, e.Play), exampleEmptyOutput; g != w {
	        t.Errorf("got Play == %q, want %q", g, w)
	    }
	    if g, w := e.Output, ""; g != w {
	        t.Errorf("got Output == %q, want %q", g, w)
	    }
	}
	
	func formatFile(t *testing.T, fset *token.FileSet, n *ast.File) string {
	    if n == nil {
	        return "<nil>"
	    }
	    var buf bytes.Buffer
	    if err := format.Node(&buf, fset, n); err != nil {
	        t.Fatal(err)
	    }
	    return buf.String()
	}
	
	// This example illustrates how to use NewFromFiles
	// to compute package documentation with examples.
	func ExampleNewFromFiles() {
	    // src and test are two source files that make up
	    // a package whose documentation will be computed.
	    const src = `
	// This is the package comment.
	package p
	
	import "fmt"
	
	// This comment is associated with the Greet function.
	func Greet(who string) {
	    fmt.Printf("Hello, %s!\n", who)
	}
	`
	    const test = `
	package p_test
	
	// This comment is associated with the ExampleGreet_world example.
	func ExampleGreet_world() {
	    Greet("world")
	}
	`
	
	    // Create the AST by parsing src and test.
	    fset := token.NewFileSet()
	    files := []*ast.File{
	        mustParse(fset, "src.go", src),
	        mustParse(fset, "src_test.go", test),
	    }
	
	    // Compute package documentation with examples.
	    p, err := doc.NewFromFiles(fset, files, "example.com/p")
	    if err != nil {
	        panic(err)
	    }
	
	    fmt.Printf("package %s - %s", p.Name, p.Doc)
	    fmt.Printf("func %s - %s", p.Funcs[0].Name, p.Funcs[0].Doc)
	    fmt.Printf(" ⤷ example with suffix %q - %s", p.Funcs[0].Examples[0].Suffix, p.Funcs[0].Examples[0].Doc)
	
	    // Output:
	    // package p - This is the package comment.
	    // func Greet - This comment is associated with the Greet function.
	    //  ⤷ example with suffix "world" - This comment is associated with the ExampleGreet_world example.
	}
	
	func TestClassifyExamples(t *testing.T) {
	    const src = `
	package p
	
	const Const1 = 0
	var   Var1   = 0
	
	type (
	    Type1     int
	    Type1_Foo int
	    Type1_foo int
	    type2     int
	
	    Embed struct { Type1 }
	)
	
	func Func1()     {}
	func Func1_Foo() {}
	func Func1_foo() {}
	func func2()     {}
	
	func (Type1) Func1() {}
	func (Type1) Func1_Foo() {}
	func (Type1) Func1_foo() {}
	func (Type1) func2() {}
	
	type (
	    Conflict          int
	    Conflict_Conflict int
	    Conflict_conflict int
	)
	
	func (Conflict) Conflict() {}
	`
	    const test = `
	package p_test
	
	func ExampleConst1() {} // invalid - no support for consts and vars
	func ExampleVar1()   {} // invalid - no support for consts and vars
	
	func Example()               {}
	func Example_()              {} // invalid - suffix must start with a lower-case letter
	func Example_suffix()        {}
	func Example_suffix_xX_X_x() {}
	func Example_世界()           {} // invalid - suffix must start with a lower-case letter
	func Example_123()           {} // invalid - suffix must start with a lower-case letter
	func Example_BadSuffix()     {} // invalid - suffix must start with a lower-case letter
	
	func ExampleType1()               {}
	func ExampleType1_()              {} // invalid - suffix must start with a lower-case letter
	func ExampleType1_suffix()        {}
	func ExampleType1_BadSuffix()     {} // invalid - suffix must start with a lower-case letter
	func ExampleType1_Foo()           {}
	func ExampleType1_Foo_suffix()    {}
	func ExampleType1_Foo_BadSuffix() {} // invalid - suffix must start with a lower-case letter
	func ExampleType1_foo()           {}
	func ExampleType1_foo_suffix()    {}
	func ExampleType1_foo_Suffix()    {} // matches Type1, instead of Type1_foo
	func Exampletype2()               {} // invalid - cannot match unexported
	
	func ExampleFunc1()               {}
	func ExampleFunc1_()              {} // invalid - suffix must start with a lower-case letter
	func ExampleFunc1_suffix()        {}
	func ExampleFunc1_BadSuffix()     {} // invalid - suffix must start with a lower-case letter
	func ExampleFunc1_Foo()           {}
	func ExampleFunc1_Foo_suffix()    {}
	func ExampleFunc1_Foo_BadSuffix() {} // invalid - suffix must start with a lower-case letter
	func ExampleFunc1_foo()           {}
	func ExampleFunc1_foo_suffix()    {}
	func ExampleFunc1_foo_Suffix()    {} // matches Func1, instead of Func1_foo
	func Examplefunc1()               {} // invalid - cannot match unexported
	
	func ExampleType1_Func1()               {}
	func ExampleType1_Func1_()              {} // invalid - suffix must start with a lower-case letter
	func ExampleType1_Func1_suffix()        {}
	func ExampleType1_Func1_BadSuffix()     {} // invalid - suffix must start with a lower-case letter
	func ExampleType1_Func1_Foo()           {}
	func ExampleType1_Func1_Foo_suffix()    {}
	func ExampleType1_Func1_Foo_BadSuffix() {} // invalid - suffix must start with a lower-case letter
	func ExampleType1_Func1_foo()           {}
	func ExampleType1_Func1_foo_suffix()    {}
	func ExampleType1_Func1_foo_Suffix()    {} // matches Type1.Func1, instead of Type1.Func1_foo
	func ExampleType1_func2()               {} // matches Type1, instead of Type1.func2
	
	func ExampleEmbed_Func1() {} // invalid - no support for forwarded methods from embedding
	
	func ExampleConflict_Conflict()        {} // ambiguous with either Conflict or Conflict_Conflict type
	func ExampleConflict_conflict()        {} // ambiguous with either Conflict or Conflict_conflict type
	func ExampleConflict_Conflict_suffix() {} // ambiguous with either Conflict or Conflict_Conflict type
	func ExampleConflict_conflict_suffix() {} // ambiguous with either Conflict or Conflict_conflict type
	`
	
	    // Parse literal source code as a *doc.Package.
	    fset := token.NewFileSet()
	    files := []*ast.File{
	        mustParse(fset, "src.go", src),
	        mustParse(fset, "src_test.go", test),
	    }
	    p, err := doc.NewFromFiles(fset, files, "example.com/p")
	    if err != nil {
	        t.Fatalf("doc.NewFromFiles: %v", err)
	    }
	
	    // Collect the association of examples to top-level identifiers.
	    got := map[string][]string{}
	    got[""] = exampleNames(p.Examples)
	    for _, f := range p.Funcs {
	        got[f.Name] = exampleNames(f.Examples)
	    }
	    for _, t := range p.Types {
	        got[t.Name] = exampleNames(t.Examples)
	        for _, f := range t.Funcs {
	            got[f.Name] = exampleNames(f.Examples)
	        }
	        for _, m := range t.Methods {
	            got[t.Name+"."+m.Name] = exampleNames(m.Examples)
	        }
	    }
	
	    want := map[string][]string{
	        "": {"", "suffix", "suffix_xX_X_x"}, // Package-level examples.
	
	        "Type1":     {"", "foo_Suffix", "func2", "suffix"},
	        "Type1_Foo": {"", "suffix"},
	        "Type1_foo": {"", "suffix"},
	
	        "Func1":     {"", "foo_Suffix", "suffix"},
	        "Func1_Foo": {"", "suffix"},
	        "Func1_foo": {"", "suffix"},
	
	        "Type1.Func1":     {"", "foo_Suffix", "suffix"},
	        "Type1.Func1_Foo": {"", "suffix"},
	        "Type1.Func1_foo": {"", "suffix"},
	
	        // These are implementation dependent due to the ambiguous parsing.
	        "Conflict_Conflict": {"", "suffix"},
	        "Conflict_conflict": {"", "suffix"},
	    }
	
	    for id := range got {
	        if !reflect.DeepEqual(got[id], want[id]) {
	            t.Errorf("classification mismatch for %q:\ngot  %q\nwant %q", id, got[id], want[id])
	        }
	    }
	}
	
	func exampleNames(exs []*doc.Example) (out []string) {
	    for _, ex := range exs {
	        out = append(out, ex.Suffix)
	    }
	    return out
	}
	
	func mustParse(fset *token.FileSet, filename, src string) *ast.File {
	    f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	    if err != nil {
	        panic(err)
	    }
	    return f
	}
```

</details>

### <a name="Package.Filter">func</a> (\*Package) [Filter](./filter.go#L99)
``` go
func (p *Package) Filter(f Filter)
```
Filter eliminates documentation for names that don't pass through the filter f.
TODO(gri): Recognize "Type.Method" as a name.

## <a name="Type">type</a> [Type](./doc.go#L50-L65)
``` go
type Type struct {
    Doc  string
    Name string
    Decl *ast.GenDecl

    // associated declarations
    Consts  []*Value // sorted list of constants of (mostly) this type
    Vars    []*Value // sorted list of variables of (mostly) this type
    Funcs   []*Func  // sorted list of functions returning this type
    Methods []*Func  // sorted list of methods (including embedded ones) of this type

    // Examples is a sorted list of examples associated with
    // this type. Examples are extracted from _test.go files
    // provided to NewFromFiles.
    Examples []*Example
}

```
Type is the documentation for a type declaration.

## <a name="Value">type</a> [Value](./doc.go#L41-L47)
``` go
type Value struct {
    Doc   string
    Names []string // var or const names in declaration order
    Decl  *ast.GenDecl
    // contains filtered or unexported fields
}

```
Value is the documentation for a (possibly grouped) var or const declaration.

- - -
Generated by [godoc2ghmd](https://github.com/gglachant/godoc2ghmd)
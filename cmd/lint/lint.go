package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"

	"golang.org/x/tools/go/loader"
)

var (
	recursivelyIgnoredPkgs arrayFlag
	ignoredPkgs            arrayFlag
	ignoredFiles           arrayFlag
	allowedFinishers       arrayFlag = []string{"Msg", "Msgf"}
	rootPkg                string
)

func init() {
	flag.Var(&recursivelyIgnoredPkgs, "ignorePkgRecursively", "ignore the specified package and all subpackages recursively")
	flag.Var(&ignoredPkgs, "ignorePkg", "ignore the specified package")
	flag.Var(&ignoredFiles, "ignoreFile", "ignore the specified file by its path and/or go path (package/file.go)")
	flag.Var(&allowedFinishers, "finisher", "allowed finisher for the event chain")
	flag.Parse()

	recursivelyIgnoredPkgs = append(recursivelyIgnoredPkgs, "github.com/rs/zerolog")
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "you must provide exactly one package path")
		os.Exit(1)
	}
	rootPkg = args[0]
}

func main() {

	conf := loader.Config{}
	conf.Import(rootPkg)
	p, err := conf.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: unable to load the root package. %s\n", err.Error())
		os.Exit(1)
	}

	event := getEvent(p)
	if event == nil {
		fmt.Fprintln(os.Stderr, "Error: github.com/rs/zerolog.Event declaration not found, maybe zerolog is not imported in the scanned package?")
		os.Exit(1)
	}

	selections := getSelectionsWithReceiverType(p, event)

	hasViolations := false
	for _, s := range selections {
		if hasBadFinisher(p, s) {
			hasViolations = true
			fmt.Printf("Error: missing or bad finisher for log chain, last call: %q at: %s:%v\n", s.fn.Name(), p.Fset.File(s.Pos()).Name(), p.Fset.Position(s.Pos()).Line)
		}
	}

	if !hasViolations {
		fmt.Println("No violations found")
		return
	}

	os.Exit(1)
}

func getEvent(p *loader.Program) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func getSelectionsWithReceiverType(p *loader.Program, targetType types.Type) map[token.Pos]selection {
	_ = "STUB: not implemented"
	return nil
}

func hasBadFinisher(p *loader.Program, s selection) bool { _ = "STUB: not implemented"; return false }

type arrayFlag []string

func (i *arrayFlag) String() string { _ = "STUB: not implemented"; return "" }

func (i *arrayFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }

type selection struct {
	*ast.SelectorExpr
	fn  *types.Func
	pkg *types.Package
}

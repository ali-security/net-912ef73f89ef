// Standalone helper: package a Go module directory into a proxy.golang.org
// module zip using golang.org/x/mod/zip.CreateFromDir.
//
// Usage: create_mod <module-path> <version> <src-dir> <out-zip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage: %s <module-path> <version> <src-dir> <out-zip>\n", os.Args[0])
		os.Exit(2)
	}
	modPath := os.Args[1]
	version := os.Args[2]
	srcDir := os.Args[3]
	outZip := os.Args[4]

	f, err := os.Create(outZip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create out zip: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "CreateFromDir: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s for %s@%s\n", outZip, modPath, version)
}

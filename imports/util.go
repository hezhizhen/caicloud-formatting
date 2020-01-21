package imports

import (
	"fmt"
	"os"
	"strings"
)

// Check checks if the error is nil, and panics if it is not
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// appendPackage appends a package to the list.
func appendPackage(pkgs []string, pkg string, first bool) []string {
	// If the package the first one in a new category and
	// the list is not empty which means there are some packages that belong
	// to other categories,
	// add a blank line to separate them.
	if first && len(pkgs) != 0 {
		pkgs = append(pkgs, "")
	}
	pkgs = append(pkgs, fmt.Sprintf("\t%s", pkg))
	return pkgs
}

// extractPackage extracts an imported package from the given line and keeps a
// copy of the line without leading and/or trailing spaces.
func extractPackage(line string) (pkg, origin string) {
	// trim leading and/or trailing spaces
	line = strings.TrimSpace(line)
	origin = line
	// skip commented line
	if strings.HasPrefix(line, "//") {
		return
	}
	pkg = line
	// trim alias
	parts := strings.Split(line, " ")
	if len(parts) == 2 {
		pkg = parts[1]
	}
	// trim double quotes
	pkg = strings.TrimPrefix(pkg, "\"")
	pkg = strings.TrimSuffix(pkg, "\"")
	return pkg, origin
}

// local is the name of the current repository.
// e.g.: /Users/caicloud/go/src/github.com/caicloud/config-admin -> config-admin
var local = func() string {
	dir, err := os.Getwd()
	Check(err)
	parts := strings.Split(dir, "github.com/caicloud/")
	return strings.Split(parts[1], "/")[0]
}()

// classify finds the correct category for the imported package.
func classify(pkg string) string {
	parts := strings.Split(pkg, "/")
	// standard libraries may have zero, one or two slashes, but none have a dot
	// other packages must be hosted in somewhere, which means the first part
	// has a dot
	if !strings.Contains(parts[0], ".") {
		return "std"
	}
	// packages defined in repositories of caicloud have caicloud in their paths
	// while others don't
	if !strings.Contains(pkg, "caicloud") {
		return "others"
	}
	// packages defined in the current repository have the repository's name in
	// their paths while others don't
	if strings.Contains(pkg, local) {
		return "local"
	}
	return "company"
}

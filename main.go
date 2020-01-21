package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hezhizhen/caicloud-formatting/imports"

	"k8s.io/klog"
)

// e.g.: go run format/main.go .
func main() {
	// the default root is the current directory, but you can specify it
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			klog.Error(err)
			return nil
		}
		// skip some files and directories
		if strings.HasPrefix(path, "vendor") || // vendor
			info.IsDir() || // directories
			strings.HasPrefix(path, ".") || // hidden files
			!strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		// read file
		bs, err := ioutil.ReadFile(path)
		imports.Check(err)
		lines := strings.Split(string(bs), "\n")
		ims, s, e := extractImports(lines)
		// write file
		var ret []string
		for i, line := range lines {
			if i == s {
				ret = append(ret, ims...)
			}
			if i >= s && i <= e {
				continue
			}
			ret = append(ret, line)
		}
		// more details about permissions: https://ss64.com/bash/chmod.html
		err = ioutil.WriteFile(path, []byte(strings.Join(ret, "\n")), 0644)
		imports.Check(err)
		return nil
	})
	imports.Check(err)
}

// extractImports extract import lines from the file and records the indices of
// start and end lines
func extractImports(lines []string) ([]string, int, int) {
	var ret []string
	// start and end indicate the first and last lines of imported packages
	start, end := -1, -1
	for i, line := range lines {
		// NOTE: suppose all imported packages are all at one place
		// check if the current line is `import (`
		if strings.HasPrefix(line, "import") {
			start = i + 1
			// If a line starts with import does not end with (,
			// there's only one import for the file.
			// Reset start to indicate doing nothing to the file.
			if !strings.HasSuffix(line, "(") {
				start = -1
			}
			continue
		}
		// check if the current line is `)`
		if start != -1 && end == -1 {
			if line == ")" {
				end = i - 1
			}
		}
		// if the current line is neither `import (` nor `)`, copy it
		if start != -1 && end == -1 {
			ret = append(ret, line)
		}
	}
	ret = imports.UpdateOrder(ret)
	return ret, start, end
}

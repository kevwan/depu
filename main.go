package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	directs, err := getDirectModules()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parseGoMod: %v\n", err)
		return
	}

	modules, err := getDepPackages()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parseModules: %v\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Package", "Current", "Latest", "GoVersion"})
	table.SetBorder(false)
	for _, mod := range modules {
		if mod.Update == nil {
			continue
		}

		if !contains(directs, mod.Path) {
			continue
		}

		table.Append([]string{
			mod.Path,
			mod.Version,
			mod.Update.Version,
			mod.GoVersion,
		})
	}
	table.Render()
}

// contains checks if str is in list.
func contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
}
